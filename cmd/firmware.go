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
	"encoding/binary"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	sparkusb "github.com/willtoth/USB-BLDC-TOOL/sparkusb"
)

var update bool

// firmwareCmd represents the firmware command
var firmwareCmd = &cobra.Command{
	Use:   "firmware",
	Short: "Get firmware version or update",
	Long: `Get the firmware version or program new firmware
	into the device. To get the firmware pass the -u flag.

The command will block until the firmware is updated. Be sure that
the device is plugged in and power is not removed during the entire
update.`,
	Run:  firmware,
	Args: cobra.MaximumNArgs(1),
}

func init() {
	rootCmd.AddCommand(firmwareCmd)

	firmwareCmd.Flags().BoolVarP(&update, "update", "u", false, "Get current firmware version from device")
}

func firmware(cmd *cobra.Command, args []string) {
	if update == true {
		fmt.Println("Firmware update is not implemented at this time")
	} else {
		//Return the firmware version
		req := sparkusb.ParameterRequest{Parameter: sparkusb.ConfigParam_FirmwareVersion}
		resp, err := sparkusb.GetParameter(&req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get firmware: %s\r\n", err.Error())
			//return
		}

		tmp, err := strconv.ParseUint(resp.Value, 10, 32)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get firmware: %s\r\n", err.Error())
			//return
		}

		tmpBytes := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmpBytes, uint32(tmp))

		versionMajor := tmpBytes[0]
		versionMinor := tmpBytes[1]
		versionBuild := (int(tmpBytes[3]) | int(tmpBytes[2])<<8)
		fmt.Printf("Firmware Version: v%d.%d.%d", versionMajor, versionMinor, versionBuild)
	}
}
