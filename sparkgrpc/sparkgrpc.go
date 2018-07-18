package sparkgrpc

import (
	"encoding/binary"
	"fmt"
	"net"

	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	sparkusb "github.com/willtoth/USB-BLDC-TOOL/sparkusb"
)

type sparkusbServer struct{}

func (s *sparkusbServer) Connect(ctx context.Context, command *RootCommand) (*RootResponse, error) {
	var resp RootResponse
	err := sparkusb.Connect(command.Device)
	if err != nil {
		//resp.Error = err.Error()
	}
	return &resp, err
}

func (s *sparkusbServer) Disconnect(ctx context.Context, command *RootCommand) (*RootResponse, error) {
	var resp RootResponse
	err := sparkusb.Disconnect()
	if err != nil {
		//resp.Error = err.Error()
	}
	return &resp, err
}

func (s *sparkusbServer) List(ctx context.Context, command *ListRequest) (*ListResponse, error) {
	var resp ListResponse
	ports := sparkusb.ListDevices(command.All)

	for _, p := range ports {
		resp.DeviceList = append(resp.DeviceList, p.Name)
	}

	return &resp, nil
}

/*
func (s *sparkusbServer) Firmware(ctx context.Context, command *FirmwareRequest) (*FirmwareResponse, error) {

}
*/

func (s *sparkusbServer) Heartbeat(ctx context.Context, command *Heartbeat) (*RootResponse, error) {
	var resp RootResponse
	frame := sparkusb.DefaultFrame()

	frame.Header.Manufacturer = sparkusb.ManuBroadcast
	frame.Header.DeviceType = sparkusb.DevBroadcast
	frame.Header.ApiClass = 0x00
	frame.Header.ApiIndex = 0x01

	if command.Enable {
		frame.Data[0] = 1
	}

	if err := sparkusb.Write(frame); err != nil {
		//resp.Error = err.Error()
		return &resp, err
	}

	frameIn, err := sparkusb.Read()

	if frameIn.Header.ApiClass != sparkusb.ApiAcknowledge {
		err = fmt.Errorf("Expected ACK, recieved :%d", frameIn.Header.ApiClass)
	}

  //resp.Error = err.Error()
  
  fmt.Println(err)

	return &resp, nil
}

/*

func (s *sparkusbServer) Address(ctx context.Context, command *AddressRequest) (*AddressResponse, error) {

}

*/

func (s *sparkusbServer) SetParameter(ctx context.Context, command *ParameterRequest) (*ParameterResponse, error) {
	var resp ParameterResponse
	frame := sparkusb.DefaultFrame()

	frame.Header.ApiClass = sparkusb.ApiConfiguration
	frame.Header.ApiIndex = 0x00

	frame.Data[0] = uint8(command.Parameter)

	binary.LittleEndian.PutUint32(frame.Data[2:6], command.Value)

	if err := sparkusb.Write(frame); err != nil {
		//resp.Root.Error = err.Error()
		return &resp, err
	}

	frameIn, err := sparkusb.Read()

	if frameIn.Header.ApiClass != sparkusb.ApiAcknowledge {
		err = fmt.Errorf("Expected ACK, recieved :%d", frameIn.Header.ApiClass)
	}

	//resp.Root.Error = err.Error()  
  fmt.Println(err)

	return &resp, nil
}

func (s *sparkusbServer) GetParameter(ctx context.Context, command *ParameterRequest) (*ParameterResponse, error) {
	var resp ParameterResponse
	frame := sparkusb.DefaultFrame()

	frame.Header.ApiClass = sparkusb.ApiConfiguration
	frame.Header.ApiIndex = 0x01

	frame.Data[0] = uint8(command.Parameter)

	fmt.Print("Outgoing Frame: ")
	fmt.Println(frame)

	if err := sparkusb.Write(frame); err != nil {
		//resp.Root.Error = err.Error()
		fmt.Println(err.Error())
		return &resp, err
	}

	frameIn, err := sparkusb.Read()

	if frameIn.Header.ApiClass != sparkusb.ApiAcknowledge {
		err = fmt.Errorf("Expected ACK, recieved :%d", frameIn.Header.ApiClass)
	}

	fmt.Print("Incoming Frame:")
	fmt.Println(frameIn)

	resp.Value = binary.LittleEndian.Uint32(frameIn.Data[:4])

	//resp.Root.Error = err.Error()
  
  fmt.Println(resp)
  
  //resp.Root.Error = err.Error()
  fmt.Println(err)

	return &resp, nil
}

/*

func (s *sparkusbServer) ListParameters(ctx context.Context, command *ParameterListRequest) (*ParameterListResponse, error) {

}

*/

func (s *sparkusbServer) Setpoint(ctx context.Context, command *SetpointRequest) (*SetpointResponse, error) {
	var resp SetpointResponse
	frame := sparkusb.DefaultFrame()

	frame.Header.ApiClass = 0x00
	frame.Header.ApiIndex = 0x02

	if command.Setpoint < 0.001 && command.Setpoint > -0.001 {
		frame.Data[0] = 0
		frame.Data[1] = 0
		frame.Data[2] = 0
	} else {
		if command.Setpoint > 1023 {
			command.Setpoint = 1023
		} else if command.Setpoint < -1024 {
			command.Setpoint = 1024
		}

		tmparray := make([]byte, 4)
		binary.BigEndian.PutUint32(tmparray[:], uint32(command.Setpoint))

		//Copy 3 LSB of number (for some reason its way that CTRE does it)
		//Implemented here (possible) for compatibility
		copy(frame.Data[:3], tmparray[1:])
	}

	if err := sparkusb.Write(frame); err != nil {
		//resp.Root.Error = err.Error()
		return &resp, err
	}

	frameIn, err := sparkusb.Read()

	if frameIn.Header.ApiClass != sparkusb.ApiAcknowledge {
		err = fmt.Errorf("Expected ACK, recieved :%d", frameIn.Header.ApiClass)
	}
  
  //resp.Root.Error = err.Error()
  fmt.Println(err)

	return &resp, nil
}

func (s *sparkusbServer) CommandLine(ctx context.Context, command *CommandLineRequest) (*CommandLineResponse, error) {
	if false {
	}
	return &CommandLineResponse{Stderr: "Error ", Stdout: "Hello"}, nil
}

// RunServer blocks with a new server based on grpc definition
func RunServer(port uint) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	RegisterSparkusbServer(s, &sparkusbServer{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		return err
	}

	return nil
}
