// Copyright (c) 2020 TypeFox GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: backup.proto

package api

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

type PrepareBackupRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrepareBackupRequest) Reset()         { *m = PrepareBackupRequest{} }
func (m *PrepareBackupRequest) String() string { return proto.CompactTextString(m) }
func (*PrepareBackupRequest) ProtoMessage()    {}
func (*PrepareBackupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_65240d19de191688, []int{0}
}

func (m *PrepareBackupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrepareBackupRequest.Unmarshal(m, b)
}
func (m *PrepareBackupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrepareBackupRequest.Marshal(b, m, deterministic)
}
func (m *PrepareBackupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrepareBackupRequest.Merge(m, src)
}
func (m *PrepareBackupRequest) XXX_Size() int {
	return xxx_messageInfo_PrepareBackupRequest.Size(m)
}
func (m *PrepareBackupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrepareBackupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrepareBackupRequest proto.InternalMessageInfo

type PrepareBackupResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrepareBackupResponse) Reset()         { *m = PrepareBackupResponse{} }
func (m *PrepareBackupResponse) String() string { return proto.CompactTextString(m) }
func (*PrepareBackupResponse) ProtoMessage()    {}
func (*PrepareBackupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_65240d19de191688, []int{1}
}

func (m *PrepareBackupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrepareBackupResponse.Unmarshal(m, b)
}
func (m *PrepareBackupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrepareBackupResponse.Marshal(b, m, deterministic)
}
func (m *PrepareBackupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrepareBackupResponse.Merge(m, src)
}
func (m *PrepareBackupResponse) XXX_Size() int {
	return xxx_messageInfo_PrepareBackupResponse.Size(m)
}
func (m *PrepareBackupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrepareBackupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrepareBackupResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PrepareBackupRequest)(nil), "supervisor.PrepareBackupRequest")
	proto.RegisterType((*PrepareBackupResponse)(nil), "supervisor.PrepareBackupResponse")
}

func init() {
	proto.RegisterFile("backup.proto", fileDescriptor_65240d19de191688)
}

var fileDescriptor_65240d19de191688 = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x4a, 0x4c, 0xce,
	0x2e, 0x2d, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2a, 0x2e, 0x2d, 0x48, 0x2d, 0x2a,
	0xcb, 0x2c, 0xce, 0x2f, 0x52, 0x12, 0xe3, 0x12, 0x09, 0x28, 0x4a, 0x2d, 0x48, 0x2c, 0x4a, 0x75,
	0x02, 0x2b, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x51, 0x12, 0xe7, 0x12, 0x45, 0x13, 0x2f,
	0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x35, 0x4a, 0xe4, 0xe2, 0x85, 0x88, 0x04, 0x83, 0x8c, 0x48, 0x4e,
	0x15, 0x0a, 0xe0, 0x62, 0x87, 0xaa, 0x14, 0x52, 0xd0, 0x43, 0x98, 0xac, 0x87, 0xcd, 0x58, 0x29,
	0x45, 0x3c, 0x2a, 0x20, 0x16, 0x28, 0x31, 0x38, 0xb1, 0x46, 0x31, 0x27, 0x16, 0x64, 0x26, 0xb1,
	0x81, 0x5d, 0x6b, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x49, 0x9a, 0x1f, 0xdf, 0xbd, 0x00, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BackupServiceClient is the client API for BackupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BackupServiceClient interface {
	// Prepare prepares a workspace backup and is intended to be called by the PreStop
	// hook of the container. It should hardly ever be neccesary to call this from any
	// other point.
	Prepare(ctx context.Context, in *PrepareBackupRequest, opts ...grpc.CallOption) (*PrepareBackupResponse, error)
}

type backupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBackupServiceClient(cc grpc.ClientConnInterface) BackupServiceClient {
	return &backupServiceClient{cc}
}

func (c *backupServiceClient) Prepare(ctx context.Context, in *PrepareBackupRequest, opts ...grpc.CallOption) (*PrepareBackupResponse, error) {
	out := new(PrepareBackupResponse)
	err := c.cc.Invoke(ctx, "/supervisor.BackupService/Prepare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackupServiceServer is the server API for BackupService service.
type BackupServiceServer interface {
	// Prepare prepares a workspace backup and is intended to be called by the PreStop
	// hook of the container. It should hardly ever be neccesary to call this from any
	// other point.
	Prepare(context.Context, *PrepareBackupRequest) (*PrepareBackupResponse, error)
}

// UnimplementedBackupServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBackupServiceServer struct {
}

func (*UnimplementedBackupServiceServer) Prepare(ctx context.Context, req *PrepareBackupRequest) (*PrepareBackupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Prepare not implemented")
}

func RegisterBackupServiceServer(s *grpc.Server, srv BackupServiceServer) {
	s.RegisterService(&_BackupService_serviceDesc, srv)
}

func _BackupService_Prepare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareBackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServiceServer).Prepare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisor.BackupService/Prepare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServiceServer).Prepare(ctx, req.(*PrepareBackupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BackupService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "supervisor.BackupService",
	HandlerType: (*BackupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Prepare",
			Handler:    _BackupService_Prepare_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backup.proto",
}