// Code generated by protoc-gen-go. DO NOT EDIT.
// source: register.proto

package register

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type User struct {
	Firstname            string   `protobuf:"bytes,1,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname             string   `protobuf:"bytes,2,opt,name=lastname,proto3" json:"lastname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_1303fe8288f4efb6, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *User) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "register.User")
}

func init() {
	proto.RegisterFile("register.proto", fileDescriptor_1303fe8288f4efb6)
}

var fileDescriptor_1303fe8288f4efb6 = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x4a, 0x4d, 0xcf,
	0x2c, 0x2e, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0xa5, 0xa4,
	0xd3, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0xc1, 0xe2, 0x49, 0xa5, 0x69, 0xfa, 0xa9, 0xb9, 0x05,
	0x25, 0x95, 0x10, 0x65, 0x4a, 0x0e, 0x5c, 0x2c, 0xa1, 0xc5, 0xa9, 0x45, 0x42, 0x32, 0x5c, 0x9c,
	0x69, 0x99, 0x45, 0xc5, 0x25, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41,
	0x08, 0x01, 0x21, 0x29, 0x2e, 0x8e, 0x9c, 0x44, 0xa8, 0x24, 0x13, 0x58, 0x12, 0xce, 0x37, 0x0a,
	0xe4, 0x92, 0x72, 0x2f, 0x2a, 0x48, 0x06, 0x99, 0x12, 0x04, 0xb5, 0xd2, 0x39, 0x3f, 0xaf, 0xa4,
	0x28, 0x3f, 0x27, 0x27, 0xb5, 0x48, 0xc8, 0x98, 0x8b, 0xdd, 0xb5, 0x22, 0x35, 0xb9, 0xb4, 0x24,
	0x55, 0x88, 0x4f, 0x0f, 0xee, 0x44, 0x90, 0x62, 0x29, 0x31, 0x3d, 0x88, 0xc3, 0xf4, 0x60, 0x0e,
	0xd3, 0x73, 0x05, 0x39, 0x4c, 0x89, 0x21, 0x89, 0x0d, 0x2c, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x65, 0x70, 0x0a, 0x21, 0xd4, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GrpcUserRegisterControllerClient is the client API for GrpcUserRegisterController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GrpcUserRegisterControllerClient interface {
	Execute(ctx context.Context, in *User, opts ...grpc.CallOption) (*empty.Empty, error)
}

type grpcUserRegisterControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewGrpcUserRegisterControllerClient(cc grpc.ClientConnInterface) GrpcUserRegisterControllerClient {
	return &grpcUserRegisterControllerClient{cc}
}

func (c *grpcUserRegisterControllerClient) Execute(ctx context.Context, in *User, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/register.GrpcUserRegisterController/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GrpcUserRegisterControllerServer is the server API for GrpcUserRegisterController service.
type GrpcUserRegisterControllerServer interface {
	Execute(context.Context, *User) (*empty.Empty, error)
}

// UnimplementedGrpcUserRegisterControllerServer can be embedded to have forward compatible implementations.
type UnimplementedGrpcUserRegisterControllerServer struct {
}

func (*UnimplementedGrpcUserRegisterControllerServer) Execute(ctx context.Context, req *User) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}

func RegisterGrpcUserRegisterControllerServer(s *grpc.Server, srv GrpcUserRegisterControllerServer) {
	s.RegisterService(&_GrpcUserRegisterController_serviceDesc, srv)
}

func _GrpcUserRegisterController_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcUserRegisterControllerServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/register.GrpcUserRegisterController/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcUserRegisterControllerServer).Execute(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

var _GrpcUserRegisterController_serviceDesc = grpc.ServiceDesc{
	ServiceName: "register.GrpcUserRegisterController",
	HandlerType: (*GrpcUserRegisterControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Execute",
			Handler:    _GrpcUserRegisterController_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "register.proto",
}