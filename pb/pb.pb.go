// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb.proto

package pb

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

type Request struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_96a777c211f3ed03, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type Response struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Tracing              string   `protobuf:"bytes,2,opt,name=tracing,proto3" json:"tracing,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_96a777c211f3ed03, []int{1}
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

func (m *Response) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Response) GetTracing() string {
	if m != nil {
		return m.Tracing
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "pb.Request")
	proto.RegisterType((*Response)(nil), "pb.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UuidServiceClient is the client API for UuidService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UuidServiceClient interface {
	Uuid(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type uuidServiceClient struct {
	cc *grpc.ClientConn
}

func NewUuidServiceClient(cc *grpc.ClientConn) UuidServiceClient {
	return &uuidServiceClient{cc}
}

func (c *uuidServiceClient) Uuid(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.UuidService/Uuid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UuidServiceServer is the server API for UuidService service.
type UuidServiceServer interface {
	Uuid(context.Context, *Request) (*Response, error)
}

func RegisterUuidServiceServer(s *grpc.Server, srv UuidServiceServer) {
	s.RegisterService(&_UuidService_serviceDesc, srv)
}

func _UuidService_Uuid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UuidServiceServer).Uuid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UuidService/Uuid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UuidServiceServer).Uuid(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _UuidService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UuidService",
	HandlerType: (*UuidServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Uuid",
			Handler:    _UuidService_Uuid_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb.proto",
}

func init() { proto.RegisterFile("pb.proto", fileDescriptor_pb_96a777c211f3ed03) }

var fileDescriptor_pb_96a777c211f3ed03 = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x28, 0x48, 0xd2, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x92, 0xe5, 0x62, 0x0f, 0x4a, 0x2d, 0x2c,
	0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0x29, 0xa9, 0x2c, 0x48, 0x95, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x0c, 0x02, 0xb3, 0x95, 0x2c, 0xb8, 0x38, 0x82, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53,
	0x41, 0xf2, 0xa5, 0xa5, 0x99, 0x29, 0x30, 0x79, 0x10, 0x5b, 0x48, 0x82, 0x8b, 0xbd, 0xa4, 0x28,
	0x31, 0x39, 0x33, 0x2f, 0x5d, 0x82, 0x09, 0x2c, 0x0c, 0xe3, 0x1a, 0x19, 0x70, 0x71, 0x87, 0x96,
	0x66, 0xa6, 0x04, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0x29, 0x72, 0xb1, 0x80, 0xb8, 0x42,
	0xdc, 0x7a, 0x05, 0x49, 0x7a, 0x50, 0x1b, 0xa5, 0x78, 0x20, 0x1c, 0x88, 0xf9, 0x49, 0x6c, 0x60,
	0x57, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x04, 0x0c, 0x7a, 0xe2, 0xa1, 0x00, 0x00, 0x00,
}
