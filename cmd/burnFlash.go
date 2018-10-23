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

type burnFlashCommand struct {
	cobra.Command
}

// setpointCmd represents the setpoint command
var burnFlashCmd = &burnFlashCommand{cobra.Command{
	Use:   "burnflash",
	Short: "Make all parameter changes perminent by burning them into flash",
	Long: `Make all parameter changes perminent by burning them into flash. Only
	call this method when needed as this will re-write the entire parameter table,
	and without wear-leveling the flash can only handle 10,000 cycles.`,
	Run:     runBurnFlash,
	Args:    cobra.ExactArgs(0),
	Aliases: []string{"burn"},
}}

func init() {
	rootCmd.AddCommand(&burnFlashCmd.Command)
	sparkmax.RegisterCommand(burnFlashCmd)
}

func burnFlash(command *sparkmax.BurnRequest) (*sparkmax.BurnResponse, error) {
	var resp sparkmax.BurnResponse
	frame := sparkmax.DefaultFrame()

	frame.Header.API = sparkmax.CmdApiBurnFlash

	frame.Data[0] = 0xA3
	frame.Data[1] = 0x3A

	_, err := sparkmax.SparkWriteFrame(frame)

	if err != nil {
		var tmp sparkmax.RootResponse
		tmp.Error = err.Error()
		resp.Root = &tmp
		resp.Verify = false
	} else {
		resp.Verify = true
	}
	return &resp, err
}

func runBurnFlash(cmd *cobra.Command, args []string) {
	req := sparkmax.BurnRequest{Verify: true}
	_, err := burnFlash(&req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in burn flash: %s\n", err.Error())
	}
}

func (s *burnFlashCommand) SparkCommandProcess(req sparkmax.RequestWire) (resp sparkmax.ResponseWire, err error) {
	r, err := burnFlash(req.GetBurn())
	if err != nil {
		tmp := sparkmax.RootResponse{Error: err.Error()}
		r.Root = &tmp
	}
	resp.Resp = &sparkmax.ResponseWire_Burn{Burn: r}
	return resp, err
}

func (s *burnFlashCommand) ExpectedType() string {
	return "Burn"
}
