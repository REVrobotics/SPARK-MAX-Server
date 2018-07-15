// Code generated by protoc-gen-go. DO NOT EDIT.
// source: commands.proto

/*
Package sparkgrpc is a generated protocol buffer package.

It is generated from these files:
	commands.proto

It has these top-level messages:
	CommandLineRequest
	CommandLineResponse
	RootCommand
	RootResponse
	ListRequest
	ListResponse
	FirmwareRequest
	FirmwareResponse
	Heartbeat
	AddressRequest
	AddressResponse
	ParameterRequest
	ParameterResponse
	ParameterListRequest
	ParameterListResponse
	SetpointRequest
	SetpointResponse
*/
package sparkgrpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CommandLineRequest struct {
	Stdin string `protobuf:"bytes,1,opt,name=stdin" json:"stdin,omitempty"`
}

func (m *CommandLineRequest) Reset()                    { *m = CommandLineRequest{} }
func (m *CommandLineRequest) String() string            { return proto.CompactTextString(m) }
func (*CommandLineRequest) ProtoMessage()               {}
func (*CommandLineRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CommandLineRequest) GetStdin() string {
	if m != nil {
		return m.Stdin
	}
	return ""
}

type CommandLineResponse struct {
	Stderr string `protobuf:"bytes,1,opt,name=stderr" json:"stderr,omitempty"`
	Stdout string `protobuf:"bytes,2,opt,name=stdout" json:"stdout,omitempty"`
}

func (m *CommandLineResponse) Reset()                    { *m = CommandLineResponse{} }
func (m *CommandLineResponse) String() string            { return proto.CompactTextString(m) }
func (*CommandLineResponse) ProtoMessage()               {}
func (*CommandLineResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CommandLineResponse) GetStderr() string {
	if m != nil {
		return m.Stderr
	}
	return ""
}

func (m *CommandLineResponse) GetStdout() string {
	if m != nil {
		return m.Stdout
	}
	return ""
}

type RootCommand struct {
	Device    string `protobuf:"bytes,1,opt,name=device" json:"device,omitempty"`
	Keepalive bool   `protobuf:"varint,2,opt,name=keepalive" json:"keepalive,omitempty"`
	Help      bool   `protobuf:"varint,3,opt,name=help" json:"help,omitempty"`
}

func (m *RootCommand) Reset()                    { *m = RootCommand{} }
func (m *RootCommand) String() string            { return proto.CompactTextString(m) }
func (*RootCommand) ProtoMessage()               {}
func (*RootCommand) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RootCommand) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *RootCommand) GetKeepalive() bool {
	if m != nil {
		return m.Keepalive
	}
	return false
}

func (m *RootCommand) GetHelp() bool {
	if m != nil {
		return m.Help
	}
	return false
}

type RootResponse struct {
	HelpString string `protobuf:"bytes,1,opt,name=helpString" json:"helpString,omitempty"`
	Error      string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *RootResponse) Reset()                    { *m = RootResponse{} }
func (m *RootResponse) String() string            { return proto.CompactTextString(m) }
func (*RootResponse) ProtoMessage()               {}
func (*RootResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RootResponse) GetHelpString() string {
	if m != nil {
		return m.HelpString
	}
	return ""
}

