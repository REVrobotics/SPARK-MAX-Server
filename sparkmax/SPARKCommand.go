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
