// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agentApplicationService.proto

package agentApplicationService

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

type LogonInfoRequest struct {
	// token is private key encrypted, can embed in code
	// including app_name, do we need serial_number for kicking out the invalid one?
	AppToken             string   `protobuf:"bytes,1,opt,name=app_token,json=appToken,proto3" json:"app_token,omitempty"`
	Env                  string   `protobuf:"bytes,2,opt,name=env,proto3" json:"env,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogonInfoRequest) Reset()         { *m = LogonInfoRequest{} }
func (m *LogonInfoRequest) String() string { return proto.CompactTextString(m) }
func (*LogonInfoRequest) ProtoMessage()    {}
func (*LogonInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_agentApplicationService_c170d34acfa1998f, []int{0}
}
func (m *LogonInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogonInfoRequest.Unmarshal(m, b)
}
func (m *LogonInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogonInfoRequest.Marshal(b, m, deterministic)
}
func (dst *LogonInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogonInfoRequest.Merge(dst, src)
}
func (m *LogonInfoRequest) XXX_Size() int {
	return xxx_messageInfo_LogonInfoRequest.Size(m)
}
func (m *LogonInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogonInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogonInfoRequest proto.InternalMessageInfo

func (m *LogonInfoRequest) GetAppToken() string {
	if m != nil {
		return m.AppToken
	}
	return ""
}

func (m *LogonInfoRequest) GetEnv() string {
	if m != nil {
		return m.Env
	}
	return ""
}

type LogonInfoResponse struct {
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Endpoints            string   `protobuf:"bytes,3,opt,name=endpoints,proto3" json:"endpoints,omitempty"`
	AppName              string   `protobuf:"bytes,4,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogonInfoResponse) Reset()         { *m = LogonInfoResponse{} }
func (m *LogonInfoResponse) String() string { return proto.CompactTextString(m) }
func (*LogonInfoResponse) ProtoMessage()    {}
func (*LogonInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_agentApplicationService_c170d34acfa1998f, []int{1}
}
func (m *LogonInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogonInfoResponse.Unmarshal(m, b)
}
func (m *LogonInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogonInfoResponse.Marshal(b, m, deterministic)
}
func (dst *LogonInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogonInfoResponse.Merge(dst, src)
}
func (m *LogonInfoResponse) XXX_Size() int {
	return xxx_messageInfo_LogonInfoResponse.Size(m)
}
func (m *LogonInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogonInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogonInfoResponse proto.InternalMessageInfo

func (m *LogonInfoResponse) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *LogonInfoResponse) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LogonInfoResponse) GetEndpoints() string {
	if m != nil {
		return m.Endpoints
	}
	return ""
}

func (m *LogonInfoResponse) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

