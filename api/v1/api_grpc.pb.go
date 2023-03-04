// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api/v1/api.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GatewayDAPIServiceClient is the client API for GatewayDAPIService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayDAPIServiceClient interface {
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
}

type gatewayDAPIServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayDAPIServiceClient(cc grpc.ClientConnInterface) GatewayDAPIServiceClient {
	return &gatewayDAPIServiceClient{cc}
}

func (c *gatewayDAPIServiceClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/api.v1.GatewayDAPIService/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayDAPIServiceServer is the server API for GatewayDAPIService service.
// All implementations must embed UnimplementedGatewayDAPIServiceServer
// for forward compatibility
type GatewayDAPIServiceServer interface {
	Version(context.Context, *emptypb.Empty) (*VersionResponse, error)
	mustEmbedUnimplementedGatewayDAPIServiceServer()
}

// UnimplementedGatewayDAPIServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayDAPIServiceServer struct {
}

func (UnimplementedGatewayDAPIServiceServer) Version(context.Context, *emptypb.Empty) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedGatewayDAPIServiceServer) mustEmbedUnimplementedGatewayDAPIServiceServer() {}

// UnsafeGatewayDAPIServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayDAPIServiceServer will
// result in compilation errors.
type UnsafeGatewayDAPIServiceServer interface {
	mustEmbedUnimplementedGatewayDAPIServiceServer()
}

func RegisterGatewayDAPIServiceServer(s grpc.ServiceRegistrar, srv GatewayDAPIServiceServer) {
	s.RegisterService(&GatewayDAPIService_ServiceDesc, srv)
}

func _GatewayDAPIService_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayDAPIServiceServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.GatewayDAPIService/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayDAPIServiceServer).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// GatewayDAPIService_ServiceDesc is the grpc.ServiceDesc for GatewayDAPIService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GatewayDAPIService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.GatewayDAPIService",
	HandlerType: (*GatewayDAPIServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _GatewayDAPIService_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/api.proto",
}
