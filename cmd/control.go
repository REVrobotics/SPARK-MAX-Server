// Copyright © 2018 REV Robotics LLC (support@revrobotics.com)
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
	sparkmax "github.com/REVrobotics/SPARK-MAX-Server/sparkmax"
)

type connectCommand struct{}
type disconnectCommand struct{}
type pingCommand struct{}

var connectCmd connectCommand
var disconnectCmd disconnectCommand
var pingCmd pingCommand

func init() {
	sparkmax.RegisterCommand(&connectCmd)
	sparkmax.RegisterCommand(&disconnectCmd)
	sparkmax.RegisterCommand(&pingCmd)
}

func (s *connectCommand) SparkCommandProcess(req sparkmax.RequestWire) (resp sparkmax.ResponseWire, err error) {
	err = sparkmax.Connect(req.GetConnect().Device)
	r := &sparkmax.ConnectResponse{Connected: sparkmax.IsConnected()}

	if err != nil {
		tmp := sparkmax.RootResponse{Error: err.Error()}
		r.Root = &tmp
	}

	resp.Resp = &sparkmax.ResponseWire_Connect{Connect: r}
	return resp, err
}

func (s *connectCommand) ExpectedType() string {
	return "Connect"
}

func (s *disconnectCommand) SparkCommandProcess(req sparkmax.RequestWire) (resp sparkmax.ResponseWire, err error) {
	err = sparkmax.Disconnect()
	r := &sparkmax.DisconnectResponse{Connected: sparkmax.IsConnected()}

	if err != nil {
		tmp := sparkmax.RootResponse{Error: err.Error()}
		r.Root = &tmp
	}

	resp.Resp = &sparkmax.ResponseWire_Disconnect{Disconnect: r}
	return resp, err
}

func (s *disconnectCommand) ExpectedType() string {
	return "Disconnect"
}

func (s *pingCommand) SparkCommandProcess(req sparkmax.RequestWire) (resp sparkmax.ResponseWire, err error) {
	r := &sparkmax.PingResponse{Connected: sparkmax.IsConnected()}
	/*
		isFWUpdating := firmwareThread.IsRunning()

		if isFWUpdating {
			r.IsUpdating = isFWUpdating
			r.UpdateStageMessage = firmwareThread.GetStatus()
			r.UpdateStagePercent = firmwareThread.GetPercent()

			fwErr := firmwareThread.GetError()
			if fwErr != nil {
				tmp := sparkmax.RootResponse{Error: fwErr.Error()}
				r.Root = &tmp
			}
		}
	*/
	resp.Resp = &sparkmax.ResponseWire_Ping{Ping: r}
	return resp, err
}

func (s *pingCommand) ExpectedType() string {
	return "Ping"
}
