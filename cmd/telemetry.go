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
	"encoding/binary"
	"fmt"
	"math"
	"strconv"

	sparkmax "github.com/REVrobotics/SPARK-MAX-Server/sparkmax"
	"github.com/spf13/cobra"
)

type telemetryCommand struct {
	cobra.Command
}

// setpointCmd represents the setpoint command
var telemetryCmd = &telemetryCommand{cobra.Command{
	Use:   "telemetry",
	Short: "Read or set telemetry data",
	Long: `Read or set telemetry data. Use no arguments to return
	all telemetry available. Use the telemetry name to get a specific
	telemetry data, or use the telemetry name and value to set the 
	telemetry data.
	
	This release only includes one telemetry option, which is to 
	set the sensor position.`,
	Run:     runTelemetry,
	Aliases: []string{"telem", "data"},
	PreRun:  preRunConnect,
	PostRun: postRunDisconnect,
}}

func init() {
	rootCmd.AddCommand(&telemetryCmd.Command)
	sparkmax.RegisterCommand(telemetryCmd)
}

func Telemetry(command *sparkmax.TelemetryRequest) (*sparkmax.TelemetryResponse, error) {
	var resp sparkmax.TelemetryResponse
	var err error
	frame := sparkmax.DefaultFrame()

	if command.GetData() != nil {
		//Attempt to set a telemetry parameter
		frame.Header.API = sparkmax.CmdApiMechPos + uint32(command.GetData().GetId())
		rawMsg := math.Float32bits(command.Data.Value)
		binary.LittleEndian.PutUint32(frame.Data[0:4], rawMsg)
		frame.Data[4] = uint8(sparkmax.ParamType_float32)

		_, err = sparkmax.SparkWriteFrame(frame)
	} else {

	}

	if err != nil {
		tmp := sparkmax.RootResponse{Error: err.Error()}
		resp.Root = &tmp
	}

	return &resp, err
}

func runTelemetry(cmd *cobra.Command, args []string) {
	var err error
	req := sparkmax.TelemetryRequest{}

	if len(args) == 0 {
		//request all telemetry be printed
	} else if len(args) == 1 {
		//request specific telemetry
	} else if len(args) == 2 {
		//Attempt to set telemetry
		if val, ok := sparkmax.TelemetryId_value[args[0]]; ok {
			tmp, err := strconv.ParseFloat(args[1], 32)
			if err != nil {
				fmt.Println("Error parsing telemetry value: ", err)
				return
			}
			req.Data = &sparkmax.TelemetryData{}
			req.Data.Id = sparkmax.TelemetryId(val)
			req.Data.Value = float32(tmp)
		}
	} else {
		fmt.Println("Invalid input")
		return
	}

	resp, err := Telemetry(&req)
	if err != nil {
		fmt.Println("Error running telemetry command: ", err)
		return
	}

	for _, val := range resp.GetData() {
		fmt.Printf("%s: %d\r\n", sparkmax.TelemetryId_name[int32(val.Id)], val.Value)
	}

	fmt.Println("Telemetry completed successfully")
}

func (s *telemetryCommand) SparkCommandProcess(req sparkmax.RequestWire) (resp sparkmax.ResponseWire, err error) {
	r, err := Telemetry(req.GetTelemetry())
	if err != nil {
		tmp := sparkmax.RootResponse{Error: err.Error()}
		r.Root = &tmp
	}
	resp.Resp = &sparkmax.ResponseWire_Telemetry{Telemetry: r}
	return resp, err
}

func (s *telemetryCommand) ExpectedType() string {
	return "telemetry"
}
