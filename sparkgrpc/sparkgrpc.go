package sparkgrpc

import (
	"fmt"
	"net"

	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type sparkusbServer struct{}

/*
func (s *sparkusbServer) Connect(ctx context.Context, command *RootCommand) (*RootResponse, error) {

}

func (s *sparkusbServer) Disconnect(ctx context.Context, command *RootCommand) (*RootResponse, error) {

}

func (s *sparkusbServer) List(ctx context.Context, command *ListRequest) (*ListResponse, error) {

}

func (s *sparkusbServer) Firmware(ctx context.Context, command *FirmwareRequest) (*FirmwareResponse, error) {

}

func (s *sparkusbServer) Heartbeat(ctx context.Context, command *Heartbeat) (*RootResponse, error) {

}

func (s *sparkusbServer) Address(ctx context.Context, command *AddressRequest) (*AddressResponse, error) {

}

func (s *sparkusbServer) SetParameter(ctx context.Context, command *ParameterRequest) (*ParameterResponse, error) {

}

func (s *sparkusbServer) GetParameter(ctx context.Context, command *ParameterRequest) (*ParameterResponse, error) {

}

func (s *sparkusbServer) ListParameters(ctx context.Context, command *ParameterListRequest) (*ParameterListResponse, error) {

}

func (s *sparkusbServer) Setpoint(ctx context.Context, command *SetpointRequest) (*SetpointResponse, error) {

}
*/

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
