// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package service

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

// PortDomainServiceClient is the client API for PortDomainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PortDomainServiceClient interface {
	GetPort(ctx context.Context, in *PortRequest, opts ...grpc.CallOption) (*PortResponse, error)
	CreateOrUpdatePort(ctx context.Context, in *PortRequest, opts ...grpc.CallOption) (*PortResponse, error)
	DeletePort(ctx context.Context, in *PortRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type portDomainServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPortDomainServiceClient(cc grpc.ClientConnInterface) PortDomainServiceClient {
	return &portDomainServiceClient{cc}
}

func (c *portDomainServiceClient) GetPort(ctx context.Context, in *PortRequest, opts ...grpc.CallOption) (*PortResponse, error) {
	out := new(PortResponse)
	err := c.cc.Invoke(ctx, "/service.PortDomainService/GetPort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portDomainServiceClient) CreateOrUpdatePort(ctx context.Context, in *PortRequest, opts ...grpc.CallOption) (*PortResponse, error) {
	out := new(PortResponse)
	err := c.cc.Invoke(ctx, "/service.PortDomainService/CreateOrUpdatePort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portDomainServiceClient) DeletePort(ctx context.Context, in *PortRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/service.PortDomainService/DeletePort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortDomainServiceServer is the server API for PortDomainService service.
// All implementations must embed UnimplementedPortDomainServiceServer
// for forward compatibility
type PortDomainServiceServer interface {
	GetPort(context.Context, *PortRequest) (*PortResponse, error)
	CreateOrUpdatePort(context.Context, *PortRequest) (*PortResponse, error)
	DeletePort(context.Context, *PortRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedPortDomainServiceServer()
}

// UnimplementedPortDomainServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPortDomainServiceServer struct {
}

func (UnimplementedPortDomainServiceServer) GetPort(context.Context, *PortRequest) (*PortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPort not implemented")
}
func (UnimplementedPortDomainServiceServer) CreateOrUpdatePort(context.Context, *PortRequest) (*PortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdatePort not implemented")
}
func (UnimplementedPortDomainServiceServer) DeletePort(context.Context, *PortRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePort not implemented")
}
func (UnimplementedPortDomainServiceServer) mustEmbedUnimplementedPortDomainServiceServer() {}

// UnsafePortDomainServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PortDomainServiceServer will
// result in compilation errors.
type UnsafePortDomainServiceServer interface {
	mustEmbedUnimplementedPortDomainServiceServer()
}

func RegisterPortDomainServiceServer(s grpc.ServiceRegistrar, srv PortDomainServiceServer) {
	s.RegisterService(&PortDomainService_ServiceDesc, srv)
}

func _PortDomainService_GetPort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortDomainServiceServer).GetPort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.PortDomainService/GetPort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortDomainServiceServer).GetPort(ctx, req.(*PortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortDomainService_CreateOrUpdatePort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortDomainServiceServer).CreateOrUpdatePort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.PortDomainService/CreateOrUpdatePort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortDomainServiceServer).CreateOrUpdatePort(ctx, req.(*PortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortDomainService_DeletePort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortDomainServiceServer).DeletePort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.PortDomainService/DeletePort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortDomainServiceServer).DeletePort(ctx, req.(*PortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PortDomainService_ServiceDesc is the grpc.ServiceDesc for PortDomainService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PortDomainService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.PortDomainService",
	HandlerType: (*PortDomainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPort",
			Handler:    _PortDomainService_GetPort_Handler,
		},
		{
			MethodName: "CreateOrUpdatePort",
			Handler:    _PortDomainService_CreateOrUpdatePort_Handler,
		},
		{
			MethodName: "DeletePort",
			Handler:    _PortDomainService_DeletePort_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "portdomainservice.proto",
}
