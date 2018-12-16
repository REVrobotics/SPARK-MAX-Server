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

	sparkmax "github.com/REVrobotics/SPARK-MAX-Server/sparkmax"
	"github.com/spf13/cobra"
)

// enable and send heartbeat
var enableMode bool

type setpointCommand struct {
	cobra.Command
}

// setpointCmd represents the setpoint command
var setpointCmd = &setpointCommand{cobra.Command{
	Use:   "setpoint",
	Short: "Set the controller setpoint",
	Long: `Set the controller setpoint. Use the -e flag
to also send an enable heartbeat.

To set a negative number, use the -- argument, example:

sparkmax.exe setpoint -- -0.75

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
		/*
			if enableMode {
				req := sparkmax.HeartbeatRequest{Enable: true}
				_, err := sparkmax.Heartbeat(&req)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Heartbeat command failed: %s", err.Error())
					return
				}
			}
		*/

		setpoint, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Setpoint argument float conversion failed: %s", err.Error())
			return
		}

		req := sparkmax.SetpointRequest{Setpoint: float32(setpoint), Enable: enableMode}
		_, err = runSetpoint(&req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Setpoint command failed: %s", err.Error())
			return
		}
	},
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"run", "Run", "Setpoint"},
	PreRun:  preRunConnect,
	PostRun: postRunDisconnect,
}}

func sendHeartbeat(enable bool) error {
	frame := sparkmax.DefaultFrame()

	frame.Header.API = sparkmax.CmdApiHeartbeat

	if enable {
		frame.Data[0] = 1
	}

	_, err := sparkmax.SparkWriteFrame(frame)

	return err
}

func runSetpoint(command *sparkmax.SetpointRequest) (*sparkmax.SetpointResponse, error) {
	var resp sparkmax.SetpointResponse
	var err error
	frame := sparkmax.DefaultFrame()

	frame.Header.API = sparkmax.CmdApiDcSet

	if command.Enable {
		err = sendHeartbeat(command.Enable)
		if err != nil {
			return &resp, err
		}
	}

	if command.Setpoint < 0.001 && command.Setpoint > -0.001 {
		frame.Data[0] = 0
		frame.Data[1] = 0
		frame.Data[2] = 0
		frame.Data[3] = 0
	} else {
		//TODO: Implement a min/max based on user setting
		tmparray := Float32ToBytes(command.Setpoint)

		copy(frame.Data[:4], tmparray[:])
	}

	_, err = sparkmax.SparkWriteFrame(frame)

	return &resp, err
}

func (s *setpointCommand) SparkCommandProcess(req sparkmax.RequestWire) (resp sparkmax.ResponseWire, err error) {
	r, err := runSetpoint(req.GetSetpoint())
	if err != nil {
		tmp := sparkmax.RootResponse{Error: err.Error()}
		r.Root = &tmp
	}
	resp.Resp = &sparkmax.ResponseWire_Setpoint{Setpoint: r}
	return resp, err
}

func (s *setpointCommand) ExpectedType() string {
	return "Setpoint"
}

func init() {
	rootCmd.AddCommand(&setpointCmd.Command)
	setpointCmd.Flags().BoolVarP(&enableMode, "enable", "e", false, "Send heartbeat with enable")

	sparkmax.RegisterCommand(setpointCmd)
}
