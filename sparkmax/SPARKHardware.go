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
	"fmt"
	"log"
	"time"

	serial "github.com/tarm/serial"
	enumerator "go.bug.st/serial.v1/enumerator"
)

const (
	sparkPID = "5740"
	sparkVID = "0483"
)

var localPort *serial.Port

func isSparkPID(pid string) (isSpark bool) {
	if sparkPID == pid {
		isSpark = true
	}
	return
}

/*ListDevices (all) set all to true to include connected CAN devices
* Note: 'All' parameter is not yet implemented
 */
func ListDevices(all bool) []*enumerator.PortDetails {
	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		log.Fatal(err)
	}

	tmp := ports[:0]
	for _, port := range ports {
		if isSparkPID(port.PID) && port.VID == sparkVID {
			tmp = append(tmp, port)
		}
	}
	return tmp
}

func IsConnected() bool {
	if localPort == nil {
		return false
	}
	return true
}

func GetDefaultDevice() (device string) {
	if devices := ListDevices(false); len(devices) > 0 {
		return ListDevices(false)[0].Name
	}
	return ""
}

func Connect(com string) error {
	if com == "" {
		com = GetDefaultDevice()
		if com == "" {
			return fmt.Errorf("No device found")
		}
	}

	serialConfig := &serial.Config{Name: com, Baud: 115200, ReadTimeout: time.Millisecond * 2000}
	port, err := serial.OpenPort(serialConfig)

	if err == nil {
		localPort = port
	}

	return err
}

func Disconnect() error {
	if localPort == nil {
		return nil
	}

	err := localPort.Close()
	localPort = nil
	return err
}

func Write(frame UsbFrame) error {
	if localPort == nil {
		return fmt.Errorf("Attempted write to uninitialized serial port")
	}
	writeData := SerializeUsbFrame(frame)
	_, err := localPort.Write(writeData)

	//TODO: figure out why this is neccessary on WIN32 to avoid recieving
	//1 byte frames... Maybe reading the buffer when size = 1?
	//Maybe on the read side call read with size 1
	//in a loop of 12 with an overall timeout?
	time.Sleep(time.Millisecond * 2)
	return err
}

func Read() (UsbFrame, error) {
	var frame UsbFrame
	if localPort == nil {
		return frame, fmt.Errorf("Attempted read to uninitialized serial port")
	}

	buf := make([]byte, FrameSize)
	len, err := localPort.Read(buf)
	if err != nil {
		return frame, err
	}

	if len%FrameSize != 0 {
		//return frame, fmt.Errorf("Packet frame unaligned, size: %d", len)
		return frame, nil
	}

	//TODO: Depending on frame size, parse multiple frames
	//Maybe a different function, as this would need to
	//return multiple frames
	frame, err = DeserializeUsbFrame(buf[:FrameSize])

	return frame, err
}
