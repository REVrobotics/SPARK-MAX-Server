package cmd

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"

	sparkmax "github.com/REVrobotics/SPARK-MAX-Server/sparkmax"
	"github.com/spf13/cobra"
)

func Float32FromBytes(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	float := math.Float32frombits(bits)
	return float
}

func Float32ToBytes(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)
	return bytes
}

func preRunConnect(cmd *cobra.Command, args []string) {
	if Remote == false && Persist == false {
		err := sparkmax.Connect(Device)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		/* Check firmware version and force update
		//Return the firmware version
		req := sparkmax.FirmwareRequest{}
		resp, err := Firmware(&req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed initialize: %v\r\n", err)
			return
		}

		if resp.Major < 1 || resp.Build < 371 {
			sparkmax.Disconnect()
			fmt.Fprintf(os.Stderr, """Firmware version too old, detected %s,
			requires at least %s\r\nGo to www.REVRobotics.com/software
			for the lastest version.""", resp.Version, "v1.0.371")
			os.Exit(1)
		}

		//TODO: Check for firmware files and display newest as available
		*/
	}
}

func postRunDisconnect(cmd *cobra.Command, args []string) {
	if Remote == false && Persist == false {
		sparkmax.Disconnect()
	}
}
