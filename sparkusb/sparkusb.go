package sparkusb

import (
	"fmt"
	"log"

	serial "go.bug.st/serial.v1"
	enumerator "go.bug.st/serial.v1/enumerator"
)

const (
	spark2PID = "5740"
	sparkVID  = "0483"
)

var localPort serial.Port

func isSparkPID(pid string) (isSpark bool) {
	if spark2PID == pid {
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
	} else {
		return ""
	}

}

func RunCommand(frame UsbFrame, device string, persist bool) error {
	if IsConnected() == true {
		if persist == false {
			Disconnect()
			err := Connect(device)
			if err != nil {
				return err
			}
		}
	} else {
		err := Connect(device)
		if err != nil {
			return err
		}
	}

	if persist == false {
		defer Disconnect()
	}

	//fmt.Println(frame)

	err := Write(frame)
	if err != nil {
		return err
	}

	_, err = Read()
	if err != nil {
		return err
	}

	//fmt.Println(out)

	return nil
}

/*
func RunCommand(frame UsbFrame, device string, persist bool) error {
	if IsConnected() == true {
		if persist == false {
			Disconnect()
			err := Connect(device)
			if err != nil {
				return err
			}
		}
	} else {
		err := Connect(device)
		if err != nil {
			return err
		}
	}

	if persist == false {
		defer Disconnect()
	}

	//write(frame)
	return nil
}
*/

func Connect(com string) error {
	//Dummy parameters when using USB CDC driver
	mode := &serial.Mode{
		BaudRate: 115200,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}
	port, err := serial.Open(com, mode)

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
	return err
}

func Read() (UsbFrame, error) {
	var frame UsbFrame
	if localPort == nil {
		return frame, fmt.Errorf("Attempted read to uninitialized serial port")
	}

	buf := make([]byte, 128)
	len, err := localPort.Read(buf)
	if err != nil {
		return frame, err
	}

	if len%FrameSize != 0 {
		return frame, fmt.Errorf("Packet frame unaligned, size: %d", len)
	}

	//TODO: Depending on frame size, parse multiple frames
	//Maybe a different function, as this would need to
	//return multiple frames
	frame, err = DeserializeUsbFrame(buf[:12])

	return frame, err
}
