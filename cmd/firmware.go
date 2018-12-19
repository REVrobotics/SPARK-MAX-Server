// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"sync"
	"time"

	sparkmax "github.com/REVrobotics/SPARK-MAX-Server/sparkmax"
	"github.com/spf13/cobra"
	"github.com/willtoth/go-dfuse/dfudevice"
	"github.com/willtoth/go-dfuse/dfufile"
	pb "gopkg.in/cheggaaa/pb.v1"
)

var update bool

type firmwareCommand struct {
	cobra.Command
}

// firmwareCmd represents the firmware command
var firmwareCmd = &firmwareCommand{cobra.Command{
	Use:   "firmware",
	Short: "Get firmware version or update",
	Long: `Get the firmware version or program new firmware
	into the device. To update the firmware pass the file path 
	to the .dfu file.

The command will block until the firmware is updated. Be sure that
the device is plugged in and power is not removed during the entire
update.`,
	Run:  firmware,
	Args: cobra.MaximumNArgs(1),
}}

func init() {
	rootCmd.AddCommand(&firmwareCmd.Command)
	sparkmax.RegisterCommand(firmwareCmd)
}

func sendBootloaderCommand() {
	frame := sparkmax.DefaultFrame()
	frame.Header.API = sparkmax.ExtCmdBootloader
	frame.Header.CommandType = sparkmax.CmdTypeExtended
	sparkmax.SparkWriteFrame(frame)
}

func Firmware(command *sparkmax.FirmwareRequest) (*sparkmax.FirmwareResponse, error) {
	var resp sparkmax.FirmwareResponse
	var err error
	var frameIn sparkmax.UsbFrame
	frame := sparkmax.BroadcastFrame()

	if command.Filename == "" {
		frame.Header.API = sparkmax.CmdBcastFirmware

		frameIn, err = sparkmax.SparkWriteFrame(frame)

		resp.Major = uint32(frameIn.Data[0])
		resp.Minor = uint32(frameIn.Data[1])
		resp.Build = uint32((uint16(frameIn.Data[2])<<8 | uint16(frameIn.Data[3])))
		resp.IsDebug = false
		resp.HardwareVersion = string(frameIn.Data[5])

		resp.Version = fmt.Sprintf("v%d.%d.%d", resp.Major, resp.Minor, resp.Build)

		if frameIn.Data[4] == 1 {
			resp.Version += ", Debug build"
			resp.IsDebug = true
		}

		resp.UpdateStarted = false

	} else {
		if sparkmax.IsConnected() != true {
			err := sparkmax.Connect(Device)
			if err == nil {
				fmt.Println("Entering bootloader...")
				sendBootloaderCommand()
				sparkmax.Disconnect()
			}
		}

		//Wait for up to 5 seconds for device to enter DFU mode
		for timeToWait := 0; timeToWait < 10; timeToWait++ {
			foundDevices := dfudevice.List(SPARKMAXDFUVID, SPARKMAXDFUPID)

			if len(foundDevices) != 0 {
				break
			}

			time.Sleep(500 * time.Millisecond)
		}

		err = startFirmwareUpdate(command.Filename)

		if err != nil {
			tmp := sparkmax.RootResponse{Error: err.Error()}
			resp.Root = &tmp
			resp.UpdateStarted = false
		}

		resp.UpdateStarted = true
	}

	return &resp, err
}

func firmware(cmd *cobra.Command, args []string) {
	if len(args) == 1 {
		req := sparkmax.FirmwareRequest{Filename: args[0]}
		resp, err := Firmware(&req)
		if err != nil || resp.UpdateStarted != true {
			fmt.Fprintf(os.Stderr, "Failed to upload firmware: %v\r\n", err)
			return
		}

		err = firmwareThread.waitOnFirmwareUpdate(20)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to upload firmware after start: %v\r\n", err)
			return
		}

	} else {
		//Run this here so we don't connect during update
		preRunConnect(cmd, args)
		//Return the firmware version
		req := sparkmax.FirmwareRequest{}
		resp, err := Firmware(&req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get firmware: %v\r\n", err)
			return
		}

		fmt.Printf("Firmware Version: %s", resp.Version)
		postRunDisconnect(cmd, args)
	}
}

type consoleProgress struct {
	pb  *pb.ProgressBar
	inc uint
	max uint
}

func (c *consoleProgress) Reset() {
	c.pb.Reset(int(c.max))
	c.pb.Set(0)
	c.pb.Update()
	c.pb.Start()
}

