package sparkusb

import (
	"encoding/binary"
	"fmt"
)

func sparkCommand(frame UsbFrame) (UsbFrame, error) {
	var resp UsbFrame
	var err error
	if err = Write(frame); err != nil {
		resp = DefaultFrame()
	} else {
		resp, err = Read()

		if resp.Header.API != CmdApiAck {
			err = fmt.Errorf("Expected ACK, recieved :%d", resp.Header.API)
		}
	}

	return resp, err
}

func Heartbeat(command *HeartbeatRequest) (*RootResponse, error) {
	var resp RootResponse
	frame := DefaultFrame()

	frame.Header.Manufacturer = ManuBroadcast
	frame.Header.DeviceType = DevBroadcast
	frame.Header.API = CmdBcastHalt

	if command.Enable {
		frame.Data[0] = 1
	}

	_, err := sparkCommand(frame)

	//resp.Error = err.Error()

	fmt.Println(err)

	return &resp, nil
}

/*

func Address(command *AddressRequest) (*AddressResponse, error) {

}

*/

func SetParameter(command *ParameterRequest) (*ParameterResponse, error) {
	var resp ParameterResponse
	frame := DefaultFrame()

	frame.Header.API = CmdApiSetCfg

	frame.Data[0] = uint8(command.Parameter)

	binary.LittleEndian.PutUint32(frame.Data[2:6], command.Value)

	_, err := sparkCommand(frame)

	//resp.Root.Error = err.Error()
	fmt.Println(err)

	return &resp, nil
}

func GetParameter(command *ParameterRequest) (*ParameterResponse, error) {
	var resp ParameterResponse
	frame := DefaultFrame()

	frame.Header.API = CmdApiGetCfg

	frame.Data[0] = uint8(command.Parameter)

	//fmt.Print("Outgoing Frame: ")
	//fmt.Println(frame)

	frameIn, err := sparkCommand(frame)

	//fmt.Print("Incoming Frame:")
	//fmt.Println(frameIn)

	resp.Value = binary.LittleEndian.Uint32(frameIn.Data[:4])

	//resp.Root.Error = err.Error()

	//fmt.Println(resp)

	//resp.Root.Error = err.Error()
	//fmt.Println(err)

	return &resp, err
}

func BurnFlash(command *RootCommand) (*RootResponse, error) {
	var resp RootResponse
	frame := DefaultFrame()

	frame.Header.API = CmdApiBurnFlash

	frame.Data[0] = 0xA3
	frame.Data[1] = 0x3A

	_, err := sparkCommand(frame)

	//fmt.Print("Incoming Frame:")
	//fmt.Println(frameIn)

	//resp.Root.Error = err.Error()

	//fmt.Println(resp)

	//resp.Root.Error = err.Error()
	//fmt.Println(err)

	return &resp, err
}

/*

func ListParameters(command *ParameterListRequest) (*ParameterListResponse, error) {

}

*/

func Setpoint(command *SetpointRequest) (*SetpointResponse, error) {
	var resp SetpointResponse
	frame := DefaultFrame()

	frame.Header.API = CmdApiDcSet

	if command.Setpoint < 0.001 && command.Setpoint > -0.001 {
		frame.Data[0] = 0
		frame.Data[1] = 0
		frame.Data[2] = 0
	} else {
		if command.Setpoint > 1023 {
			command.Setpoint = 1023
		} else if command.Setpoint < -1024 {
			command.Setpoint = 1024
		}

		tmparray := make([]byte, 4)
		binary.BigEndian.PutUint32(tmparray[:], uint32(command.Setpoint))

		//Copy 3 LSB of number (for some reason its way that CTRE does it)
		//Implemented here (possible) for compatibility
		copy(frame.Data[:3], tmparray[1:])
	}

	_, err := sparkCommand(frame)

	//resp.Root.Error = err.Error()
	fmt.Println(err)

	return &resp, nil
}
