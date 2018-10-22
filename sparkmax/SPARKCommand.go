package sparkmax

import (
	"fmt"
	"log"
	"reflect"
)

var registeredCommands map[string]SparkMaxCommand

type SparkMaxCommand interface {
	sparkCommandProcess(RequestWire) (ResponseWire, error)

	expectedType() string
}

func RegisterCommand(cmd SparkMaxCommand) {
	if _, exists := registeredCommands[cmd.expectedType()]; exists {
		log.Fatal("Command request already registered, command being overwritten")
	}
	registeredCommands[cmd.expectedType()] = cmd
}

func RunCommand(req RequestWire) (ResponseWire, error) {
	var err error
	var resp ResponseWire

	typename := reflect.TypeOf(req.Req).Name()

	if val, exists := registeredCommands[typename]; exists {
		resp, err = val.sparkCommandProcess(req)
	} else {
		err = fmt.Errorf("Command not found")
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
