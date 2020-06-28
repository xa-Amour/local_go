// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ch7/pb/string.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type StringRequest struct {
	A                    string   `protobuf:"bytes,1,opt,name=A,proto3" json:"A,omitempty"`
	B                    string   `protobuf:"bytes,2,opt,name=B,proto3" json:"B,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringRequest) Reset()         { *m = StringRequest{} }
func (m *StringRequest) String() string { return proto.CompactTextString(m) }
func (*StringRequest) ProtoMessage()    {}
func (*StringRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_18fc4e3e38082adc, []int{0}
}

func (m *StringRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringRequest.Unmarshal(m, b)
}
func (m *StringRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringRequest.Marshal(b, m, deterministic)
}
func (m *StringRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringRequest.Merge(m, src)
}
func (m *StringRequest) XXX_Size() int {
	return xxx_messageInfo_StringRequest.Size(m)
}
func (m *StringRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StringRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StringRequest proto.InternalMessageInfo

func (m *StringRequest) GetA() string {
	if m != nil {
		return m.A
	}
	return ""
}

func (m *StringRequest) GetB() string {
	if m != nil {
		return m.B
	}
	return ""
}

type StringResponse struct {
	Ret                  string   `protobuf:"bytes,1,opt,name=Ret,proto3" json:"Ret,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringResponse) Reset()         { *m = StringResponse{} }
func (m *StringResponse) String() string { return proto.CompactTextString(m) }
func (*StringResponse) ProtoMessage()    {}
func (*StringResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_18fc4e3e38082adc, []int{1}
}

func (m *StringResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringResponse.Unmarshal(m, b)
}
func (m *StringResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringResponse.Marshal(b, m, deterministic)
}
func (m *StringResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringResponse.Merge(m, src)
}
func (m *StringResponse) XXX_Size() int {
	return xxx_messageInfo_StringResponse.Size(m)
}
func (m *StringResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StringResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StringResponse proto.InternalMessageInfo

func (m *StringResponse) GetRet() string {
	if m != nil {
		return m.Ret
	}
	return ""
}

func (m *StringResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*StringRequest)(nil), "pb.StringRequest")
	proto.RegisterType((*StringResponse)(nil), "pb.StringResponse")
}

func init() { proto.RegisterFile("ch7/pb/string.proto", fileDescriptor_18fc4e3e38082adc) }

var fileDescriptor_18fc4e3e38082adc = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0xce, 0x30, 0xd7,
	0x2f, 0x48, 0xd2, 0x2f, 0x2e, 0x29, 0xca, 0xcc, 0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x2a, 0x48, 0x52, 0xd2, 0xe6, 0xe2, 0x0d, 0x06, 0x8b, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16,
	0x97, 0x08, 0xf1, 0x70, 0x31, 0x3a, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31, 0x3a, 0x82,
	0x78, 0x4e, 0x12, 0x4c, 0x10, 0x9e, 0x93, 0x92, 0x09, 0x17, 0x1f, 0x4c, 0x71, 0x71, 0x41, 0x7e,
	0x5e, 0x71, 0xaa, 0x90, 0x00, 0x17, 0x73, 0x50, 0x6a, 0x09, 0x54, 0x3d, 0x88, 0x09, 0x12, 0x49,
	0x2d, 0x2a, 0x82, 0xea, 0x01, 0x31, 0x8d, 0x8a, 0x60, 0x56, 0x04, 0xa7, 0x16, 0x95, 0x65, 0x26,
	0xa7, 0x0a, 0x19, 0x72, 0xb1, 0x39, 0xe7, 0xe7, 0x25, 0x27, 0x96, 0x08, 0x09, 0xea, 0x15, 0x24,
	0xe9, 0xa1, 0xd8, 0x2f, 0x25, 0x84, 0x2c, 0x04, 0xb1, 0x45, 0x89, 0x41, 0x48, 0x8f, 0x8b, 0xc5,
	0x25, 0x33, 0x2d, 0x0d, 0x9b, 0x06, 0x4c, 0x21, 0x25, 0x86, 0x24, 0x36, 0xb0, 0x0f, 0x8d, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x2e, 0xb0, 0x9d, 0x84, 0xf8, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StringServiceClient is the client API for StringService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StringServiceClient interface {
	Concat(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*StringResponse, error)
	Diff(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*StringRequest, error)
}

type stringServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStringServiceClient(cc grpc.ClientConnInterface) StringServiceClient {
	return &stringServiceClient{cc}
}

func (c *stringServiceClient) Concat(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*StringResponse, error) {
	out := new(StringResponse)
	err := c.cc.Invoke(ctx, "/pb.StringService/Concat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stringServiceClient) Diff(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*StringRequest, error) {
	out := new(StringRequest)
	err := c.cc.Invoke(ctx, "/pb.StringService/Diff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StringServiceServer is the server API for StringService service.
type StringServiceServer interface {
	Concat(context.Context, *StringRequest) (*StringResponse, error)
	Diff(context.Context, *StringRequest) (*StringRequest, error)
}

// UnimplementedStringServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStringServiceServer struct {
}

func (*UnimplementedStringServiceServer) Concat(ctx context.Context, req *StringRequest) (*StringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Concat not implemented")
}
func (*UnimplementedStringServiceServer) Diff(ctx context.Context, req *StringRequest) (*StringRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Diff not implemented")
}

func RegisterStringServiceServer(s *grpc.Server, srv StringServiceServer) {
	s.RegisterService(&_StringService_serviceDesc, srv)
}

func _StringService_Concat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StringServiceServer).Concat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StringService/Concat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StringServiceServer).Concat(ctx, req.(*StringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StringService_Diff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StringServiceServer).Diff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StringService/Diff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StringServiceServer).Diff(ctx, req.(*StringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StringService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.StringService",
	HandlerType: (*StringServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Concat",
			Handler:    _StringService_Concat_Handler,
		},
		{
			MethodName: "Diff",
			Handler:    _StringService_Diff_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ch7/pb/string.proto",
}