func (c *consoleProgress) Increment() {
	c.pb.Add(int(c.inc))
	c.pb.Update()
}

func (c *consoleProgress) SetStatus(status string) {
	c.pb.Prefix(status)
}

func (c *consoleProgress) SetIncrement(increment uint) {
	c.inc = increment
}

func (c *consoleProgress) SetMax(max uint) {
	c.pb.SetTotal(int(max))
	c.max = max
}

func StartNew() consoleProgress {
	var c consoleProgress
	c.pb = pb.New(1)
	c.pb.SetMaxWidth(120)
	c.pb.ShowTimeLeft = false

	//Manually update the progress bar
	c.pb.SetRefreshRate(time.Second * 10000)
	return c
}

const (
	SPARKMAXDFUVID = 0x0483
	SPARKMAXDFUPID = 0xdf11
)

type firmwareUpdateThread struct {
	sync.Mutex
	status  string
	running bool
	err     error
}

var firmwareThread firmwareUpdateThread

func startFirmwareUpdate(filename string) error {
	firmwareThread.Lock()
	if firmwareThread.running == true {
		return fmt.Errorf("Firmware update already running, must wait for completion")
	}
	firmwareThread.Unlock()

	firmwareThread = firmwareUpdateThread{}
	firmwareThread.running = true

	go firmwareThread.updateFirmware(filename)

	return nil
}

//Block until firmware update is done
func (f *firmwareUpdateThread) waitOnFirmwareUpdate(timeout time.Duration) error {
	timeLeft := timeout * 1000

	pollTime := time.Duration(100)

	for timeLeft > 0 {
		firmwareThread.Lock()
		running := firmwareThread.running
		firmwareThread.Unlock()

		if running == false {
			break
		}

		time.Sleep(pollTime * time.Millisecond)
		timeLeft = timeLeft - pollTime
	}

	firmwareThread.Lock()
	err := firmwareThread.err
	firmwareThread.Unlock()

	return err
}

func (f *firmwareUpdateThread) updateFirmware(filename string) {
	defer func() {
		f.Lock()
		f.running = false
		f.Unlock()
	}()

	fmt.Println("Starting Update!")

	foundDevices := dfudevice.List(SPARKMAXDFUVID, SPARKMAXDFUPID)

	if len(foundDevices) == 0 {
		f.Lock()
		f.err = fmt.Errorf("No DFU Devices Found")
		f.Unlock()
		return
	}

	dfu, err := dfufile.Read(filename)

	f.Lock()
	if err != nil {
		f.err = fmt.Errorf("DFU File Format Failed: %v", err)
		return
	}
	f.status = "Connecting to device..."
	f.Unlock()
	fmt.Println(f.status)

	dev, err := dfudevice.Open(SPARKMAXDFUVID, SPARKMAXDFUPID)
	defer dev.Close()

	if err != nil {
		f.Lock()
		f.err = fmt.Errorf("Failed to initialize: %v", err)
		f.Unlock()
		return
	}

	bar := StartNew()
	dev.RegisterProgress(&bar)

	err = dfudevice.WriteImage(dfu.Images[0], dev)

	if err != nil {
		f.Lock()
		f.err = fmt.Errorf("Write DFUFile Failed %v", err)
		f.Unlock()
		return
	}

	verify, err := dfudevice.VerifyImage(dfu.Images[0], dev)

	if err != nil || verify == false {
		f.Lock()
		f.err = fmt.Errorf("Failed to verify DFU Image: %v", err)
		f.Unlock()
		return
	}

	err = dev.ExitDFU(uint(dfu.Images[0].Targets[0].Prefix.Address))

	if err != nil || verify == false {
		f.Lock()
		f.err = fmt.Errorf("Failed to exit DFU mode: %v", err)
		f.Unlock()
		return
	}

	fmt.Println("")
	fmt.Println("Success!")
}

func (s *firmwareCommand) SparkCommandProcess(req sparkmax.RequestWire) (resp sparkmax.ResponseWire, err error) {
	r, err := Firmware(req.GetFirmware())
	if err != nil {
		tmp := sparkmax.RootResponse{Error: err.Error()}
		r.Root = &tmp
	}
	resp.Resp = &sparkmax.ResponseWire_Firmware{Firmware: r}
	return resp, err
}

func (s *firmwareCommand) ExpectedType() string {
	return "Firmware"
}
