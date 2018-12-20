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

package sparkmax

import (
	"fmt"
	"log"
	"reflect"
)

var registeredCommands map[string]SparkMaxCommand

type SparkMaxCommand interface {
	SparkCommandProcess(RequestWire) (ResponseWire, error)

	ExpectedType() string
}

func RegisterCommand(cmd SparkMaxCommand) {
	if _, exists := registeredCommands[cmd.ExpectedType()]; exists {
		log.Fatal("Command request already registered, command being overwritten")
	}

	if registeredCommands == nil {
		registeredCommands = make(map[string]SparkMaxCommand)
	}

	typeName := "RequestWire_" + cmd.ExpectedType()

	registeredCommands[typeName] = cmd
}

func getType(myvar interface{}) string {
	t := reflect.TypeOf(myvar)
	if t == nil {
		return ""
	}
	if t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	}
	return t.Name()
}

func RunCommand(req RequestWire) (ResponseWire, error) {
	var err error
	var resp ResponseWire

	typename := getType(req.Req)

	//fmt.Println("running command: " + typename)

	if val, exists := registeredCommands[typename]; exists {
		resp, err = val.SparkCommandProcess(req)
	} else {
		err = fmt.Errorf("Command not found: " + typename)
	}
	return resp, err
}

func SparkWriteFrame(frame UsbFrame) (UsbFrame, error) {
	var resp UsbFrame
	var err error
	if err = Write(frame); err != nil {
		resp = DefaultFrame()
	} else {
		resp, err = Read()
	}

	return resp, err
}
