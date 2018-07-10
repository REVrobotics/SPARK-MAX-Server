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
	devBroadcast            = iota
	devRobotController      = iota
	devMotorController      = iota
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
	manuBroadcast   = 0
	manuNI          = 1
	manuLM          = 2 //(TI)
	manuDEKA        = 3
	manuCTRE        = 4
	manuREV         = 5
	manuGrapple     = 6
	manuMindSensors = 7
	manuTeamUse     = 8
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

type usbFrameHeader struct {
	packetNum    uint32
	deviceType   uint32
	manufacturer uint32
	apiClass     uint32
	apiIndex     uint32
	deviceID     uint32
}

// UsbFrame structure sent to device
type UsbFrame struct {
	header usbFrameHeader
	data   [8]uint8
}

func DefaultFrame() UsbFrame {
	var frame UsbFrame

	return frame
}

// SerializeUsbFrame convert frame to byte array
func SerializeUsbFrame(frame UsbFrame) (out []byte) {
	out = make([]byte, 12)
	header := usbFrameHeaderToUint32(frame.header)

	copy(frame.data[:], out[4:])
	binary.LittleEndian.PutUint32(out[:4], header)
	return
}

// DeserializeUsbFrame convert byte array to frame
func DeserializeUsbFrame(in []byte) (UsbFrame, error) {
	var frame UsbFrame
	if len(in) != FrameSize {
		return frame, fmt.Errorf("Frame size mismatch, expected %d, was %d", FrameSize, len(in))
	}

	copy(in, frame.data[:])
	tmp := binary.BigEndian.Uint32(in)
	frame.header = uint32ToUsbFrameHeader(tmp)

	return frame, nil
}

func usbFrameHeaderToUint32(header usbFrameHeader) (output uint32) {
	output |= header.packetNum << shiftPacketNum
	output |= header.deviceType << shiftDeviceType
	output |= header.manufacturer << shiftManufacturer
	output |= header.apiClass << shiftAPIClass
	output |= header.apiIndex << shiftAPIIndex
	output |= header.deviceID << shiftDeviceID
	return
}

func uint32ToUsbFrameHeader(input uint32) (header usbFrameHeader) {
	header.packetNum = (input & bitsPacketNum) >> shiftPacketNum
	header.deviceType = (input & bitsDeviceType) >> shiftDeviceType
	header.manufacturer = (input & bitsManufacturer) >> shiftManufacturer
	header.apiClass = (input & bitsAPIClass) >> shiftAPIClass
	header.apiIndex = (input & bitsAPIIndex) >> shiftAPIIndex
	header.deviceID = (input & bitsDeviceID) >> shiftDeviceID
	return
}
