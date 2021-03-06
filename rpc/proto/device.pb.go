// Code generated by protoc-gen-go. DO NOT EDIT.
// source: device.proto

package proto

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

type SignUpRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Addr                 string   `protobuf:"bytes,3,opt,name=addr" json:"addr,omitempty"`
	Capabilities         string   `protobuf:"bytes,4,opt,name=capabilities" json:"capabilities,omitempty"`
	Action               string   `protobuf:"bytes,5,opt,name=action" json:"action,omitempty"`
	Ts                   string   `protobuf:"bytes,6,opt,name=ts" json:"ts,omitempty"`
	Services             string   `protobuf:"bytes,7,opt,name=services" json:"services,omitempty"`
	Uuid                 string   `protobuf:"bytes,8,opt,name=uuid" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignUpRequest) Reset()         { *m = SignUpRequest{} }
func (m *SignUpRequest) String() string { return proto.CompactTextString(m) }
func (*SignUpRequest) ProtoMessage()    {}
func (*SignUpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_53620f3171f9fc26, []int{0}
}
func (m *SignUpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpRequest.Unmarshal(m, b)
}
func (m *SignUpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpRequest.Marshal(b, m, deterministic)
}
func (dst *SignUpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpRequest.Merge(dst, src)
}
func (m *SignUpRequest) XXX_Size() int {
	return xxx_messageInfo_SignUpRequest.Size(m)
}
func (m *SignUpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpRequest proto.InternalMessageInfo

func (m *SignUpRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SignUpRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SignUpRequest) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *SignUpRequest) GetCapabilities() string {
	if m != nil {
		return m.Capabilities
	}
	return ""
}

func (m *SignUpRequest) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *SignUpRequest) GetTs() string {
	if m != nil {
		return m.Ts
	}
	return ""
}

func (m *SignUpRequest) GetServices() string {
	if m != nil {
		return m.Services
	}
	return ""
}

func (m *SignUpRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type Response struct {
	Res                  string   `protobuf:"bytes,1,opt,name=res" json:"res,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_53620f3171f9fc26, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetRes() string {
	if m != nil {
		return m.Res
	}
	return ""
}

type DataRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Start                string   `protobuf:"bytes,2,opt,name=start" json:"start,omitempty"`
	End                  string   `protobuf:"bytes,3,opt,name=end" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataRequest) Reset()         { *m = DataRequest{} }
func (m *DataRequest) String() string { return proto.CompactTextString(m) }
func (*DataRequest) ProtoMessage()    {}
func (*DataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_53620f3171f9fc26, []int{2}
}
func (m *DataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataRequest.Unmarshal(m, b)
}
func (m *DataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataRequest.Marshal(b, m, deterministic)
}
func (dst *DataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataRequest.Merge(dst, src)
}
func (m *DataRequest) XXX_Size() int {
	return xxx_messageInfo_DataRequest.Size(m)
}
func (m *DataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DataRequest proto.InternalMessageInfo

func (m *DataRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DataRequest) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *DataRequest) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

func init() {
	proto.RegisterType((*SignUpRequest)(nil), "proto.SignUpRequest")
	proto.RegisterType((*Response)(nil), "proto.Response")
	proto.RegisterType((*DataRequest)(nil), "proto.DataRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DeviceClient is the client API for Device service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeviceClient interface {
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*Response, error)
	GetDeviceInfo(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*Response, error)
}

type deviceClient struct {
	cc *grpc.ClientConn
}

func NewDeviceClient(cc *grpc.ClientConn) DeviceClient {
	return &deviceClient{cc}
}

func (c *deviceClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Device/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceClient) GetDeviceInfo(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Device/GetDeviceInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Device service

type DeviceServer interface {
	SignUp(context.Context, *SignUpRequest) (*Response, error)
	GetDeviceInfo(context.Context, *DataRequest) (*Response, error)
}

func RegisterDeviceServer(s *grpc.Server, srv DeviceServer) {
	s.RegisterService(&_Device_serviceDesc, srv)
}

func _Device_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Device/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Device_GetDeviceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).GetDeviceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Device/GetDeviceInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).GetDeviceInfo(ctx, req.(*DataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Device_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Device",
	HandlerType: (*DeviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _Device_SignUp_Handler,
		},
		{
			MethodName: "GetDeviceInfo",
			Handler:    _Device_GetDeviceInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "device.proto",
}

func init() { proto.RegisterFile("device.proto", fileDescriptor_device_53620f3171f9fc26) }

var fileDescriptor_device_53620f3171f9fc26 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xdf, 0x4a, 0xf3, 0x40,
	0x10, 0xc5, 0xbf, 0x24, 0x4d, 0xbe, 0x74, 0x6c, 0xb5, 0x0c, 0x45, 0x96, 0xe2, 0x45, 0xc9, 0x55,
	0xaf, 0x0a, 0x2a, 0xf8, 0x04, 0x85, 0xe2, 0x6d, 0xc4, 0x07, 0xd8, 0x66, 0x47, 0x59, 0xaa, 0x9b,
	0x98, 0xd9, 0x08, 0x7d, 0x3e, 0x5f, 0x4c, 0xf6, 0x4f, 0xc4, 0x82, 0x57, 0x7b, 0xe6, 0xb7, 0x33,
	0xcc, 0x9c, 0x03, 0x33, 0x45, 0x9f, 0xba, 0xa1, 0x6d, 0xd7, 0xb7, 0xb6, 0xc5, 0xdc, 0x3f, 0xd5,
	0x57, 0x02, 0xf3, 0x27, 0xfd, 0x6a, 0x9e, 0xbb, 0x9a, 0x3e, 0x06, 0x62, 0x8b, 0x97, 0x90, 0x6a,
	0x25, 0x92, 0x75, 0xb2, 0xc9, 0xeb, 0x54, 0x2b, 0x44, 0x98, 0x18, 0xf9, 0x4e, 0x22, 0x5d, 0x27,
	0x9b, 0x69, 0xed, 0xb5, 0x63, 0x52, 0xa9, 0x5e, 0x64, 0x81, 0x39, 0x8d, 0x15, 0xcc, 0x1a, 0xd9,
	0xc9, 0x83, 0x7e, 0xd3, 0x56, 0x13, 0x8b, 0x89, 0xff, 0x3b, 0x63, 0x78, 0x0d, 0x85, 0x6c, 0xac,
	0x6e, 0x8d, 0xc8, 0xfd, 0x6f, 0xac, 0xdc, 0x4e, 0xcb, 0xa2, 0xf0, 0x2c, 0xb5, 0x8c, 0x2b, 0x28,
	0x99, 0x7a, 0x77, 0x2d, 0x8b, 0xff, 0x9e, 0xfe, 0xd4, 0x6e, 0xf7, 0x30, 0x68, 0x25, 0xca, 0xb0,
	0xdb, 0xe9, 0xea, 0x06, 0xca, 0x9a, 0xb8, 0x6b, 0x0d, 0x13, 0x2e, 0x20, 0xeb, 0x89, 0xbd, 0x81,
	0x69, 0xed, 0x64, 0xb5, 0x87, 0x8b, 0x9d, 0xb4, 0x72, 0x34, 0xb8, 0x80, 0xec, 0x48, 0xa7, 0xb1,
	0xe1, 0x48, 0x27, 0x5c, 0x42, 0xce, 0x56, 0xf6, 0x36, 0x7a, 0x0c, 0x85, 0xeb, 0x23, 0xa3, 0xa2,
	0x47, 0x27, 0xef, 0x18, 0x8a, 0x9d, 0xcf, 0x10, 0x6f, 0xa1, 0x08, 0xa9, 0xe1, 0x32, 0xe4, 0xb9,
	0x3d, 0x0b, 0x71, 0x75, 0x15, 0xe9, 0x78, 0x55, 0xf5, 0x0f, 0x1f, 0x60, 0xbe, 0x27, 0x1b, 0xe6,
	0x1f, 0xcd, 0x4b, 0x8b, 0x18, 0x7b, 0x7e, 0xdd, 0xf6, 0xc7, 0xdc, 0xa1, 0xf0, 0xe4, 0xfe, 0x3b,
	0x00, 0x00, 0xff, 0xff, 0x88, 0x19, 0x5e, 0x65, 0xbf, 0x01, 0x00, 0x00,
}
