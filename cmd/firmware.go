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

		resp.UpdateSuccess = false

	} else {
		err := updateFirmware(command.Filename)

		if err != nil {
			tmp := sparkmax.RootResponse{Error: err.Error()}
			resp.Root = &tmp
			resp.UpdateSuccess = false
		}

		resp.UpdateSuccess = true
	}

	return &resp, err
}

func firmware(cmd *cobra.Command, args []string) {
	if len(args) == 1 {
		updateFirmware(args[0])
	} else {
		//Run this here so we don't connect during update
		preRunConnect(cmd, args)
		//Return the firmware version
		req := sparkmax.FirmwareRequest{}
		resp, err := Firmware(&req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get firmware: %s\r\n", err.Error())
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

func updateFirmware(filename string) error {
	foundDevices := dfudevice.List(SPARKMAXDFUVID, SPARKMAXDFUPID)

	if len(foundDevices) == 0 {
		fmt.Println("No DFU Devices Found")
		return nil
	}

	dfu, err := dfufile.Read(filename)

	if err != nil {
		fmt.Println("DFU File Format Failed: ", err)
		return err
	}

	fmt.Println("Connecting to device...")

	dev, err := dfudevice.Open(SPARKMAXDFUVID, SPARKMAXDFUPID)
	defer dev.Close()

	if err != nil {
		fmt.Println("Failed to initialize ", err)
		return err
	}

	bar := StartNew()
	dev.RegisterProgress(&bar)

	err = dfudevice.WriteImage(dfu.Images[0], dev)

	if err != nil {
		fmt.Println("Write DFUFile Failed ", err)
		return err
	}

	verify, err := dfudevice.VerifyImage(dfu.Images[0], dev)

	if err != nil || verify == false {
		fmt.Println("Failed to verify DFU Image: ", err)
		return err
	}

	err = dev.ExitDFU(uint(dfu.Images[0].Targets[0].Prefix.Address))

	if err != nil || verify == false {
		fmt.Println("Failed to exit DFU mode: ", err)
		return err
	}

	fmt.Println("")
	fmt.Println("Success!")

	return nil
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
