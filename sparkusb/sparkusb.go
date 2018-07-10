package sparkusb

import (
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

func connect(com string) {
	mode := &serial.Mode{
		BaudRate: 9600,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}
	port, err := serial.Open(com, mode)

	if err != nil {
		log.Fatal(err)
	}

	localPort = port
}

func disconnect() {
	if localPort == nil {
		return
	}

	err := localPort.Close()
	if err != nil {
		log.Fatal(err)
	}

	localPort = nil
}