type LogonError struct {
	Detail               string   `protobuf:"bytes,1,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogonError) Reset()         { *m = LogonError{} }
func (m *LogonError) String() string { return proto.CompactTextString(m) }
func (*LogonError) ProtoMessage()    {}
func (*LogonError) Descriptor() ([]byte, []int) {
	return fileDescriptor_agentApplicationService_c170d34acfa1998f, []int{2}
}
func (m *LogonError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogonError.Unmarshal(m, b)
}
func (m *LogonError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogonError.Marshal(b, m, deterministic)
}
func (dst *LogonError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogonError.Merge(dst, src)
}
func (m *LogonError) XXX_Size() int {
	return xxx_messageInfo_LogonError.Size(m)
}
func (m *LogonError) XXX_DiscardUnknown() {
	xxx_messageInfo_LogonError.DiscardUnknown(m)
}

var xxx_messageInfo_LogonError proto.InternalMessageInfo

func (m *LogonError) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

func init() {
	proto.RegisterType((*LogonInfoRequest)(nil), "agentApplicationService.LogonInfoRequest")
	proto.RegisterType((*LogonInfoResponse)(nil), "agentApplicationService.LogonInfoResponse")
	proto.RegisterType((*LogonError)(nil), "agentApplicationService.LogonError")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AgentApplicationServiceClient is the client API for AgentApplicationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentApplicationServiceClient interface {
	GetLogonInfo(ctx context.Context, in *LogonInfoRequest, opts ...grpc.CallOption) (*LogonInfoResponse, error)
}

type agentApplicationServiceClient struct {
	cc *grpc.ClientConn
}

func NewAgentApplicationServiceClient(cc *grpc.ClientConn) AgentApplicationServiceClient {
	return &agentApplicationServiceClient{cc}
}

func (c *agentApplicationServiceClient) GetLogonInfo(ctx context.Context, in *LogonInfoRequest, opts ...grpc.CallOption) (*LogonInfoResponse, error) {
	out := new(LogonInfoResponse)
	err := c.cc.Invoke(ctx, "/agentApplicationService.AgentApplicationService/getLogonInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentApplicationServiceServer is the server API for AgentApplicationService service.
type AgentApplicationServiceServer interface {
	GetLogonInfo(context.Context, *LogonInfoRequest) (*LogonInfoResponse, error)
}

func RegisterAgentApplicationServiceServer(s *grpc.Server, srv AgentApplicationServiceServer) {
	s.RegisterService(&_AgentApplicationService_serviceDesc, srv)
}

func _AgentApplicationService_GetLogonInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogonInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentApplicationServiceServer).GetLogonInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentApplicationService.AgentApplicationService/GetLogonInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentApplicationServiceServer).GetLogonInfo(ctx, req.(*LogonInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AgentApplicationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "agentApplicationService.AgentApplicationService",
	HandlerType: (*AgentApplicationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getLogonInfo",
			Handler:    _AgentApplicationService_GetLogonInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agentApplicationService.proto",
}

func init() {
	proto.RegisterFile("agentApplicationService.proto", fileDescriptor_agentApplicationService_c170d34acfa1998f)
}

var fileDescriptor_agentApplicationService_c170d34acfa1998f = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xa9, 0x2d, 0xb5, 0x19, 0x44, 0xea, 0x82, 0x36, 0x46, 0x05, 0x09, 0x1e, 0xd4, 0x43,
	0x0f, 0xfa, 0x04, 0x39, 0x78, 0x10, 0xc4, 0x43, 0xf5, 0x1e, 0xd6, 0x64, 0x0c, 0x8b, 0xcd, 0xcc,
	0xb8, 0xbb, 0xa9, 0x17, 0x1f, 0xc2, 0x77, 0xf5, 0xe6, 0x49, 0xb2, 0x0d, 0xb5, 0x88, 0x05, 0x6f,
	0xf3, 0xcf, 0x3f, 0xb3, 0x33, 0xfb, 0x0d, 0x9c, 0xe8, 0x0a, 0xc9, 0x67, 0x22, 0x73, 0x53, 0x68,
	0x6f, 0x98, 0x1e, 0xd0, 0x2e, 0x4c, 0x81, 0x53, 0xb1, 0xec, 0x59, 0x4d, 0x36, 0xd8, 0xc9, 0x7e,
	0xf0, 0xb5, 0x98, 0xbc, 0xe0, 0xba, 0x66, 0x5a, 0xd6, 0xa7, 0x19, 0x8c, 0xef, 0xb8, 0x62, 0xba,
	0xa5, 0x67, 0x9e, 0xe1, 0x6b, 0x83, 0xce, 0xab, 0x23, 0x88, 0xb4, 0x48, 0xee, 0xf9, 0x05, 0x29,
	0xee, 0x9d, 0xf6, 0xce, 0xa3, 0xd9, 0x48, 0x8b, 0x3c, 0xb6, 0x5a, 0x8d, 0xa1, 0x8f, 0xb4, 0x88,
	0xb7, 0x42, 0xba, 0x0d, 0xd3, 0x77, 0xd8, 0x5b, 0x7b, 0xc2, 0x09, 0x93, 0x43, 0xa5, 0x60, 0xd0,
	0x38, 0xb4, 0x5d, 0x7b, 0x88, 0x55, 0x02, 0x23, 0xd1, 0xce, 0xbd, 0xb1, 0x2d, 0xbb, 0xfe, 0x95,
	0x56, 0xc7, 0x10, 0x21, 0x95, 0xc2, 0x86, 0xbc, 0x8b, 0xfb, 0xc1, 0xfc, 0x49, 0xa8, 0x43, 0x68,
	0x17, 0xc8, 0x49, 0xd7, 0x18, 0x0f, 0x82, 0xb9, 0xad, 0x45, 0xee, 0x75, 0x8d, 0xe9, 0x19, 0x40,
	0x98, 0x7e, 0x63, 0x2d, 0x5b, 0x75, 0x00, 0xc3, 0x12, 0xbd, 0x36, 0xf3, 0x6e, 0x70, 0xa7, 0xae,
	0x3e, 0x7a, 0x30, 0xc9, 0xfe, 0x26, 0xa3, 0x1a, 0xd8, 0xa9, 0xd0, 0xaf, 0xbe, 0xa0, 0x2e, 0xa6,
	0x9b, 0x10, 0xff, 0x26, 0x95, 0x5c, 0xfe, 0xa7, 0x74, 0x49, 0x24, 0xdd, 0xfd, 0xfa, 0x8c, 0xd7,
	0x56, 0x7d, 0x1a, 0x86, 0x03, 0x5c, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x27, 0x2a, 0x33, 0xd0,
	0xd1, 0x01, 0x00, 0x00,
}
