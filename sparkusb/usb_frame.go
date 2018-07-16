package sparkusb

import (
	"encoding/binary"
	"fmt"
)

const (
	//FrameSize for each frame
	FrameSize = 12

	//PacketSizeMax limited by CDC driver
	PacketSizeMax = 64
)

const (
	DevBroadcast            = iota
	devRobotController      = iota
	DevMotorController      = iota
	devRelayController      = iota
	devGyroSensor           = iota
	devAccelerometerSensor  = iota
	devUltrasonicSensor     = iota
	devGearToothSensor      = iota
	devPowerDistribution    = iota
	devPneumaticsController = iota
	devMiscCANDevice        = iota
	devIOBreakout           = iota
)

const (
	ManuBroadcast   = 0
	ManuNI          = 1
	ManuLM          = 2 //(TI)
	ManuDEKA        = 3
	ManuCTRE        = 4
	ManuREV         = 5
	ManuGrapple     = 6
	ManuMindSensors = 7
	ManuTeamUse     = 8
)

const (
	shiftPacketNum    = 29
	shiftDeviceType   = 24
	shiftManufacturer = 16
	shiftAPIClass     = 10
	shiftAPIIndex     = 6
	shiftDeviceID     = 0
)

const (
	bitsPacketNum    = 0xE0000000
	bitsDeviceType   = 0x1F000000
	bitsManufacturer = 0xFF0000
	bitsAPIClass     = 0xFC00
	bitsAPIIndex     = 0x3C0
	bitsDeviceID     = 0x3F
)

//This enum is unique to this controller
const (
	ApiDutyCycleControl = 0
	ApiSpeedControl     = 1
	ApiVoltageControl   = 2
	ApiPositionControl  = 3
	ApiCurrentControl   = 4
	ApiStatus           = 5
	ApiPeriodicStatus   = 6
	ApiConfiguration    = 7
	ApiAcknowledge      = 8
)

type UsbFrameHeader struct {
	PacketNum    uint32
	DeviceType   uint32
	Manufacturer uint32
	ApiClass     uint32
	ApiIndex     uint32
	DeviceID     uint32
}

// UsbFrame structure sent to device
type UsbFrame struct {
	Header UsbFrameHeader
	Data   [8]uint8
}

func DefaultFrame() UsbFrame {
	var frame UsbFrame

	frame.Header.DeviceType = DevMotorController
	frame.Header.Manufacturer = ManuREV

	return frame
}

// SerializeUsbFrame convert frame to byte array
func SerializeUsbFrame(frame UsbFrame) (out []byte) {
	out = make([]byte, 12)
	header := usbFrameHeaderToUint32(frame.Header)

	copy(out[4:], frame.Data[:])
	binary.LittleEndian.PutUint32(out[:4], header)
	return
}

// DeserializeUsbFrame convert byte array to frame
func DeserializeUsbFrame(in []byte) (UsbFrame, error) {
	var frame UsbFrame
	if len(in) != FrameSize {
		return frame, fmt.Errorf("Frame size mismatch, expected %d, was %d", FrameSize, len(in))
	}

	copy(frame.Data[:], in[4:])
	tmp := binary.LittleEndian.Uint32(in)
	frame.Header = uint32ToUsbFrameHeader(tmp)

	return frame, nil
}

func usbFrameHeaderToUint32(header UsbFrameHeader) (output uint32) {
	output |= header.PacketNum << shiftPacketNum
	output |= header.DeviceType << shiftDeviceType
	output |= header.Manufacturer << shiftManufacturer
	output |= header.ApiClass << shiftAPIClass
	output |= header.ApiIndex << shiftAPIIndex
	output |= header.DeviceID << shiftDeviceID
	return
}

func uint32ToUsbFrameHeader(input uint32) (header UsbFrameHeader) {
	header.PacketNum = (input & bitsPacketNum) >> shiftPacketNum
	header.DeviceType = (input & bitsDeviceType) >> shiftDeviceType
	header.Manufacturer = (input & bitsManufacturer) >> shiftManufacturer
	header.ApiClass = (input & bitsAPIClass) >> shiftAPIClass
	header.ApiIndex = (input & bitsAPIIndex) >> shiftAPIIndex
	header.DeviceID = (input & bitsDeviceID) >> shiftDeviceID
	return
}
