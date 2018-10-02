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
	"strings"

	"github.com/spf13/cobra"
	sparkusb "github.com/willtoth/USB-BLDC-TOOL/sparkusb"
)

// neutralmodeCmd represents the neutralmode command
var neutralmodeCmd = &cobra.Command{
	Use:   "neutralmode",
	Short: "Set or get the neutral mode",
	Long: `Set or get the neutral mode, options are:

coast
brake,

This is the same as calling the command:

parameter IdleMode <x>`,
	Run:     runNeutralMode,
	Aliases: []string{"NeutralMode", "IdleMode", "idlemode"},
}

func init() {
	rootCmd.AddCommand(neutralmodeCmd)
}

func runNeutralMode(cmd *cobra.Command, args []string) {
	req := sparkusb.ParameterRequest{Parameter: sparkusb.ConfigParam_kIdleMode}
	if len(args) < 1 {
		resp, err := sparkusb.GetParameter(&req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get neutral mode: %s\n", err.Error())
		}
		idx, err := strconv.Atoi(resp.Value)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error in neutralmode: %s\n", err.Error())
		}
		fmt.Println(sparkusb.IdleMode_name[int32(idx)])
	} else {
		var idleMode sparkusb.IdleMode
		switch mt := strings.ToLower(args[0]); mt {
		case "b", "brake":
			idleMode = sparkusb.IdleMode_Brake
		case "c", "coast":
			idleMode = sparkusb.IdleMode_Coast
		}
		req.Value = strconv.FormatInt(int64(idleMode), 10)
		_, err := sparkusb.SetParameter(&req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error in motortype: %s\n", err.Error())
		}
	}
}
