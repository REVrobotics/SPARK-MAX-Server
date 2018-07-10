package sparkusb

import (
	"fmt"
	"log"

	serial "go.bug.st/serial.v1"
	enumerator "go.bug.st/serial.v1/enumerator"
)

const (
	spark2PID   = "5740"
	sparkTNGPID = "5740"
	sparkVID    = "0483"
)

var localPort serial.Port

func isSparkPID(pid string) (isSpark bool) {
	if spark2PID == pid || sparkTNGPID == pid {
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

/*
func GetDefaultDevice() (device string) {
	return ListDevices(false)[0].Name
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

	//write(frame)
	return nil
}
*/

func Connect(com string) error {
	//Dummy parameters when using USB CDC driver
	mode := &serial.Mode{
		BaudRate: 9600,
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

func write(frame UsbFrame) error {
	if localPort == nil {
		return fmt.Errorf("Attempted write to uninitialized serial port")
	}
	_, err := localPort.Write(SerializeUsbFrame(frame))
	return err
}

func read() (UsbFrame, error) {
	var frame UsbFrame
	if localPort == nil {
		return frame, fmt.Errorf("Attempted read to uninitialized serial port")
	}

	var buf []byte
	len, err := localPort.Read(buf)
	if err != nil {
		return frame, err
	}

	if len%FrameSize != 0 {
		return frame, fmt.Errorf("Packet frame unaligned, size: %d", len)
	}

	frame, err = DeserializeUsbFrame(buf)

	return frame, err
}
