package sparkusb

import (
	"encoding/binary"
	"math"
	"strconv"
)

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

	frame.Header.Manufacturer = ManuBroadcast
	frame.Header.DeviceType = DevBroadcast
	frame.Header.API = CmdBcastHalt

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

func getParameterType(paramID ConfigParam) string {
	var paramTypeList = map[int32]string{
		int32(ConfigParam_CanID):           "uint",
		int32(ConfigParam_InputMode):       "uint",
		int32(ConfigParam_MotorType):       "uint",
		int32(ConfigParam_CommAdv):         "float",
		int32(ConfigParam_SensorType):      "uint",
		int32(ConfigParam_CtrlType):        "uint",
		int32(ConfigParam_IdleMode):        "uint",
		int32(ConfigParam_InputDeadband):   "float",
		int32(ConfigParam_FirmwareVersion): "uint",
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

	//Parse string param to raw bytes of the appropriate type
	switch getParameterType(command.Parameter) {
	case "uint":
		tmp, err := strconv.ParseUint(command.Value, 10, 32)
		if err != nil {
			return &resp, err
		}
		rawMsg = uint32(tmp)
	case "int":
		tmp, err := strconv.ParseInt(command.Value, 10, 32)
		if err != nil {
			return &resp, err
		}
		rawMsg = uint32(tmp)
	case "float":
		tmp, err := strconv.ParseFloat(command.Value, 32)
		if err != nil {
			return &resp, err
		}
		rawMsg = math.Float32bits(float32(tmp))
	}
	binary.LittleEndian.PutUint32(frame.Data[2:6], rawMsg)

	_, err = sparkCommand(frame)

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

	//Parse to string from raw bytes of the appropriate type
	switch getParameterType(command.Parameter) {
	case "int":
		resp.Value = strconv.FormatInt(int64(rawMsg), 10)
	case "uint":
		resp.Value = strconv.FormatUint(uint64(rawMsg), 10)
	case "float":
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

	_, err = sparkCommand(frame)

	return &resp, err
}
