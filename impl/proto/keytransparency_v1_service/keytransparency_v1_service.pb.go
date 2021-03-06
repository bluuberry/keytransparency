// Code generated by protoc-gen-go.
// source: keytransparency_v1_service.proto
// DO NOT EDIT!

/*
Package keytransparency_v1_service is a generated protocol buffer package.

Key Transparency Service

The Key Transparency Service API consists of a map of user names to public
keys. Each user name also has a history of public keys that have been
associated with it.

It is generated from these files:
	keytransparency_v1_service.proto

It has these top-level messages:
*/
package keytransparency_v1_service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import keytransparency_v1_types "github.com/google/keytransparency/core/proto/keytransparency_v1_types"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for KeyTransparencyService service

type KeyTransparencyServiceClient interface {
	// GetEntry returns a user's entry in the Merkle Tree.
	//
	// Entries contain signed commitments to a profile, which is also returned.
	GetEntry(ctx context.Context, in *keytransparency_v1_types.GetEntryRequest, opts ...grpc.CallOption) (*keytransparency_v1_types.GetEntryResponse, error)
	// ListEntryHistory returns a list of historic GetEntry values.
	//
	// Clients verify their account history by observing correct values for their
	// account over time.
	ListEntryHistory(ctx context.Context, in *keytransparency_v1_types.ListEntryHistoryRequest, opts ...grpc.CallOption) (*keytransparency_v1_types.ListEntryHistoryResponse, error)
	// UpdateEntry updates a user's profile.
	//
	// Returns the current user profile.
	// Clients must retry until this function returns a proof containing the desired value.
	UpdateEntry(ctx context.Context, in *keytransparency_v1_types.UpdateEntryRequest, opts ...grpc.CallOption) (*keytransparency_v1_types.UpdateEntryResponse, error)
}

type keyTransparencyServiceClient struct {
	cc *grpc.ClientConn
}

func NewKeyTransparencyServiceClient(cc *grpc.ClientConn) KeyTransparencyServiceClient {
	return &keyTransparencyServiceClient{cc}
}

