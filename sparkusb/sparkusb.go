package sparkusb

import (
	"log"

	enumerator "go.bug.st/serial.v1/enumerator"
)

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
		if port.PID == "5740" && port.VID == "0483" {
			tmp = append(tmp, port)
		}
	}
	return tmp
}
