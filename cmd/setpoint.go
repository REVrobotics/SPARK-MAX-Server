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
	"strconv"

	"github.com/spf13/cobra"
	sparkusb "github.com/willtoth/USB-BLDC-TOOL/sparkusb"
)

// enable and send heartbeat
var enableMode bool

// setpointCmd represents the setpoint command
var setpointCmd = &cobra.Command{
	Use:   "setpoint",
	Short: "Set the controller setpoint",
	Long: `Set the controller setpoint. Use the -e flag
to also send an enable heartbeat.

**Note** If not run in interactive
mode (not implemented yet) you cannot reliably set the 
motor controller as this shell will connect --> 
enable --> send setpoint --> disconnect on every call, 
and this will likely take more than 100ms on your system, 
which will disable the controller. Best not to be 
controlling powerful motors from the command line. Use the 
built in remote server instead to control the motor with a 
GUI.`,
	Run: func(cmd *cobra.Command, args []string) {
		if enableMode {
			req := sparkusb.HeartbeatRequest{Enable: true}
			_, err := sparkusb.Heartbeat(&req)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Heartbeat command failed: %s", err.Error())
				return
			}
		}

		setpoint, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Setpoint argument float conversion failed: %s", err.Error())
			return
		}

		req := sparkusb.SetpointRequest{Setpoint: float32(setpoint)}
		_, err = sparkusb.Setpoint(&req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Setpoint command failed: %s", err.Error())
			return
		}
	},
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"run", "Run", "Setpoint"},
}

func init() {
	rootCmd.AddCommand(setpointCmd)
	setpointCmd.Flags().BoolVarP(&enableMode, "enable", "e", false, "Send heartbeat with enable")
}
