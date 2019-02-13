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

package cmd

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"

	sparkmax "github.com/REVrobotics/SPARK-MAX-Server/sparkmax"
	"github.com/spf13/cobra"
)

type drvStatusCommand struct {
	cobra.Command
}

// setpointCmd represents the setpoint command
var drvStatusCmd = &drvStatusCommand{cobra.Command{
	Use:   "faults",
	Short: "Read the faults, sticky faults and SPI status registers from the DRV8320",
	Long: `Read the faults, sticky faults  the two status registers 
	from the DRV8320 over SPI. The definition of each register can 
	be found in the device datasheet here: 
	http://www.ti.com/lit/ds/symlink/drv8320.pdf`,
	Run:     runDRVStatus,
	Args:    cobra.ExactArgs(0),
	Aliases: []string{"drv", "drvstatus", "stat", "status"},
	PreRun:  preRunConnect,
	PostRun: postRunDisconnect,
}}

func init() {
	rootCmd.AddCommand(&drvStatusCmd.Command)
	sparkmax.RegisterCommand(drvStatusCmd)
}

func bitLocationToMask(bitLocation int32) (mask uint16) {
	mask = 1 << uint16(bitLocation)
	return
}

func drvUintToStat0(Stat0Bits uint16, Stat0 *sparkmax.DRVStat0) {
	Stat0.VDS_LC = (Stat0Bits & bitLocationToMask(sparkmax.DRVStat0_Bits_value["VDS_LC_Bit"])) != 0
	Stat0.VDS_HC = (Stat0Bits & bitLocationToMask(sparkmax.DRVStat0_Bits_value["VDS_HC_Bit"])) != 0
	Stat0.VDS_LB = (Stat0Bits & bitLocationToMask(sparkmax.DRVStat0_Bits_value["VDS_LB_Bit"])) != 0
	Stat0.VDS_HB = (Stat0Bits & bitLocationToMask(sparkmax.DRVStat0_Bits_value["VDS_HB_Bit"])) != 0
	Stat0.VDS_LA = (Stat0Bits & bitLocationToMask(sparkmax.DRVStat0_Bits_value["VDS_LA_Bit"])) != 0
	Stat0.VDS_HA = (Stat0Bits & bitLocationToMask(sparkmax.DRVStat0_Bits_value["VDS_HA_Bit"])) != 0
	Stat0.OTSD = (Stat0Bits & bitLocationToMask(sparkmax.DRVStat0_Bits_value["OTSD_Bit"])) != 0
	Stat0.UVLO = (Stat0Bits & bitLocationToMask(sparkmax.DRVStat0_Bits_value["UVLO_Bit"])) != 0
	Stat0.GDF = (Stat0Bits & bitLocationToMask(sparkmax.DRVStat0_Bits_value["GDF_Bit"])) != 0
	Stat0.VDS_OCP = (Stat0Bits & bitLocationToMask(sparkmax.DRVStat0_Bits_value["VDS_OCP_Bit"])) != 0
	Stat0.FAULT = (Stat0Bits & bitLocationToMask(sparkmax.DRVStat0_Bits_value["FAULT_Bit"])) != 0
}

func drvUintToStat1(Stat1Bits uint16, Stat1 *sparkmax.DRVStat1) {
	Stat1.VGS_LC = (Stat1Bits & bitLocationToMask(sparkmax.DRVStat1_Bits_value["VGS_LC_Bit"])) != 0
	Stat1.VGS_HC = (Stat1Bits & bitLocationToMask(sparkmax.DRVStat1_Bits_value["VGS_HC_Bit"])) != 0
	Stat1.VGS_LB = (Stat1Bits & bitLocationToMask(sparkmax.DRVStat1_Bits_value["VGS_LB_Bit"])) != 0
	Stat1.VGS_HB = (Stat1Bits & bitLocationToMask(sparkmax.DRVStat1_Bits_value["VGS_HB_Bit"])) != 0
	Stat1.VGS_LA = (Stat1Bits & bitLocationToMask(sparkmax.DRVStat1_Bits_value["VGS_LA_Bit"])) != 0
	Stat1.VGS_HA = (Stat1Bits & bitLocationToMask(sparkmax.DRVStat1_Bits_value["VGS_HA_Bit"])) != 0
	Stat1.CPUV = (Stat1Bits & bitLocationToMask(sparkmax.DRVStat1_Bits_value["CPUV_Bit"])) != 0
	Stat1.OTW = (Stat1Bits & bitLocationToMask(sparkmax.DRVStat1_Bits_value["OTW_Bit"])) != 0
	Stat1.SC_OC = false
	Stat1.SB_OC = false
	Stat1.SA_OC = false
}