func (m *RootResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type ListRequest struct {
	Root *RootCommand `protobuf:"bytes,1,opt,name=root" json:"root,omitempty"`
	All  bool         `protobuf:"varint,2,opt,name=all" json:"all,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ListRequest) GetRoot() *RootCommand {
	if m != nil {
		return m.Root
	}
	return nil
}

func (m *ListRequest) GetAll() bool {
	if m != nil {
		return m.All
	}
	return false
}

type ListResponse struct {
	DeviceList    []string      `protobuf:"bytes,1,rep,name=deviceList" json:"deviceList,omitempty"`
	DeviceDetails []string      `protobuf:"bytes,2,rep,name=deviceDetails" json:"deviceDetails,omitempty"`
	Root          *RootResponse `protobuf:"bytes,3,opt,name=root" json:"root,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ListResponse) GetDeviceList() []string {
	if m != nil {
		return m.DeviceList
	}
	return nil
}

func (m *ListResponse) GetDeviceDetails() []string {
	if m != nil {
		return m.DeviceDetails
	}
	return nil
}

func (m *ListResponse) GetRoot() *RootResponse {
	if m != nil {
		return m.Root
	}
	return nil
}

type FirmwareRequest struct {
	Root     *RootCommand `protobuf:"bytes,1,opt,name=root" json:"root,omitempty"`
	Filename string       `protobuf:"bytes,2,opt,name=filename" json:"filename,omitempty"`
}

func (m *FirmwareRequest) Reset()                    { *m = FirmwareRequest{} }
func (m *FirmwareRequest) String() string            { return proto.CompactTextString(m) }
func (*FirmwareRequest) ProtoMessage()               {}
func (*FirmwareRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *FirmwareRequest) GetRoot() *RootCommand {
	if m != nil {
		return m.Root
	}
	return nil
}

func (m *FirmwareRequest) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

type FirmwareResponse struct {
	Version string        `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	Root    *RootResponse `protobuf:"bytes,3,opt,name=root" json:"root,omitempty"`
}

func (m *FirmwareResponse) Reset()                    { *m = FirmwareResponse{} }
func (m *FirmwareResponse) String() string            { return proto.CompactTextString(m) }
func (*FirmwareResponse) ProtoMessage()               {}
func (*FirmwareResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *FirmwareResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *FirmwareResponse) GetRoot() *RootResponse {
	if m != nil {
		return m.Root
	}
	return nil
}

type Heartbeat struct {
	Root *RootCommand `protobuf:"bytes,1,opt,name=root" json:"root,omitempty"`
	Ok   bool         `protobuf:"varint,2,opt,name=ok" json:"ok,omitempty"`
}

func (m *Heartbeat) Reset()                    { *m = Heartbeat{} }
func (m *Heartbeat) String() string            { return proto.CompactTextString(m) }
func (*Heartbeat) ProtoMessage()               {}
func (*Heartbeat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Heartbeat) GetRoot() *RootCommand {
	if m != nil {
		return m.Root
	}
	return nil
}

func (m *Heartbeat) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type AddressRequest struct {
	Root    *RootCommand `protobuf:"bytes,1,opt,name=root" json:"root,omitempty"`
	Address uint32       `protobuf:"varint,2,opt,name=address" json:"address,omitempty"`
}

func (m *AddressRequest) Reset()                    { *m = AddressRequest{} }
func (m *AddressRequest) String() string            { return proto.CompactTextString(m) }
func (*AddressRequest) ProtoMessage()               {}
func (*AddressRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *AddressRequest) GetRoot() *RootCommand {
	if m != nil {
		return m.Root
	}
	return nil
}

func (m *AddressRequest) GetAddress() uint32 {
	if m != nil {
		return m.Address
	}
	return 0
}

type AddressResponse struct {
	CurrentAddress  uint32        `protobuf:"varint,1,opt,name=currentAddress" json:"currentAddress,omitempty"`
	PreviousAddress uint32        `protobuf:"varint,2,opt,name=previousAddress" json:"previousAddress,omitempty"`
	Root            *RootResponse `protobuf:"bytes,3,opt,name=root" json:"root,omitempty"`
}

func (m *AddressResponse) Reset()                    { *m = AddressResponse{} }
func (m *AddressResponse) String() string            { return proto.CompactTextString(m) }
func (*AddressResponse) ProtoMessage()               {}
func (*AddressResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *AddressResponse) GetCurrentAddress() uint32 {
	if m != nil {
		return m.CurrentAddress
	}
	return 0
}

func (m *AddressResponse) GetPreviousAddress() uint32 {
	if m != nil {
		return m.PreviousAddress
	}
	return 0
}

func (m *AddressResponse) GetRoot() *RootResponse {
	if m != nil {
		return m.Root
	}
	return nil
}

type ParameterRequest struct {
	Root      *RootCommand `protobuf:"bytes,1,opt,name=root" json:"root,omitempty"`
	Parameter string       `protobuf:"bytes,2,opt,name=parameter" json:"parameter,omitempty"`
	Value     []byte       `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *ParameterRequest) Reset()                    { *m = ParameterRequest{} }
func (m *ParameterRequest) String() string            { return proto.CompactTextString(m) }
func (*ParameterRequest) ProtoMessage()               {}
func (*ParameterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ParameterRequest) GetRoot() *RootCommand {
	if m != nil {
		return m.Root
	}
	return nil
}

func (m *ParameterRequest) GetParameter() string {
	if m != nil {
		return m.Parameter
	}
	return ""
}

func (m *ParameterRequest) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type ParameterResponse struct {
	OldParameter string        `protobuf:"bytes,1,opt,name=oldParameter" json:"oldParameter,omitempty"`
	NewParameter string        `protobuf:"bytes,2,opt,name=newParameter" json:"newParameter,omitempty"`
	Root         *RootResponse `protobuf:"bytes,3,opt,name=root" json:"root,omitempty"`
}

func (m *ParameterResponse) Reset()                    { *m = ParameterResponse{} }
func (m *ParameterResponse) String() string            { return proto.CompactTextString(m) }
func (*ParameterResponse) ProtoMessage()               {}
func (*ParameterResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *ParameterResponse) GetOldParameter() string {
	if m != nil {
		return m.OldParameter
	}
	return ""
}

func (m *ParameterResponse) GetNewParameter() string {
	if m != nil {
		return m.NewParameter
	}
	return ""
}

func (m *ParameterResponse) GetRoot() *RootResponse {
	if m != nil {
		return m.Root
	}
	return nil
}

type ParameterListRequest struct {
	Root *RootCommand `protobuf:"bytes,1,opt,name=root" json:"root,omitempty"`
}

func (m *ParameterListRequest) Reset()                    { *m = ParameterListRequest{} }
func (m *ParameterListRequest) String() string            { return proto.CompactTextString(m) }
func (*ParameterListRequest) ProtoMessage()               {}
func (*ParameterListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *ParameterListRequest) GetRoot() *RootCommand {
	if m != nil {
		return m.Root
	}
	return nil
}

type ParameterListResponse struct {
	Parameter []string      `protobuf:"bytes,1,rep,name=parameter" json:"parameter,omitempty"`
	Root      *RootResponse `protobuf:"bytes,3,opt,name=root" json:"root,omitempty"`
}

func (m *ParameterListResponse) Reset()                    { *m = ParameterListResponse{} }
func (m *ParameterListResponse) String() string            { return proto.CompactTextString(m) }
func (*ParameterListResponse) ProtoMessage()               {}
func (*ParameterListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *ParameterListResponse) GetParameter() []string {
	if m != nil {
		return m.Parameter
	}
	return nil
}

func (m *ParameterListResponse) GetRoot() *RootResponse {
	if m != nil {
		return m.Root
	}
	return nil
}

type SetpointRequest struct {
	Root     *RootCommand `protobuf:"bytes,1,opt,name=root" json:"root,omitempty"`
	Setpoint float32      `protobuf:"fixed32,2,opt,name=setpoint" json:"setpoint,omitempty"`
}

func (m *SetpointRequest) Reset()                    { *m = SetpointRequest{} }
func (m *SetpointRequest) String() string            { return proto.CompactTextString(m) }
func (*SetpointRequest) ProtoMessage()               {}
func (*SetpointRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *SetpointRequest) GetRoot() *RootCommand {
	if m != nil {
		return m.Root
	}
	return nil
}

func (m *SetpointRequest) GetSetpoint() float32 {
	if m != nil {
		return m.Setpoint
	}
	return 0
}

type SetpointResponse struct {
	Setpoint  float32       `protobuf:"fixed32,1,opt,name=setpoint" json:"setpoint,omitempty"`
	IsRunning bool          `protobuf:"varint,2,opt,name=isRunning" json:"isRunning,omitempty"`
	Root      *RootResponse `protobuf:"bytes,3,opt,name=root" json:"root,omitempty"`
}

func (m *SetpointResponse) Reset()                    { *m = SetpointResponse{} }
func (m *SetpointResponse) String() string            { return proto.CompactTextString(m) }
func (*SetpointResponse) ProtoMessage()               {}
func (*SetpointResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *SetpointResponse) GetSetpoint() float32 {
	if m != nil {
		return m.Setpoint
	}
	return 0
}

func (m *SetpointResponse) GetIsRunning() bool {
	if m != nil {
		return m.IsRunning
	}
	return false
}

func (m *SetpointResponse) GetRoot() *RootResponse {
	if m != nil {
		return m.Root
	}
	return nil
}

func init() {
	proto.RegisterType((*CommandLineRequest)(nil), "sparkgrpc.commandLineRequest")
	proto.RegisterType((*CommandLineResponse)(nil), "sparkgrpc.commandLineResponse")
	proto.RegisterType((*RootCommand)(nil), "sparkgrpc.rootCommand")
	proto.RegisterType((*RootResponse)(nil), "sparkgrpc.rootResponse")
	proto.RegisterType((*ListRequest)(nil), "sparkgrpc.listRequest")
	proto.RegisterType((*ListResponse)(nil), "sparkgrpc.listResponse")
	proto.RegisterType((*FirmwareRequest)(nil), "sparkgrpc.firmwareRequest")
	proto.RegisterType((*FirmwareResponse)(nil), "sparkgrpc.firmwareResponse")
	proto.RegisterType((*Heartbeat)(nil), "sparkgrpc.heartbeat")
	proto.RegisterType((*AddressRequest)(nil), "sparkgrpc.addressRequest")
	proto.RegisterType((*AddressResponse)(nil), "sparkgrpc.addressResponse")
	proto.RegisterType((*ParameterRequest)(nil), "sparkgrpc.parameterRequest")
	proto.RegisterType((*ParameterResponse)(nil), "sparkgrpc.parameterResponse")
	proto.RegisterType((*ParameterListRequest)(nil), "sparkgrpc.parameterListRequest")
	proto.RegisterType((*ParameterListResponse)(nil), "sparkgrpc.parameterListResponse")
	proto.RegisterType((*SetpointRequest)(nil), "sparkgrpc.setpointRequest")
	proto.RegisterType((*SetpointResponse)(nil), "sparkgrpc.setpointResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Sparkusb service

type SparkusbClient interface {
	Connect(ctx context.Context, in *RootCommand, opts ...grpc.CallOption) (*RootResponse, error)
	Disconnect(ctx context.Context, in *RootCommand, opts ...grpc.CallOption) (*RootResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	CommandLine(ctx context.Context, in *CommandLineRequest, opts ...grpc.CallOption) (*CommandLineResponse, error)
}

type sparkusbClient struct {
	cc *grpc.ClientConn
}

func NewSparkusbClient(cc *grpc.ClientConn) SparkusbClient {
	return &sparkusbClient{cc}
}

func (c *sparkusbClient) Connect(ctx context.Context, in *RootCommand, opts ...grpc.CallOption) (*RootResponse, error) {
	out := new(RootResponse)
	err := grpc.Invoke(ctx, "/sparkgrpc.sparkusb/Connect", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sparkusbClient) Disconnect(ctx context.Context, in *RootCommand, opts ...grpc.CallOption) (*RootResponse, error) {
	out := new(RootResponse)
	err := grpc.Invoke(ctx, "/sparkgrpc.sparkusb/Disconnect", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sparkusbClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/sparkgrpc.sparkusb/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sparkusbClient) CommandLine(ctx context.Context, in *CommandLineRequest, opts ...grpc.CallOption) (*CommandLineResponse, error) {
	out := new(CommandLineResponse)
	err := grpc.Invoke(ctx, "/sparkgrpc.sparkusb/CommandLine", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Sparkusb service

type SparkusbServer interface {
	Connect(context.Context, *RootCommand) (*RootResponse, error)
	Disconnect(context.Context, *RootCommand) (*RootResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	CommandLine(context.Context, *CommandLineRequest) (*CommandLineResponse, error)
}

func RegisterSparkusbServer(s *grpc.Server, srv SparkusbServer) {
	s.RegisterService(&_Sparkusb_serviceDesc, srv)
}

func _Sparkusb_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RootCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SparkusbServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sparkgrpc.sparkusb/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SparkusbServer).Connect(ctx, req.(*RootCommand))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sparkusb_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RootCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SparkusbServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sparkgrpc.sparkusb/Disconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SparkusbServer).Disconnect(ctx, req.(*RootCommand))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sparkusb_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SparkusbServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sparkgrpc.sparkusb/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SparkusbServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sparkusb_CommandLine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandLineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SparkusbServer).CommandLine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sparkgrpc.sparkusb/CommandLine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SparkusbServer).CommandLine(ctx, req.(*CommandLineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Sparkusb_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sparkgrpc.sparkusb",
	HandlerType: (*SparkusbServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _Sparkusb_Connect_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _Sparkusb_Disconnect_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Sparkusb_List_Handler,
		},
		{
			MethodName: "CommandLine",
			Handler:    _Sparkusb_CommandLine_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "commands.proto",
}

func init() { proto.RegisterFile("commands.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 624 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0x5e, 0xda, 0xb1, 0xad, 0xaf, 0x5d, 0x57, 0xcc, 0x18, 0xd5, 0x34, 0x26, 0x64, 0x21, 0x34,
	0x0d, 0x69, 0x87, 0x71, 0x42, 0x42, 0x42, 0x63, 0x43, 0x1c, 0x98, 0xd0, 0x14, 0x24, 0xd0, 0x8e,
	0x6e, 0xf2, 0xb6, 0x59, 0x4d, 0xed, 0x60, 0x3b, 0x9d, 0x38, 0x72, 0xe3, 0xc8, 0x81, 0x3f, 0x18,
	0xd9, 0x71, 0x7e, 0x15, 0xed, 0x90, 0xdc, 0xf2, 0xbe, 0x7c, 0xef, 0xf3, 0xf7, 0xbd, 0xc4, 0x36,
	0x8c, 0x23, 0xb9, 0x58, 0x30, 0x11, 0xeb, 0x93, 0x54, 0x49, 0x23, 0xc9, 0x40, 0xa7, 0x4c, 0xcd,
	0x6f, 0x55, 0x1a, 0xd1, 0x63, 0x20, 0xfe, 0xe5, 0x25, 0x17, 0x18, 0xe2, 0x8f, 0x0c, 0xb5, 0x21,
	0xbb, 0xf0, 0x48, 0x9b, 0x98, 0x8b, 0x69, 0xf0, 0x22, 0x38, 0x1a, 0x84, 0x79, 0x41, 0x3f, 0xc2,
	0x93, 0x06, 0x57, 0xa7, 0x52, 0x68, 0x24, 0x7b, 0xb0, 0xa1, 0x4d, 0x8c, 0x4a, 0x79, 0xb6, 0xaf,
	0x3c, 0x2e, 0x33, 0x33, 0xed, 0x95, 0xb8, 0xcc, 0x0c, 0xfd, 0x0e, 0x43, 0x25, 0xa5, 0x39, 0xcf,
	0xa5, 0x2c, 0x2d, 0xc6, 0x25, 0x8f, 0xb0, 0x68, 0xcf, 0x2b, 0x72, 0x00, 0x83, 0x39, 0x62, 0xca,
	0x12, 0xbe, 0x44, 0xa7, 0xb0, 0x15, 0x56, 0x00, 0x21, 0xb0, 0x7e, 0x87, 0x49, 0x3a, 0xed, 0xbb,
	0x17, 0xee, 0x99, 0x5e, 0xc0, 0xc8, 0x0a, 0x97, 0xc6, 0x0e, 0x01, 0x2c, 0xfe, 0xd5, 0x28, 0x2e,
	0x6e, 0xbd, 0x7a, 0x0d, 0xb1, 0x29, 0x51, 0x29, 0xa9, 0xbc, 0xbf, 0xbc, 0xa0, 0x9f, 0x61, 0x98,
	0x70, 0x6d, 0x8a, 0x51, 0x1c, 0xc3, 0xba, 0x15, 0x75, 0xed, 0xc3, 0xd3, 0xbd, 0x93, 0x72, 0x74,
	0x27, 0xb5, 0x10, 0xa1, 0xe3, 0x90, 0x09, 0xf4, 0x59, 0x92, 0x78, 0xb3, 0xf6, 0x91, 0xfe, 0x0a,
	0x60, 0x94, 0xab, 0x55, 0x9e, 0xf2, 0x7c, 0x97, 0x5c, 0x5b, 0xd1, 0xbe, 0xf5, 0x54, 0x21, 0xe4,
	0x25, 0x6c, 0xe7, 0xd5, 0x05, 0x1a, 0xc6, 0x13, 0x3d, 0xed, 0x39, 0x4a, 0x13, 0x24, 0xaf, 0xbd,
	0xa9, 0xbe, 0x33, 0xf5, 0x6c, 0xc5, 0x54, 0xb1, 0x58, 0xee, 0x8a, 0x5e, 0xc3, 0xce, 0x0d, 0x57,
	0x8b, 0x7b, 0xa6, 0xb0, 0x4b, 0xa8, 0x7d, 0xd8, 0xba, 0xe1, 0x09, 0x0a, 0xb6, 0x40, 0x3f, 0xa8,
	0xb2, 0xa6, 0xd7, 0x30, 0xa9, 0xa4, 0x7d, 0xc2, 0x29, 0x6c, 0x2e, 0x51, 0x69, 0x2e, 0x8b, 0xbf,
	0xa7, 0x28, 0xdb, 0xb9, 0xfe, 0x04, 0x83, 0x3b, 0x64, 0xca, 0xcc, 0x90, 0xb5, 0xf3, 0x3b, 0x86,
	0x9e, 0x9c, 0xfb, 0x6f, 0xd0, 0x93, 0x73, 0xfa, 0x0d, 0xc6, 0x2c, 0x8e, 0x15, 0x6a, 0xdd, 0x25,
	0xfd, 0x14, 0x36, 0x7d, 0xb7, 0x93, 0xdc, 0x0e, 0x8b, 0x92, 0xfe, 0x09, 0x60, 0xa7, 0x14, 0xf6,
	0xd9, 0x5f, 0xc1, 0x38, 0xca, 0x94, 0x42, 0x61, 0xce, 0x7c, 0x53, 0xe0, 0x9a, 0x56, 0x50, 0x72,
	0x04, 0x3b, 0xa9, 0xc2, 0x25, 0x97, 0x99, 0x3e, 0x6b, 0xa8, 0xaf, 0xc2, 0xed, 0x66, 0xa6, 0x60,
	0x92, 0x32, 0xc5, 0x16, 0x68, 0x50, 0x75, 0x09, 0x7b, 0x00, 0x83, 0xb2, 0xdf, 0x7f, 0xeb, 0x0a,
	0xb0, 0xdb, 0x65, 0xc9, 0x92, 0x0c, 0x9d, 0x97, 0x51, 0x98, 0x17, 0xf4, 0x77, 0x00, 0x8f, 0x6b,
	0x8b, 0xfa, 0x41, 0x50, 0x18, 0xc9, 0x24, 0xbe, 0x2a, 0xc5, 0xf2, 0x3f, 0xa1, 0x81, 0x59, 0x8e,
	0xc0, 0xfb, 0xab, 0x95, 0x05, 0x1b, 0x58, 0xbb, 0xf8, 0x1f, 0x60, 0xb7, 0x74, 0x72, 0xd9, 0x6d,
	0x0b, 0xd3, 0x19, 0x3c, 0x5d, 0xd1, 0xf0, 0x89, 0x1a, 0xb3, 0xc9, 0xf7, 0x6d, 0x6d, 0x36, 0x6d,
	0x37, 0xa4, 0x46, 0x93, 0x4a, 0x2e, 0x4c, 0xc7, 0x0d, 0x59, 0xb4, 0xbb, 0x99, 0xf5, 0xc2, 0xb2,
	0xa6, 0x3f, 0x61, 0x52, 0x49, 0x7b, 0xe7, 0x75, 0x7e, 0xd0, 0xe4, 0xdb, 0x54, 0x5c, 0x87, 0x99,
	0x10, 0xf6, 0x84, 0xf4, 0x87, 0x6c, 0x09, 0xb4, 0x4a, 0x75, 0xfa, 0xb7, 0x07, 0x5b, 0x8e, 0x90,
	0xe9, 0x19, 0x79, 0x07, 0x9b, 0xe7, 0x52, 0x08, 0x8c, 0x0c, 0x79, 0x20, 0xcc, 0xfe, 0x43, 0x72,
	0x74, 0x8d, 0xbc, 0x07, 0xb8, 0xe0, 0x3a, 0xea, 0x2e, 0xf0, 0x16, 0xd6, 0xdd, 0x69, 0x5a, 0x6f,
	0xad, 0x1d, 0xea, 0x8d, 0xd6, 0xfa, 0xf1, 0x4c, 0xd7, 0xc8, 0x17, 0x18, 0x9e, 0x57, 0x97, 0x1c,
	0x79, 0x5e, 0x63, 0xfe, 0x7f, 0x51, 0xee, 0x1f, 0x3e, 0xf4, 0xba, 0xd0, 0x9b, 0x6d, 0xb8, 0x2b,
	0xf7, 0xcd, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x2c, 0x99, 0x9d, 0x84, 0x07, 0x00, 0x00,
}
