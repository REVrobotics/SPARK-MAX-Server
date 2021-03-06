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
	DevRobotController      = iota
	DevMotorController      = iota
	DevRelayController      = iota
	DevGyroSensor           = iota
	DevAccelerometerSensor  = iota
	DevUltrasonicSensor     = iota
	DevGearToothSensor      = iota
	DevPowerDistribution    = iota
	DevPneumaticsController = iota
	DevMiscCANDevice        = iota
	DevIOBreakout           = iota
)

const (
	ManuBroadcast   = 0
	ManuNI          = 1
	ManuLM          = 2
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
	shiftAPI          = shiftAPIIndex
	shiftDeviceID     = 0
)

const (
	bitsPacketNum    = 0xE0000000
	bitsDeviceType   = 0x1F000000
	bitsManufacturer = 0xFF0000
	bitsAPIClass     = 0xFC00
	bitsAPIIndex     = 0x3C0
	bitsAPI          = bitsAPIClass | bitsAPIIndex
	bitsDeviceID     = 0x3F
)

//This enum is unique to this controller
const (
	CmdBcastDisable      = 0x000
	CmdBcastHalt         = 0x001
	CmdBcastReset        = 0x002
	CmdBcastAssign       = 0x003
	CmdBcastHeartbeat    = 0x005
	CmdBcastFirmware     = 0x008
	CmdBcastEnum         = 0x009
	CmdBcastResume       = 0x00A
	CmdApiSetpntSet      = 0x001
	CmdApiDcSet          = 0x002
	CmdApiSpdSet         = 0x012
	CmdApiPosSet         = 0x032
	CmdApiVolSet         = 0x042
	CmdApiSmartmotionSet = 0x052
	CmdApiStat0          = 0x060
	CmdApiStat1          = 0x061
	CmdApiUsrStat0       = 0x062
	CmdApiUsrStat1       = 0x063
	CmdApiUsrStat2       = 0x064
	CmdApiUsrStat3       = 0x065
	CmdApiDrvStat        = 0x06A
	CmdApiClrFaults      = 0x06E
	CmdApiSetCfg         = 0x070
	CmdApiGetCfg         = 0x071
	CmdApiBurnFlash      = 0x072
	CmdApiSetFollower    = 0x073
	CmdApiFactoryDefault = 0x074
	CmdApiFactoryReset   = 0x075
	CmdApiNack           = 0x080
	CmdApiAck            = 0x081
	CmdApiBroadcast      = 0x090
	CmdApiHeartbeat      = 0x092
	CmdApiSync           = 0x093
	CmdApiIdQuery        = 0x094
	CmdApiIdAssign       = 0x095
	CmdApiFirmware       = 0x098
	CmdApiEnum           = 0x099
	CmdApiUSBHB          = 0x09A
	CmdApiMechPos        = 0x0A0

	ExtCmdBootloader = 0x010
)

const (
	CmdTypeStandard = 0
	CmdTypeExtended = 1
	CmdTypeRemote   = 2
)

type UsbFrameHeader struct {
	CommandType  uint32
	DeviceType   uint32
	Manufacturer uint32
	API          uint32
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

func BroadcastFrame() UsbFrame {
	var frame UsbFrame

	frame.Header.DeviceType = DevBroadcast
	frame.Header.Manufacturer = ManuBroadcast

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
	output |= header.CommandType << shiftPacketNum
	output |= header.DeviceType << shiftDeviceType
	output |= header.Manufacturer << shiftManufacturer
	output |= header.API << shiftAPI
	output |= header.DeviceID << shiftDeviceID
	return
}

func uint32ToUsbFrameHeader(input uint32) (header UsbFrameHeader) {
	header.CommandType = (input & bitsPacketNum) >> shiftPacketNum
	header.DeviceType = (input & bitsDeviceType) >> shiftDeviceType
	header.Manufacturer = (input & bitsManufacturer) >> shiftManufacturer
	header.API = (input & bitsAPI) >> shiftAPI
	header.DeviceID = (input & bitsDeviceID) >> shiftDeviceID
	return
}
