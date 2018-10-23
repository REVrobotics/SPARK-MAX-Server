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
	sparkmax "github.com/willtoth/USB-BLDC-TOOL/sparkmax"
)

var listAll bool
var verbose bool

func listDevices(cmd *cobra.Command, args []string) {
	ports := sparkmax.ListDevices(listAll)
	for _, port := range ports {
		if verbose {
			fmt.Printf("Device: %v\n", port)
		} else {
			fmt.Printf("Device: %v %v\n", port.SerialNumber, port.Name)
		}
	}
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
	//overwrite these since we don't want ot connect/disconnect
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	listCmd.PersistentFlags().BoolVarP(&listAll, "all", "a", false, "List all devices including over CAN")
	listCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "List more details for devices")

	sparkmax.RegisterCommand(&listDevicesCmd)
}

type listDevicesCommand struct{}

var listDevicesCmd = listDevicesCommand{}

func (s *listDevicesCommand) SparkCommandProcess(req sparkmax.RequestWire) (resp sparkmax.ResponseWire, err error) {
	ports := sparkmax.ListDevices(req.GetList().All)

	devList := make([]string, 0)
	devDetails := make([]string, 0)
	for _, port := range ports {
		devList = append(devList, port.Name)

		result := fmt.Sprintf("Device: %s,\t%s:%s\t%s", port.SerialNumber, port.VID, port.PID, port.Name)
		devDetails = append(devDetails, result)
	}
	tmp := sparkmax.ListResponse{DeviceList: devList, DeviceDetails: devDetails}
	resp.Resp = &sparkmax.ResponseWire_List{List: &tmp}
	return resp, err
}

func (s *listDevicesCommand) ExpectedType() string {
	return "List"
}
