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

//Enable flag
var enable bool

// heartbeatCmd represents the heartbeat command
var heartbeatCmd = &cobra.Command{
	Use:   "heartbeat",
	Short: "Send a heartbeat command to the controller",
	Long: `Send a heartbeat command to the controller
, use -e flag to send with enable`,
	Run: func(cmd *cobra.Command, args []string) {
		req := sparkusb.HeartbeatRequest{Enable: enable}
		_, err := sparkusb.Heartbeat(&req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Heartbeat command failed: %s", err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(heartbeatCmd)
	heartbeatCmd.Flags().BoolVarP(&enable, "enable", "e", false, "Send heartbeat with enable")
}
