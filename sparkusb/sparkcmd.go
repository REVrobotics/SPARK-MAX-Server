package sparkusb

import (
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
)

func float32FromBytes(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	float := math.Float32frombits(bits)
	return float
}

func float32ToBytes(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)
	return bytes
}

func sparkCommand(frame UsbFrame) (UsbFrame, error) {
	var resp UsbFrame
	var err error
	if err = Write(frame); err != nil {
		resp = DefaultFrame()
	} else {
		resp, err = Read()
	}

	return resp, err
}

func sendHeartbeat(enable bool) error {
	frame := DefaultFrame()

	frame.Header.API = CmdApiHeartbeat

	if enable {
		frame.Data[0] = 1
	}

	_, err := sparkCommand(frame)

	return err
}

func Heartbeat(command *HeartbeatRequest) (*RootResponse, error) {
	var resp RootResponse

	err := sendHeartbeat(command.Enable)

	return &resp, err
}

/*

func Address(command *AddressRequest) (*AddressResponse, error) {

}

*/

func getParameterType(paramID ConfigParam) ParamType {
	var paramTypeList = map[int32]ParamType{
		int32(ConfigParam_kCanID):         ParamType_uint32,
		int32(ConfigParam_kInputMode):     ParamType_uint32,
		int32(ConfigParam_kMotorType):     ParamType_uint32,
		int32(ConfigParam_kCommAdvance):   ParamType_float32,
		int32(ConfigParam_kSensorType):    ParamType_uint32,
		int32(ConfigParam_kCtrlType):      ParamType_uint32,
		int32(ConfigParam_kIdleMode):      ParamType_uint32,
		int32(ConfigParam_kInputDeadband): ParamType_float32,
	}

	return paramTypeList[int32(paramID)]
}

func SetParameter(command *ParameterRequest) (*ParameterResponse, error) {
	var resp ParameterResponse
	frame := DefaultFrame()

	frame.Header.API = CmdApiSetCfg

	frame.Data[0] = uint8(command.Parameter)

	//rawMsg := frame.Data[2:6]
	var rawMsg uint32
	var err error
	resp.Type = getParameterType(command.Parameter)

	//Parse to string from raw bytes of the appropriate type
	switch resp.Type {
	case ParamType_uint32:
		tmp, err := strconv.ParseUint(command.Value, 10, 32)
		if err != nil {
			return &resp, err
		}
		rawMsg = uint32(tmp)
	case ParamType_int32:
		tmp, err := strconv.ParseInt(command.Value, 10, 32)
		if err != nil {
			return &resp, err
		}
		rawMsg = uint32(tmp)
	case ParamType_float32:
		tmp, err := strconv.ParseFloat(command.Value, 32)
		if err != nil {
			return &resp, err
		}
		rawMsg = math.Float32bits(float32(tmp))
	}
	binary.LittleEndian.PutUint32(frame.Data[2:6], rawMsg)
	frame.Data[6] = uint8(resp.Type)

	_, err = sparkCommand(frame)

	//TODO: Check response for correct type and status flag

	return &resp, err
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

	rawMsg := binary.LittleEndian.Uint32(frameIn.Data[:4])
	resp.Type = ParamType(frameIn.Data[4])

	//Parse to string from raw bytes of the appropriate type
	switch resp.Type {
	case ParamType_int32:
		resp.Value = strconv.FormatInt(int64(rawMsg), 10)
	case ParamType_uint32:
		resp.Value = strconv.FormatUint(uint64(rawMsg), 10)
	case ParamType_float32:
		rawMsgFloat := math.Float32frombits(rawMsg)
		resp.Value = strconv.FormatFloat(float64(rawMsgFloat), 'f', 6, 32)
	}

	return &resp, err
}

func BurnFlash(command *BurnRequest) (*BurnResponse, error) {
	var resp BurnResponse
	frame := DefaultFrame()

	frame.Header.API = CmdApiBurnFlash

	frame.Data[0] = 0xA3
	frame.Data[1] = 0x3A

	_, err := sparkCommand(frame)

	if err != nil {
		var tmp RootResponse
		tmp.Error = err.Error()
		resp.Root = &tmp
		resp.Verify = false
	} else {
		resp.Verify = true
	}
	return &resp, err
}

func Firmware(command *FirmwareRequest) (*FirmwareResponse, error) {
	var resp FirmwareResponse
	var err error
	var frameIn UsbFrame
	frame := BroadcastFrame()

	if command.Filename == "" {
		frame.Header.API = CmdBcastFirmware

		frameIn, err = sparkCommand(frame)

		resp.Version = fmt.Sprintf("v%d.%d.%d", frameIn.Data[0], frameIn.Data[1], uint16(frameIn.Data[2])<<8|uint16(frameIn.Data[3]))

		if frameIn.Data[4] == 1 {
			resp.Version += ", Debug build"
		}
	} else {
		//TODO: Firmware update
	}

	return &resp, err
}

/*

func ListParameters(command *ParameterListRequest) (*ParameterListResponse, error) {

}

*/

func Setpoint(command *SetpointRequest) (*SetpointResponse, error) {
	var resp SetpointResponse
	var err error
	frame := DefaultFrame()

	frame.Header.API = CmdApiDcSet

	if command.Enable {
		err = sendHeartbeat(command.Enable)
		if err != nil {
			return &resp, err
		}
	}

	if command.Setpoint < 0.001 && command.Setpoint > -0.001 {
		frame.Data[0] = 0
		frame.Data[1] = 0
		frame.Data[2] = 0
		frame.Data[3] = 0
	} else {
		//TODO: Implement a min/max based on user setting
		tmparray := float32ToBytes(command.Setpoint)

		copy(frame.Data[:4], tmparray[:])
	}

	_, err = sparkCommand(frame)

	return &resp, err
}
