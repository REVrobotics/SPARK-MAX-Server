// Copyright © 2018 REV Robotics LLC (support@revrobotics.com)
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

	sparkmax "github.com/REVrobotics/SPARK-MAX-Server/sparkmax"
	"github.com/spf13/cobra"
)

// clearFaultsCmd represents the address command
var clearFaultsCmd = &cobra.Command{
	Use:     "clearfaults",
	Short:   "Clear sticky faults",
	Long:    `Clear the sticky faults of the device.`,
	Aliases: []string{"clear", "clr"},
	Run:     runClearFaults,
	PreRun:  preRunConnect,
	PostRun: postRunDisconnect,
}

func init() {
	rootCmd.AddCommand(clearFaultsCmd)
}

func ClearFaults(command *sparkmax.ClearFaultsRequest) (*sparkmax.ClearFaultsResponse, error) {
	var resp sparkmax.ClearFaultsResponse
	frame := sparkmax.DefaultFrame()

	frame.Header.API = sparkmax.CmdApiClrFaults
	_, err := sparkmax.SparkWriteFrame(frame)
	if err != nil {
		tmp := sparkmax.RootResponse{Error: err.Error()}
		resp.Root = &tmp
	}
	return &resp, err
}

func runClearFaults(cmd *cobra.Command, args []string) {
	req := sparkmax.ClearFaultsRequest{}

	_, err := ClearFaults(&req)
	if err != nil {
		fmt.Println("Failed to clear faults: ", err)
	}
}
