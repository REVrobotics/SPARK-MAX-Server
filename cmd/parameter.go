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
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	sparkmax "github.com/willtoth/USB-BLDC-TOOL/sparkmax"
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
	for key := range sparkmax.ConfigParam_value {
		val += key
		val += "\n"
	}
	return val
}

func init() {
	parameterCmd.Long = buildLongStr()
	rootCmd.AddCommand(parameterCmd)
	sparkmax.RegisterCommand(&getParameterCmd)
	sparkmax.RegisterCommand(&setParameterCmd)
}

func getParameterType(paramID sparkmax.ConfigParam) sparkmax.ParamType {
	tmp := sparkmax.ConfigParam_name[int32(paramID)] + "_t"
	return sparkmax.ParamType(sparkmax.ConfigParamTypes_value[tmp])
}

func SetParameter(command *sparkmax.SetParameterRequest) (*sparkmax.ParameterResponse, error) {
	var resp sparkmax.ParameterResponse
	frame := sparkmax.DefaultFrame()

	frame.Header.API = sparkmax.CmdApiSetCfg

	frame.Data[0] = uint8(command.Parameter)

	//rawMsg := frame.Data[2:6]
	var rawMsg uint32
	var err error
	resp.Type = getParameterType(command.Parameter)

	//Parse to string from raw bytes of the appropriate type
	switch resp.Type {
	case sparkmax.ParamType_bool:
		fallthrough
	case sparkmax.ParamType_uint32:
		tmp, err := strconv.ParseUint(command.Value, 10, 32)
		if err != nil {
			return &resp, err
		}
		rawMsg = uint32(tmp)
	case sparkmax.ParamType_int32:
		tmp, err := strconv.ParseInt(command.Value, 10, 32)
		if err != nil {
			return &resp, err
		}
		rawMsg = uint32(tmp)
	case sparkmax.ParamType_float32:
		tmp, err := strconv.ParseFloat(command.Value, 32)
		if err != nil {
			return &resp, err
		}
		rawMsg = math.Float32bits(float32(tmp))
	}
	binary.LittleEndian.PutUint32(frame.Data[2:6], rawMsg)
	frame.Data[6] = uint8(resp.Type)

	//fmt.Print("Outgoing Frame: ")
	//fmt.Println(frame)

	frameIn, err := sparkmax.SparkWriteFrame(frame)

	//TODO: Check response for correct type and status flag
	resp.Status = sparkmax.ParamStatus(frameIn.Data[7])

	//fmt.Print("Incoming Frame:")
	//fmt.Println(frameIn)

	return &resp, err
}

func GetParameter(command *sparkmax.GetParameterRequest) (*sparkmax.ParameterResponse, error) {
	var resp sparkmax.ParameterResponse
	frame := sparkmax.DefaultFrame()

	frame.Header.API = sparkmax.CmdApiGetCfg

	frame.Data[0] = uint8(command.Parameter)

	//fmt.Print("Outgoing Frame: ")
	//fmt.Println(frame)

	frameIn, err := sparkmax.SparkWriteFrame(frame)

	//fmt.Print("Incoming Frame:")
	//fmt.Println(frameIn)

	rawMsg := binary.LittleEndian.Uint32(frameIn.Data[:4])
	resp.Type = sparkmax.ParamType(frameIn.Data[4])
	resp.Status = sparkmax.ParamStatus(frameIn.Data[5])

	//Parse to string from raw bytes of the appropriate type
	switch resp.Type {
	case sparkmax.ParamType_int32:
		resp.Value = strconv.FormatInt(int64(rawMsg), 10)
	case sparkmax.ParamType_bool:
		fallthrough
	case sparkmax.ParamType_uint32:
		resp.Value = strconv.FormatUint(uint64(rawMsg), 10)
	case sparkmax.ParamType_float32:
		rawMsgFloat := math.Float32frombits(rawMsg)
		resp.Value = strconv.FormatFloat(float64(rawMsgFloat), 'f', 6, 32)
	}

	return &resp, err
}

func runParameter(cmd *cobra.Command, args []string) {
	//TODO: Allow non-exact spelling
	if val, ok := sparkmax.ConfigParam_value[args[0]]; ok {
		if len(args) == 1 {
			req := sparkmax.GetParameterRequest{Parameter: sparkmax.ConfigParam(val)}
			resp, err := GetParameter(&req)
			if resp.Status != sparkmax.ParamStatus_paramOK {
				fmt.Fprintf(os.Stderr, "Failed to get parameter %s: %s\r\n",
					sparkmax.ConfigParam_value[args[0]],
					sparkmax.ParamStatus_name[int32(resp.Status)])
			}
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to get parameter %s: %s\r\n", val, err.Error())
			}
			fmt.Println(resp.Value)
		} else {
			req := sparkmax.SetParameterRequest{Parameter: sparkmax.ConfigParam(val)}
			req.Value = args[1]
			resp, err := SetParameter(&req)

			if resp.Status != sparkmax.ParamStatus_paramOK {
				fmt.Fprintf(os.Stderr, "Set parameter failed for %s: %s\r\n",
					val,
					sparkmax.ParamStatus_name[int32(resp.Status)])
			}

			if err != nil {
				fmt.Fprintf(os.Stderr, "Set parameter failed for %s: %s\r\n", val, err.Error())
			}
		}
	} else {
		fmt.Fprintf(os.Stderr, "Invalid parameter: %s", args[0])
	}
}

type setParameterCommand struct{}
type getParameterCommand struct{}

var getParameterCmd = getParameterCommand{}
var setParameterCmd = setParameterCommand{}

func (s *setParameterCommand) SparkCommandProcess(req sparkmax.RequestWire) (resp sparkmax.ResponseWire, err error) {
	r, err := SetParameter(req.GetSetParameter())
	if err != nil {
		tmp := sparkmax.RootResponse{Error: err.Error()}
		r.Root = &tmp
	}
	resp.Resp = &sparkmax.ResponseWire_Parameter{Parameter: r}
	return resp, err
}

func (s *setParameterCommand) ExpectedType() string {
	return "SetParameter"
}

func (g *getParameterCommand) SparkCommandProcess(req sparkmax.RequestWire) (resp sparkmax.ResponseWire, err error) {
	r, err := GetParameter(req.GetGetParameter())
	if err != nil {
		tmp := sparkmax.RootResponse{Error: err.Error()}
		r.Root = &tmp
	}
	resp.Resp = &sparkmax.ResponseWire_Parameter{Parameter: r}
	return resp, err
}

func (g *getParameterCommand) ExpectedType() string {
	return "GetParameter"
}
