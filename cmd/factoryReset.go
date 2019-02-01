// Copyright © 2018 - 2019 REV Robotics LLC (support@revrobotics.com)
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

type factoryResetCommand struct {
	cobra.Command
}

// full rset
var doFullReset bool

// Don't burn table
var dontBurnTable bool

// factoryResetCmd represents the factory reset command
var factoryResetCmd = &factoryResetCommand{cobra.Command{
	Use:   "factoryreset",
	Short: "Reset the controller to factory defaults",
	Long: `Reset the controller to factory defaults. Use the
	--full to include parameters that are not normally marked for reset
	including CAN ID and others.`,
	Run:     runFactoryReset,
	Aliases: []string{"factory", "paramreset"},
	PreRun:  preRunConnect,
	PostRun: postRunDisconnect,
}}

func init() {
	factoryResetCmd.Flags().BoolVar(&doFullReset, "full", false, "Reset all parameters including CAN ID and others")
	factoryResetCmd.Flags().BoolVar(&dontBurnTable, "no-burn", false, "Don't burn the parameter table after writing, after power cycle values will revert to previous values")
	rootCmd.AddCommand(&factoryResetCmd.Command)
	sparkmax.RegisterCommand(factoryResetCmd)
}

func FactoryReset(command *sparkmax.FactoryResetRequest) (*sparkmax.RootResponse, error) {
	var resp sparkmax.RootResponse
	var err error
	frame := sparkmax.DefaultFrame()

	if command.GetFullWipe() == true {
		frame.Header.API = sparkmax.CmdApiFactoryReset
	} else {
		frame.Header.API = sparkmax.CmdApiFactoryDefault
	}

	if command.GetBurnAfterWrite() {
		frame.Data[0] = 1
	}

	frame.Data[4] = uint8(sparkmax.ParamType_bool)

	_, err = sparkmax.SparkWriteFrame(frame)

	if err != nil {
		resp.Error = err.Error()
	}

	return &resp, err
}

func runFactoryReset(cmd *cobra.Command, args []string) {
	var err error
	req := sparkmax.FactoryResetRequest{BurnAfterWrite: !dontBurnTable, FullWipe: doFullReset}

	resp, err := FactoryReset(&req)
	if err != nil {
		fmt.Println("Error running factory reset command: ", err)
		return
	}

	if resp.GetError() != "" {
		fmt.Println("Failed to run factory reset: %s", resp.GetError())
		return
	}

	fmt.Println("Factory reset completed successfully")
}

func (s *factoryResetCommand) SparkCommandProcess(req sparkmax.RequestWire) (resp sparkmax.ResponseWire, err error) {
	r, err := Telemetry(req.GetTelemetry())
	if err != nil {
		tmp := sparkmax.RootResponse{Error: err.Error()}
		r.Root = &tmp
	}
	resp.Resp = &sparkmax.ResponseWire_Telemetry{Telemetry: r}
	return resp, err
}

func (s *factoryResetCommand) ExpectedType() string {
	return "factoryReset"
}
