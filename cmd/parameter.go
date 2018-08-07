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
	sparkusb "github.com/willtoth/USB-BLDC-TOOL/sparkusb"
)

// parameterCmd represents the parameter command
var parameterCmd = &cobra.Command{
	Use:     "parameter",
	Args:    cobra.RangeArgs(1, 2),
	Short:   "Get or set parameter",
	Run:     runParameter,
	Aliases: []string{"param"},
}

const (
	longPrefix = `The first argument is <parameter ID>, the second
optional parameter is the value for that parameter. 

Parameter ID list is as follows:

`
)

func buildLongStr() string {
	val := longPrefix
	for key := range sparkusb.ConfigParam_value {
		val += key
		val += "\n"
	}
	return val
}

func init() {
	parameterCmd.Long = buildLongStr()
	rootCmd.AddCommand(parameterCmd)
}

func runParameter(cmd *cobra.Command, args []string) {
	if val, ok := sparkusb.ConfigParam_value[args[0]]; ok {
		req := sparkusb.ParameterRequest{Parameter: sparkusb.ConfigParam(val)}

		if len(args) == 1 {
			resp, err := sparkusb.GetParameter(&req)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to get parameter %s: %s\r\n", val, err.Error())
			}
			fmt.Println(resp.Value)
		} else {
			req.Value = args[1]
			_, err := sparkusb.SetParameter(&req)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Set parameter failed for %s: %s\r\n", val, err.Error())
			}
		}
	} else {
		fmt.Fprintf(os.Stderr, "Invalid parameter: %s", args[0])
	}
}
