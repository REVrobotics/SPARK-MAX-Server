package sparkusb

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

type usbFrame struct {
	packetNum    uint32
	deviceType   uint32
	manufacturer uint32
	apiClass     uint32
	apiIndex     uint32
	deviceID     uint32
}

func usbFrameToUint32(frame usbFrame) (output uint32) {
	output |= frame.packetNum << shiftPacketNum
	output |= frame.deviceType << shiftDeviceType
	output |= frame.manufacturer << shiftManufacturer
	output |= frame.apiClass << shiftAPIClass
	output |= frame.apiIndex << shiftAPIIndex
	output |= frame.deviceID << shiftDeviceID
	return
}

func uint32ToUsbFrame(input uint32) (frame usbFrame) {
	frame.packetNum = (input & bitsPacketNum) >> shiftPacketNum
	frame.deviceType = (input & bitsDeviceType) >> shiftDeviceType
	frame.manufacturer = (input & bitsManufacturer) >> shiftManufacturer
	frame.apiClass = (input & bitsAPIClass) >> shiftAPIClass
	frame.apiIndex = (input & bitsAPIIndex) >> shiftAPIIndex
	frame.deviceID = (input & bitsDeviceID) >> shiftDeviceID
	return
}
