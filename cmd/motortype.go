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

// motortypeCmd represents the motortype command
var motortypeCmd = &cobra.Command{
	Use:   "motortype",
	Short: "Set the motor type",
	Long: `Set the motor type, options are:
	
bldc
bdc,
brushless,
brushed

This is the same as calling the command:

parameter MotorType <x>`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		req := sparkusb.ParameterRequest{Parameter: sparkusb.ConfigParam_MotorType}
		if len(args) < 1 {
			resp, err := sparkusb.GetParameter(&req)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to get motor type: %s\n", err.Error())
			}
			idx, err := strconv.Atoi(resp.Value)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error in motortype: %s\n", err.Error())
			}
			fmt.Println(sparkusb.MotorType_name[int32(idx)])
		} else {
			var motorType sparkusb.MotorType
			switch mt := strings.ToLower(args[0]); mt {
			case "bdc", "brushed":
				motorType = sparkusb.MotorType_Brushed
			case "bldc", "brushless":
				motorType = sparkusb.MotorType_Brushless
			}
			req.Value = strconv.FormatInt(int64(motorType), 10)
			_, err := sparkusb.SetParameter(&req)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error in motortype: %s\n", err.Error())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(motortypeCmd)
}
