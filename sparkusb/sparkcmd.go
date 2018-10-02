package sparkusb

import (
	"encoding/binary"
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
		int32(ConfigParam_CanID):           ParamType_uint32,
		int32(ConfigParam_InputMode):       ParamType_uint32,
		int32(ConfigParam_MotorType):       ParamType_uint32,
		int32(ConfigParam_CommAdv):         ParamType_float32,
		int32(ConfigParam_SensorType):      ParamType_uint32,
		int32(ConfigParam_CtrlType):        ParamType_uint32,
		int32(ConfigParam_IdleMode):        ParamType_uint32,
		int32(ConfigParam_InputDeadband):   ParamType_float32,
		int32(ConfigParam_FirmwareVersion): ParamType_uint32,
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

func BurnFlash(command *RootCommand) (*RootResponse, error) {
	var resp RootResponse
	frame := DefaultFrame()

	frame.Header.API = CmdApiBurnFlash

	frame.Data[0] = 0xA3
	frame.Data[1] = 0x3A

	_, err := sparkCommand(frame)

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