func (c *keyTransparencyServiceClient) GetEntry(ctx context.Context, in *keytransparency_v1_types.GetEntryRequest, opts ...grpc.CallOption) (*keytransparency_v1_types.GetEntryResponse, error) {
	out := new(keytransparency_v1_types.GetEntryResponse)
	err := grpc.Invoke(ctx, "/keytransparency.v1.service.KeyTransparencyService/GetEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyTransparencyServiceClient) ListEntryHistory(ctx context.Context, in *keytransparency_v1_types.ListEntryHistoryRequest, opts ...grpc.CallOption) (*keytransparency_v1_types.ListEntryHistoryResponse, error) {
	out := new(keytransparency_v1_types.ListEntryHistoryResponse)
	err := grpc.Invoke(ctx, "/keytransparency.v1.service.KeyTransparencyService/ListEntryHistory", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyTransparencyServiceClient) UpdateEntry(ctx context.Context, in *keytransparency_v1_types.UpdateEntryRequest, opts ...grpc.CallOption) (*keytransparency_v1_types.UpdateEntryResponse, error) {
	out := new(keytransparency_v1_types.UpdateEntryResponse)
	err := grpc.Invoke(ctx, "/keytransparency.v1.service.KeyTransparencyService/UpdateEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for KeyTransparencyService service

type KeyTransparencyServiceServer interface {
	// GetEntry returns a user's entry in the Merkle Tree.
	//
	// Entries contain signed commitments to a profile, which is also returned.
	GetEntry(context.Context, *keytransparency_v1_types.GetEntryRequest) (*keytransparency_v1_types.GetEntryResponse, error)
	// ListEntryHistory returns a list of historic GetEntry values.
	//
	// Clients verify their account history by observing correct values for their
	// account over time.
	ListEntryHistory(context.Context, *keytransparency_v1_types.ListEntryHistoryRequest) (*keytransparency_v1_types.ListEntryHistoryResponse, error)
	// UpdateEntry updates a user's profile.
	//
	// Returns the current user profile.
	// Clients must retry until this function returns a proof containing the desired value.
	UpdateEntry(context.Context, *keytransparency_v1_types.UpdateEntryRequest) (*keytransparency_v1_types.UpdateEntryResponse, error)
}

func RegisterKeyTransparencyServiceServer(s *grpc.Server, srv KeyTransparencyServiceServer) {
	s.RegisterService(&_KeyTransparencyService_serviceDesc, srv)
}

func _KeyTransparencyService_GetEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(keytransparency_v1_types.GetEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyTransparencyServiceServer).GetEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keytransparency.v1.service.KeyTransparencyService/GetEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyTransparencyServiceServer).GetEntry(ctx, req.(*keytransparency_v1_types.GetEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyTransparencyService_ListEntryHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(keytransparency_v1_types.ListEntryHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyTransparencyServiceServer).ListEntryHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keytransparency.v1.service.KeyTransparencyService/ListEntryHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyTransparencyServiceServer).ListEntryHistory(ctx, req.(*keytransparency_v1_types.ListEntryHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyTransparencyService_UpdateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(keytransparency_v1_types.UpdateEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyTransparencyServiceServer).UpdateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keytransparency.v1.service.KeyTransparencyService/UpdateEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyTransparencyServiceServer).UpdateEntry(ctx, req.(*keytransparency_v1_types.UpdateEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeyTransparencyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "keytransparency.v1.service.KeyTransparencyService",
	HandlerType: (*KeyTransparencyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEntry",
			Handler:    _KeyTransparencyService_GetEntry_Handler,
		},
		{
			MethodName: "ListEntryHistory",
			Handler:    _KeyTransparencyService_ListEntryHistory_Handler,
		},
		{
			MethodName: "UpdateEntry",
			Handler:    _KeyTransparencyService_UpdateEntry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "keytransparency_v1_service.proto",
}

func init() { proto.RegisterFile("keytransparency_v1_service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x91, 0xbf, 0x4a, 0xc4, 0x40,
	0x10, 0xc6, 0x39, 0x05, 0x91, 0x68, 0x21, 0x2b, 0x5a, 0xe4, 0x14, 0xc4, 0xab, 0x4e, 0x34, 0x4b,
	0xce, 0xce, 0x5e, 0x14, 0xb4, 0xf2, 0x4f, 0x1d, 0xf6, 0x92, 0x21, 0xb7, 0xa8, 0xbb, 0xeb, 0xce,
	0x24, 0xb0, 0x88, 0x16, 0xbe, 0x82, 0xd8, 0xfa, 0x4e, 0xe2, 0x2b, 0xf8, 0x20, 0x72, 0xc9, 0x06,
	0x8e, 0x90, 0xc0, 0x59, 0xa5, 0x98, 0xdf, 0x7c, 0xdf, 0x2f, 0xb3, 0xc1, 0xc1, 0x03, 0x38, 0xb2,
	0x42, 0xa1, 0x11, 0x16, 0x54, 0xea, 0x92, 0x32, 0x4e, 0x10, 0x6c, 0x29, 0x53, 0x88, 0x8c, 0xd5,
	0xa4, 0x59, 0xd8, 0x22, 0xa2, 0x32, 0x8e, 0x3c, 0x11, 0x66, 0xb9, 0xa4, 0x59, 0x31, 0x8d, 0x52,
	0xfd, 0xc4, 0x73, 0xad, 0xf3, 0x47, 0xe0, 0x2d, 0x9a, 0xa7, 0xda, 0x02, 0xaf, 0x92, 0x78, 0x47,
	0x15, 0x39, 0x03, 0xd8, 0x3b, 0xa8, 0x0d, 0xc2, 0x3d, 0x1f, 0x2d, 0x8c, 0xe4, 0x42, 0x29, 0x4d,
	0x82, 0xa4, 0x56, 0x7e, 0x3a, 0xf9, 0x5e, 0x0d, 0x76, 0xaf, 0xc0, 0xdd, 0x2d, 0x04, 0xdc, 0xd6,
	0x7a, 0xec, 0x2d, 0x58, 0xbf, 0x00, 0x3a, 0x57, 0x64, 0x1d, 0x1b, 0x47, 0x1d, 0xff, 0x51, 0xb7,
	0x34, 0xcc, 0x0d, 0x3c, 0x17, 0x80, 0x14, 0x1e, 0x2d, 0x83, 0xa2, 0xd1, 0x0a, 0xe1, 0x70, 0xf8,
	0xfe, 0xf3, 0xfb, 0xb1, 0xb2, 0xc3, 0xb6, 0x79, 0x19, 0xf3, 0x02, 0xc1, 0x22, 0x7f, 0x99, 0x7f,
	0x12, 0x99, 0xbd, 0xb2, 0xaf, 0x41, 0xb0, 0x75, 0x2d, 0xb1, 0x5e, 0xb9, 0x94, 0x48, 0xda, 0x3a,
	0x16, 0xf7, 0xa7, 0xb7, 0xd9, 0x46, 0x68, 0xf2, 0x9f, 0x15, 0x2f, 0x36, 0xaa, 0xc4, 0xf6, 0xd9,
	0xb0, 0x43, 0x8c, 0xcf, 0xbc, 0xcb, 0xe7, 0x20, 0xd8, 0xb8, 0x37, 0x99, 0x20, 0xa8, 0x8f, 0x74,
	0xdc, 0x5f, 0xb4, 0x80, 0x35, 0x5a, 0x27, 0x4b, 0xd2, 0xde, 0x68, 0x5c, 0x19, 0x8d, 0xc2, 0xae,
	0x53, 0x9d, 0x6d, 0xc2, 0x9c, 0x4d, 0x8a, 0x6a, 0x6f, 0xba, 0x56, 0x3d, 0xed, 0xe9, 0x5f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x47, 0xf4, 0xa4, 0xe3, 0x9e, 0x02, 0x00, 0x00,
}
