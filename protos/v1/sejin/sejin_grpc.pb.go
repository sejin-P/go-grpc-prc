// Code generated by protoc-gen-go-grpc-prc. DO NOT EDIT.

package sejin

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SejinClient is the client API for Sejin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SejinClient interface {
	GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoResponse, error)
}

type sejinClient struct {
	cc grpc.ClientConnInterface
}

func NewSejinClient(cc grpc.ClientConnInterface) SejinClient {
	return &sejinClient{cc}
}

func (c *sejinClient) GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoResponse, error) {
	out := new(GetInfoResponse)
	err := c.cc.Invoke(ctx, "/v1.sejin.Sejin/GetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SejinServer is the server API for Sejin service.
// All implementations must embed UnimplementedSejinServer
// for forward compatibility
type SejinServer interface {
	GetInfo(context.Context, *GetInfoRequest) (*GetInfoResponse, error)
	mustEmbedUnimplementedSejinServer()
}

// UnimplementedSejinServer must be embedded to have forward compatible implementations.
type UnimplementedSejinServer struct {
}

func (UnimplementedSejinServer) GetInfo(context.Context, *GetInfoRequest) (*GetInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}
func (UnimplementedSejinServer) mustEmbedUnimplementedSejinServer() {}

// UnsafeSejinServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SejinServer will
// result in compilation errors.
type UnsafeSejinServer interface {
	mustEmbedUnimplementedSejinServer()
}

func RegisterSejinServer(s grpc.ServiceRegistrar, srv SejinServer) {
	s.RegisterService(&Sejin_ServiceDesc, srv)
}

func _Sejin_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SejinServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.sejin.Sejin/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SejinServer).GetInfo(ctx, req.(*GetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Sejin_ServiceDesc is the grpc.ServiceDesc for Sejin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sejin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.sejin.Sejin",
	HandlerType: (*SejinServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInfo",
			Handler:    _Sejin_GetInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/v1/sejin/sejin.proto",
}