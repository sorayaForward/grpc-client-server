// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: services.proto

package services

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

const (
	Calcule_Send_FullMethodName = "/services.calcule/send"
)

// CalculeClient is the client API for Calcule service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculeClient interface {
	Send(ctx context.Context, in *Number, opts ...grpc.CallOption) (*Result, error)
}

type calculeClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculeClient(cc grpc.ClientConnInterface) CalculeClient {
	return &calculeClient{cc}
}

func (c *calculeClient) Send(ctx context.Context, in *Number, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, Calcule_Send_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculeServer is the server API for Calcule service.
// All implementations must embed UnimplementedCalculeServer
// for forward compatibility
type CalculeServer interface {
	Send(context.Context, *Number) (*Result, error)
	mustEmbedUnimplementedCalculeServer()
}

// UnimplementedCalculeServer must be embedded to have forward compatible implementations.
type UnimplementedCalculeServer struct {
}

func (UnimplementedCalculeServer) Send(context.Context, *Number) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedCalculeServer) mustEmbedUnimplementedCalculeServer() {}

// UnsafeCalculeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculeServer will
// result in compilation errors.
type UnsafeCalculeServer interface {
	mustEmbedUnimplementedCalculeServer()
}

func RegisterCalculeServer(s grpc.ServiceRegistrar, srv CalculeServer) {
	s.RegisterService(&Calcule_ServiceDesc, srv)
}

func _Calcule_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Number)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculeServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Calcule_Send_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculeServer).Send(ctx, req.(*Number))
	}
	return interceptor(ctx, in, info, handler)
}

// Calcule_ServiceDesc is the grpc.ServiceDesc for Calcule service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Calcule_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.calcule",
	HandlerType: (*CalculeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "send",
			Handler:    _Calcule_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}
