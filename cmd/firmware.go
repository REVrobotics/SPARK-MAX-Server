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

	"github.com/spf13/cobra"
	sparkmax "github.com/willtoth/USB-BLDC-TOOL/sparkmax"
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
	into the device. To get the firmware pass the -u flag.

The command will block until the firmware is updated. Be sure that
the device is plugged in and power is not removed during the entire
update.`,
	Run:  firmware,
	Args: cobra.MaximumNArgs(1),
}}

func init() {
	rootCmd.AddCommand(&firmwareCmd.Command)
	sparkmax.RegisterCommand(firmwareCmd)

	firmwareCmd.Flags().BoolVarP(&update, "update", "u", false, "Get current firmware version from device")
}

func Firmware(command *sparkmax.FirmwareRequest) (*sparkmax.FirmwareResponse, error) {
	var resp sparkmax.FirmwareResponse
	var err error
	var frameIn sparkmax.UsbFrame
	frame := sparkmax.BroadcastFrame()

	if command.Filename == "" {
		frame.Header.API = sparkmax.CmdBcastFirmware

		frameIn, err = sparkmax.SparkWriteFrame(frame)

		resp.Version = fmt.Sprintf("v%d.%d.%d", frameIn.Data[0], frameIn.Data[1], uint16(frameIn.Data[2])<<8|uint16(frameIn.Data[3]))

		if frameIn.Data[4] == 1 {
			resp.Version += ", Debug build"
		}
	} else {
		//TODO: Firmware update
	}

	return &resp, err
}

func firmware(cmd *cobra.Command, args []string) {
	if update == true {
		fmt.Println("Firmware update is not implemented at this time")
	} else {
		//Return the firmware version
		req := sparkmax.FirmwareRequest{}
		resp, err := Firmware(&req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get firmware: %s\r\n", err.Error())
			return
		}

		fmt.Printf("Firmware Version: %s", resp.Version)
	}
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
	return "FirmwareRequest"
}
