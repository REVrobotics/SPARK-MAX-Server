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

// addressCmd represents the address command
var addressCmd = &cobra.Command{
	Use:   "address",
	Short: "Get or set device CAN ID",
	Long: `Get or set the device CAN ID. Run with 
no arguments to get the CAN ID. Run with a valid ID to 
set the CAN ID. The CAN ID must be a number between 1 - 62.`,
	Run:     runAddress,
	PreRun:  preRunConnect,
	PostRun: postRunDisconnect,
}

func init() {
	rootCmd.AddCommand(addressCmd)
}

func runAddress(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		req := sparkmax.GetParameterRequest{Parameter: sparkmax.ConfigParam_kCanID}
		resp, err := GetParameter(&req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get address: %s\n", err.Error())
		}
		val, err := strconv.Atoi(resp.Value)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error in address: %s\n", err.Error())
		}
		fmt.Println(int32(val))
	} else {
		req := sparkmax.SetParameterRequest{Parameter: sparkmax.ConfigParam_kCanID}
		req.Value = args[0]
		_, err := SetParameter(&req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error in address: %s\n", err.Error())
		}
	}
}
