// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/delivery/grpc/user/user.proto

package user

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
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
	ID                   int64                `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	TeamId               int64                `protobuf:"varint,2,opt,name=TeamId,proto3" json:"TeamId,omitempty"`
	Name                 string               `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	IsAdmin              bool                 `protobuf:"varint,4,opt,name=IsAdmin,proto3" json:"IsAdmin,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_31cef1674bc06c59, []int{0}
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

func (m *User) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *User) GetTeamId() int64 {
	if m != nil {
		return m.TeamId
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetIsAdmin() bool {
	if m != nil {
		return m.IsAdmin
	}
	return false
}

func (m *User) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *User) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type SingleRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SingleRequest) Reset()         { *m = SingleRequest{} }
func (m *SingleRequest) String() string { return proto.CompactTextString(m) }
func (*SingleRequest) ProtoMessage()    {}
func (*SingleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31cef1674bc06c59, []int{1}
}

func (m *SingleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SingleRequest.Unmarshal(m, b)
}
func (m *SingleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SingleRequest.Marshal(b, m, deterministic)
}
func (m *SingleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SingleRequest.Merge(m, src)
}
func (m *SingleRequest) XXX_Size() int {
	return xxx_messageInfo_SingleRequest.Size(m)
}
func (m *SingleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SingleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SingleRequest proto.InternalMessageInfo

func (m *SingleRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Users struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=Users,proto3" json:"Users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Users) Reset()         { *m = Users{} }
func (m *Users) String() string { return proto.CompactTextString(m) }
func (*Users) ProtoMessage()    {}
func (*Users) Descriptor() ([]byte, []int) {
	return fileDescriptor_31cef1674bc06c59, []int{2}
}

func (m *Users) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Users.Unmarshal(m, b)
}
func (m *Users) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Users.Marshal(b, m, deterministic)
}
func (m *Users) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Users.Merge(m, src)
}
func (m *Users) XXX_Size() int {
	return xxx_messageInfo_Users.Size(m)
}
func (m *Users) XXX_DiscardUnknown() {
	xxx_messageInfo_Users.DiscardUnknown(m)
}

var xxx_messageInfo_Users proto.InternalMessageInfo

func (m *Users) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*SingleRequest)(nil), "user.SingleRequest")
	proto.RegisterType((*Users)(nil), "user.Users")
}

func init() { proto.RegisterFile("app/delivery/grpc/user/user.proto", fileDescriptor_31cef1674bc06c59) }

var fileDescriptor_31cef1674bc06c59 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x50, 0xcd, 0x6a, 0xf3, 0x30,
	0x10, 0xfc, 0x94, 0x38, 0xc9, 0x97, 0x0d, 0xed, 0x41, 0x85, 0x22, 0x72, 0x89, 0xea, 0x93, 0x0b,
	0xc5, 0x86, 0xf4, 0x52, 0xe8, 0x29, 0x34, 0x50, 0x7c, 0xe9, 0x41, 0x49, 0x1e, 0xc0, 0x89, 0xb6,
	0x46, 0x10, 0xc7, 0xaa, 0x24, 0x07, 0xfa, 0x96, 0x7d, 0xa4, 0x22, 0xd9, 0xee, 0xcf, 0xa9, 0x17,
	0x31, 0x3b, 0x9a, 0x91, 0x66, 0x16, 0x6e, 0x0a, 0xad, 0x33, 0x89, 0x47, 0x75, 0x46, 0xf3, 0x9e,
	0x95, 0x46, 0x1f, 0xb2, 0xc6, 0xa2, 0x09, 0x47, 0xaa, 0x4d, 0xed, 0x6a, 0x1a, 0x79, 0x3c, 0x5f,
	0x94, 0x75, 0x5d, 0x1e, 0x31, 0x0b, 0xdc, 0xbe, 0x79, 0xcd, 0x9c, 0xaa, 0xd0, 0xba, 0xa2, 0xd2,
	0xad, 0x2c, 0xfe, 0x20, 0x10, 0xed, 0x2c, 0x1a, 0x7a, 0x09, 0x83, 0x7c, 0xcd, 0x08, 0x27, 0xc9,
	0x50, 0x0c, 0xf2, 0x35, 0xbd, 0x86, 0xf1, 0x16, 0x8b, 0x2a, 0x97, 0x6c, 0x10, 0xb8, 0x6e, 0xa2,
	0x14, 0xa2, 0x97, 0xa2, 0x42, 0x36, 0xe4, 0x24, 0x99, 0x8a, 0x80, 0x29, 0x83, 0x49, 0x6e, 0x57,
	0xb2, 0x52, 0x27, 0x16, 0x71, 0x92, 0xfc, 0x17, 0xfd, 0x48, 0x1f, 0x60, 0xba, 0xd3, 0xb2, 0x70,
	0x28, 0x57, 0x8e, 0x8d, 0x38, 0x49, 0x66, 0xcb, 0x79, 0xda, 0x66, 0x4a, 0xfb, 0x4c, 0xe9, 0xb6,
	0xcf, 0x24, 0xbe, 0xc5, 0xde, 0xf9, 0x64, 0xb0, 0x73, 0x8e, 0xff, 0x76, 0x7e, 0x89, 0xe3, 0x05,
	0x5c, 0x6c, 0xd4, 0xa9, 0x3c, 0xa2, 0xc0, 0xb7, 0x06, 0xad, 0xf3, 0xd5, 0x94, 0xec, 0xab, 0x29,
	0x19, 0xdf, 0xc2, 0xc8, 0x57, 0xb6, 0x94, 0x77, 0x80, 0x11, 0x3e, 0x4c, 0x66, 0x4b, 0x48, 0xc3,
	0xfe, 0x3c, 0x25, 0xda, 0x8b, 0xe5, 0x23, 0xcc, 0x3c, 0xd8, 0xa0, 0x39, 0xab, 0x03, 0xd2, 0x3b,
	0x98, 0x3c, 0xa3, 0x0b, 0xfb, 0xba, 0x6a, 0xc5, 0xbf, 0x7e, 0x9a, 0xff, 0x78, 0x21, 0xfe, 0xb7,
	0x1f, 0x87, 0x9c, 0xf7, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x40, 0xa7, 0x19, 0xdf, 0xae, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUser(ctx context.Context, in *SingleRequest, opts ...grpc.CallOption) (*User, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUser(ctx context.Context, in *SingleRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	GetUser(context.Context, *SingleRequest) (*User, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*SingleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/delivery/grpc/user/user.proto",
}
