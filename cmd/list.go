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

	"github.com/spf13/cobra"
	sparkusb "github.com/willtoth/USB-BLDC-TOOL/sparkusb"
)

var listAll bool
var verbose bool

func listDevices(cmd *cobra.Command, args []string) {
	ports := sparkusb.ListDevices(listAll)
	for _, port := range ports {
		if verbose {
			fmt.Printf("Device: %v\n", port)
		} else {
			fmt.Printf("Device: %v %v\n", port.SerialNumber, port.Name)
		}
	}

	spName := sparkusb.GetDefaultDevice()
	frame := sparkusb.DefaultFrame()

	frame.Data[1] = 1
	frame.Data[3] = 3
	frame.Header.DeviceType = 2

	err := sparkusb.Connect(spName)
	if err != nil {
		fmt.Println(err)
	}

	err = sparkusb.RunCommand(frame, spName, false)
	if err != nil {
		fmt.Println(err)
	}

	sparkusb.Disconnect()
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available CAN BLDC devices",
	Long: `List of available devices either connected to
USB or selected device

Use this command to list available connected devices if
more than one device is connected. Output of this command
can be used to specify device for other commands`,
	Run: listDevices,
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	listCmd.PersistentFlags().BoolVarP(&listAll, "all", "a", false, "List all devices including over CAN")
	listCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "List more details for devices")
}