func faultUintToFaults(FaultBits uint16, Faults *sparkmax.FaultFlags) {
	Faults.Brownout = (FaultBits & bitLocationToMask(sparkmax.FaultBits_value["Brownout"])) != 0
	Faults.Overcurrent = (FaultBits & bitLocationToMask(sparkmax.FaultBits_value["Overcurrent"])) != 0
	Faults.Overvoltage = (FaultBits & bitLocationToMask(sparkmax.FaultBits_value["Overvoltage"])) != 0
	Faults.MotorFault = (FaultBits & bitLocationToMask(sparkmax.FaultBits_value["MotorFault"])) != 0
	Faults.SensorFault = (FaultBits & bitLocationToMask(sparkmax.FaultBits_value["SensorFault"])) != 0
	Faults.Stall = (FaultBits & bitLocationToMask(sparkmax.FaultBits_value["Stall"])) != 0
	Faults.EEPROMCRC = (FaultBits & bitLocationToMask(sparkmax.FaultBits_value["EEPROMCRC"])) != 0
	Faults.CANTX = (FaultBits & bitLocationToMask(sparkmax.FaultBits_value["CANTX"])) != 0
	Faults.CANRX = (FaultBits & bitLocationToMask(sparkmax.FaultBits_value["CANRX"])) != 0
	Faults.OtherFault = (FaultBits & bitLocationToMask(sparkmax.FaultBits_value["OtherFault"])) != 0
	Faults.SoftLimitFwd = (FaultBits & bitLocationToMask(sparkmax.FaultBits_value["SoftLimitFwd"])) != 0
	Faults.SoftLimitRev = (FaultBits & bitLocationToMask(sparkmax.FaultBits_value["SoftLimitRev"])) != 0
	Faults.HardLimitFwd = (FaultBits & bitLocationToMask(sparkmax.FaultBits_value["HardLimitFwd"])) != 0
	Faults.HardLimitRev = (FaultBits & bitLocationToMask(sparkmax.FaultBits_value["HardLimitRev"])) != 0
}

func drvStatus(command *sparkmax.DRVStatusRequest) (*sparkmax.DRVStatusResponse, error) {
	var resp sparkmax.DRVStatusResponse
	frame := sparkmax.DefaultFrame()

	frame.Header.API = sparkmax.CmdApiDrvStat

	frameIn, err := sparkmax.SparkWriteFrame(frame)

	if err != nil {
		var tmp sparkmax.RootResponse
		tmp.Error = err.Error()
		resp.Root = &tmp
	} else {
		Stat0Bits := binary.LittleEndian.Uint16(frameIn.Data[0:2])
		Stat1Bits := binary.LittleEndian.Uint16(frameIn.Data[2:4])
		faultBits := binary.LittleEndian.Uint16(frameIn.Data[4:6])
		stickyFaultBits := binary.LittleEndian.Uint16(frameIn.Data[6:8])
		resp.Stat0 = &sparkmax.DRVStat0{}
		resp.Stat1 = &sparkmax.DRVStat1{}
		resp.Faults = &sparkmax.FaultFlags{}
		resp.StickyFaults = &sparkmax.FaultFlags{}

		drvUintToStat0(Stat0Bits, resp.Stat0)
		drvUintToStat1(Stat1Bits, resp.Stat1)
		faultUintToFaults(faultBits, resp.Faults)
		faultUintToFaults(stickyFaultBits, resp.StickyFaults)
	}
	return &resp, err
}

func runDRVStatus(cmd *cobra.Command, args []string) {
	req := sparkmax.DRVStatusRequest{}
	resp, err := drvStatus(&req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in drvStatus flash: %s\n", err.Error())
	} else {
		jsonFaults, _ := json.MarshalIndent(resp.Faults, "", "\t")
		formattedLine := fmt.Sprintf("Faults Set:\r\n%s", string(jsonFaults))
		fmt.Print(formattedLine)

		jsonStickyFaults, _ := json.MarshalIndent(resp.StickyFaults, "", "\t")
		formattedLine = fmt.Sprintf("\r\nSticky Faults Set:\r\n%s", string(jsonStickyFaults))
		fmt.Print(formattedLine)

		jsonStat0, _ := json.MarshalIndent(resp.Stat0, "", "\t")
		formattedLine = fmt.Sprintf("\r\nFlags Set for DRV Stat0:\r\n%s", string(jsonStat0))
		fmt.Print(formattedLine)

		jsonStat1, _ := json.MarshalIndent(resp.Stat1, "", "\t")
		formattedLine = fmt.Sprintf("\r\nFlags Set for DRV Stat1:\r\n%s", string(jsonStat1))
		fmt.Print(formattedLine)
	}
}

func (s *drvStatusCommand) SparkCommandProcess(req sparkmax.RequestWire) (resp sparkmax.ResponseWire, err error) {
	r, err := drvStatus(req.GetDrvStatus())
	if err != nil {
		tmp := sparkmax.RootResponse{Error: err.Error()}
		r.Root = &tmp
	}
	resp.Resp = &sparkmax.ResponseWire_DrvStatus{DrvStatus: r}
	return resp, err
}

func (s *drvStatusCommand) ExpectedType() string {
	return "DrvStatus"
}
