// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/verificationcode/v1/verificationcode.proto

package verificationcode

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type SendCodeRequest struct {
	// The ID of the user to send the verification code
	JuvoId               string   `protobuf:"bytes,1,opt,name=juvo_id,json=juvoId,proto3" json:"juvo_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendCodeRequest) Reset()         { *m = SendCodeRequest{} }
func (m *SendCodeRequest) String() string { return proto.CompactTextString(m) }
func (*SendCodeRequest) ProtoMessage()    {}
func (*SendCodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e687964a8cdbecd7, []int{0}
}

func (m *SendCodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendCodeRequest.Unmarshal(m, b)
}
func (m *SendCodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendCodeRequest.Marshal(b, m, deterministic)
}
func (m *SendCodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendCodeRequest.Merge(m, src)
}
func (m *SendCodeRequest) XXX_Size() int {
	return xxx_messageInfo_SendCodeRequest.Size(m)
}
func (m *SendCodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendCodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendCodeRequest proto.InternalMessageInfo

func (m *SendCodeRequest) GetJuvoId() string {
	if m != nil {
		return m.JuvoId
	}
	return ""
}

type SendCodeResponse struct {
	VerificationCode     string   `protobuf:"bytes,1,opt,name=verification_code,json=verificationCode,proto3" json:"verification_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendCodeResponse) Reset()         { *m = SendCodeResponse{} }
func (m *SendCodeResponse) String() string { return proto.CompactTextString(m) }
func (*SendCodeResponse) ProtoMessage()    {}
func (*SendCodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e687964a8cdbecd7, []int{1}
}

func (m *SendCodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendCodeResponse.Unmarshal(m, b)
}
func (m *SendCodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendCodeResponse.Marshal(b, m, deterministic)
}
func (m *SendCodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendCodeResponse.Merge(m, src)
}
func (m *SendCodeResponse) XXX_Size() int {
	return xxx_messageInfo_SendCodeResponse.Size(m)
}
func (m *SendCodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendCodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendCodeResponse proto.InternalMessageInfo

func (m *SendCodeResponse) GetVerificationCode() string {
	if m != nil {
		return m.VerificationCode
	}
	return ""
}

type VerifyCodeRequest struct {
	// The ID of the user to send the verification code
	JuvoId string `protobuf:"bytes,1,opt,name=juvo_id,json=juvoId,proto3" json:"juvo_id,omitempty"`
	// The verification code to test
	VerificationCode     string   `protobuf:"bytes,2,opt,name=verification_code,json=verificationCode,proto3" json:"verification_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyCodeRequest) Reset()         { *m = VerifyCodeRequest{} }
func (m *VerifyCodeRequest) String() string { return proto.CompactTextString(m) }
func (*VerifyCodeRequest) ProtoMessage()    {}
func (*VerifyCodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e687964a8cdbecd7, []int{2}
}

func (m *VerifyCodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyCodeRequest.Unmarshal(m, b)
}
func (m *VerifyCodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyCodeRequest.Marshal(b, m, deterministic)
}
func (m *VerifyCodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyCodeRequest.Merge(m, src)
}
func (m *VerifyCodeRequest) XXX_Size() int {
	return xxx_messageInfo_VerifyCodeRequest.Size(m)
}
func (m *VerifyCodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyCodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyCodeRequest proto.InternalMessageInfo

func (m *VerifyCodeRequest) GetJuvoId() string {
	if m != nil {
		return m.JuvoId
	}
	return ""
}

func (m *VerifyCodeRequest) GetVerificationCode() string {
	if m != nil {
		return m.VerificationCode
	}
	return ""
}

type VerifyCodeResponse struct {
	// True if the code is correct (just return 200?)
	Verified             bool     `protobuf:"varint,1,opt,name=verified,proto3" json:"verified,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyCodeResponse) Reset()         { *m = VerifyCodeResponse{} }
func (m *VerifyCodeResponse) String() string { return proto.CompactTextString(m) }
func (*VerifyCodeResponse) ProtoMessage()    {}
func (*VerifyCodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e687964a8cdbecd7, []int{3}
}

func (m *VerifyCodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyCodeResponse.Unmarshal(m, b)
}
func (m *VerifyCodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyCodeResponse.Marshal(b, m, deterministic)
}
func (m *VerifyCodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyCodeResponse.Merge(m, src)
}
func (m *VerifyCodeResponse) XXX_Size() int {
	return xxx_messageInfo_VerifyCodeResponse.Size(m)
}
func (m *VerifyCodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyCodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyCodeResponse proto.InternalMessageInfo

func (m *VerifyCodeResponse) GetVerified() bool {
	if m != nil {
		return m.Verified
	}
	return false
}

func init() {
	proto.RegisterType((*SendCodeRequest)(nil), "juvo.verification_code.v1.SendCodeRequest")
	proto.RegisterType((*SendCodeResponse)(nil), "juvo.verification_code.v1.SendCodeResponse")
	proto.RegisterType((*VerifyCodeRequest)(nil), "juvo.verification_code.v1.VerifyCodeRequest")
	proto.RegisterType((*VerifyCodeResponse)(nil), "juvo.verification_code.v1.VerifyCodeResponse")
}

func init() {
	proto.RegisterFile("api/verificationcode/v1/verificationcode.proto", fileDescriptor_e687964a8cdbecd7)
}

var fileDescriptor_e687964a8cdbecd7 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x4a, 0xc3, 0x40,
	0x14, 0x26, 0x5d, 0xd4, 0xfa, 0x10, 0xda, 0xce, 0xa6, 0x36, 0x74, 0x21, 0x01, 0xa9, 0xb6, 0x36,
	0x63, 0x75, 0xa7, 0x88, 0x58, 0x45, 0x70, 0x57, 0x2a, 0x14, 0x74, 0x53, 0x62, 0x32, 0x86, 0x91,
	0x76, 0x5e, 0x6c, 0xd2, 0x40, 0x11, 0x37, 0x7a, 0x03, 0xdd, 0x79, 0x24, 0xb7, 0xe2, 0x0d, 0x3c,
	0x88, 0xcc, 0x24, 0xda, 0x90, 0xb4, 0xda, 0xe5, 0xbc, 0xef, 0x7d, 0x3f, 0xf3, 0x65, 0x02, 0xa6,
	0xe5, 0x71, 0x1a, 0xb2, 0x31, 0xbf, 0xe5, 0xb6, 0x15, 0x70, 0x14, 0x36, 0x3a, 0x8c, 0x86, 0xed,
	0xcc, 0xcc, 0xf4, 0xc6, 0x18, 0x20, 0xa9, 0xde, 0x4d, 0x42, 0x34, 0x93, 0xe0, 0x40, 0xa1, 0x61,
	0x5b, 0xaf, 0xb9, 0x88, 0xee, 0x90, 0x51, 0xa9, 0x68, 0x09, 0x81, 0x81, 0x82, 0xfd, 0x88, 0xa8,
	0x57, 0x12, 0xa8, 0x3d, 0xe4, 0x4c, 0x04, 0x11, 0x60, 0x34, 0xa0, 0x78, 0xc9, 0x84, 0x73, 0x8a,
	0x0e, 0xeb, 0xb1, 0xfb, 0x09, 0xf3, 0x03, 0x52, 0x81, 0x15, 0x69, 0x33, 0xe0, 0xce, 0xba, 0xb6,
	0xa1, 0x6d, 0xad, 0xf6, 0xf2, 0xf2, 0x78, 0xe1, 0x18, 0xc7, 0x50, 0x9a, 0xed, 0xfa, 0x1e, 0x0a,
	0x9f, 0x91, 0x26, 0x94, 0x33, 0x71, 0x62, 0x5a, 0x29, 0x09, 0x48, 0x92, 0x71, 0x05, 0xe5, 0xbe,
	0x9c, 0x4d, 0x97, 0xb1, 0x9b, 0x2f, 0x9d, 0x5b, 0x20, 0xbd, 0x0b, 0x24, 0x29, 0x1d, 0xa7, 0xd3,
	0xa1, 0x10, 0x6d, 0xb2, 0x48, 0xbc, 0xd0, 0xfb, 0x3d, 0xef, 0x7d, 0xe6, 0xa0, 0xd4, 0x4f, 0xc9,
	0x90, 0x17, 0x0d, 0x0a, 0x3f, 0x77, 0x24, 0x0d, 0x73, 0x61, 0xdd, 0x66, 0xaa, 0x34, 0xbd, 0xb9,
	0xd4, 0x6e, 0x14, 0xcb, 0x68, 0x3d, 0x7d, 0x7c, 0xbd, 0xe6, 0xea, 0x64, 0x73, 0xde, 0xa7, 0xa6,
	0x0f, 0x71, 0x1d, 0x47, 0x8d, 0xc7, 0x03, 0x9f, 0x09, 0x87, 0xbc, 0x69, 0x00, 0xb3, 0xcb, 0x91,
	0x9d, 0x3f, 0xac, 0x32, 0xf5, 0xea, 0xad, 0x25, 0xb7, 0xe3, 0x68, 0x54, 0x45, 0xdb, 0x26, 0xf5,
	0x7f, 0xa3, 0x29, 0x70, 0xaa, 0x17, 0xdf, 0x4f, 0xd6, 0x2c, 0x8f, 0x47, 0x2e, 0x36, 0x8e, 0x3a,
	0xcf, 0x1a, 0xd4, 0x6c, 0x1c, 0x65, 0x6d, 0x63, 0xd7, 0x4e, 0x35, 0xdd, 0xfa, 0x99, 0x15, 0x58,
	0x5d, 0xf9, 0x1a, 0xbb, 0xda, 0xf5, 0xb9, 0xa2, 0xb9, 0x38, 0xb4, 0x84, 0x6b, 0xe2, 0xd8, 0xa5,
	0x2e, 0x13, 0xea, 0xa5, 0x52, 0x09, 0xd0, 0x05, 0xff, 0xcb, 0x61, 0x7a, 0x76, 0x93, 0x57, 0xa4,
	0xfd, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x44, 0x44, 0x6f, 0x62, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VerificationCodeClient is the client API for VerificationCode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VerificationCodeClient interface {
	// Send the user a verification code.  Returns xxx if the user has exceeded the limit
	SendCode(ctx context.Context, in *SendCodeRequest, opts ...grpc.CallOption) (*SendCodeResponse, error)
	// Verify that the code is correct
	VerifyCode(ctx context.Context, in *VerifyCodeRequest, opts ...grpc.CallOption) (*VerifyCodeResponse, error)
}

type verificationCodeClient struct {
	cc *grpc.ClientConn
}

func NewVerificationCodeClient(cc *grpc.ClientConn) VerificationCodeClient {
	return &verificationCodeClient{cc}
}

func (c *verificationCodeClient) SendCode(ctx context.Context, in *SendCodeRequest, opts ...grpc.CallOption) (*SendCodeResponse, error) {
	out := new(SendCodeResponse)
	err := c.cc.Invoke(ctx, "/juvo.verification_code.v1.VerificationCode/SendCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verificationCodeClient) VerifyCode(ctx context.Context, in *VerifyCodeRequest, opts ...grpc.CallOption) (*VerifyCodeResponse, error) {
	out := new(VerifyCodeResponse)
	err := c.cc.Invoke(ctx, "/juvo.verification_code.v1.VerificationCode/VerifyCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationCodeServer is the server API for VerificationCode service.
type VerificationCodeServer interface {
	// Send the user a verification code.  Returns xxx if the user has exceeded the limit
	SendCode(context.Context, *SendCodeRequest) (*SendCodeResponse, error)
	// Verify that the code is correct
	VerifyCode(context.Context, *VerifyCodeRequest) (*VerifyCodeResponse, error)
}

// UnimplementedVerificationCodeServer can be embedded to have forward compatible implementations.
type UnimplementedVerificationCodeServer struct {
}

func (*UnimplementedVerificationCodeServer) SendCode(ctx context.Context, req *SendCodeRequest) (*SendCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCode not implemented")
}
func (*UnimplementedVerificationCodeServer) VerifyCode(ctx context.Context, req *VerifyCodeRequest) (*VerifyCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyCode not implemented")
}

func RegisterVerificationCodeServer(s *grpc.Server, srv VerificationCodeServer) {
	s.RegisterService(&_VerificationCode_serviceDesc, srv)
}

func _VerificationCode_SendCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerificationCodeServer).SendCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/juvo.verification_code.v1.VerificationCode/SendCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerificationCodeServer).SendCode(ctx, req.(*SendCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VerificationCode_VerifyCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerificationCodeServer).VerifyCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/juvo.verification_code.v1.VerificationCode/VerifyCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerificationCodeServer).VerifyCode(ctx, req.(*VerifyCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VerificationCode_serviceDesc = grpc.ServiceDesc{
	ServiceName: "juvo.verification_code.v1.VerificationCode",
	HandlerType: (*VerificationCodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendCode",
			Handler:    _VerificationCode_SendCode_Handler,
		},
		{
			MethodName: "VerifyCode",
			Handler:    _VerificationCode_VerifyCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/verificationcode/v1/verificationcode.proto",
}
