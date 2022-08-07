// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: proto/calculator.proto

package pb

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

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	Add(ctx context.Context, in *CalculatorPayload, opts ...grpc.CallOption) (*CalculatorResult, error)
	Substract(ctx context.Context, in *CalculatorPayload, opts ...grpc.CallOption) (*CalculatorResult, error)
	Multiply(ctx context.Context, in *CalculatorPayload, opts ...grpc.CallOption) (*CalculatorResult, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Add(ctx context.Context, in *CalculatorPayload, opts ...grpc.CallOption) (*CalculatorResult, error) {
	out := new(CalculatorResult)
	err := c.cc.Invoke(ctx, "/pb.CalculatorService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Substract(ctx context.Context, in *CalculatorPayload, opts ...grpc.CallOption) (*CalculatorResult, error) {
	out := new(CalculatorResult)
	err := c.cc.Invoke(ctx, "/pb.CalculatorService/Substract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Multiply(ctx context.Context, in *CalculatorPayload, opts ...grpc.CallOption) (*CalculatorResult, error) {
	out := new(CalculatorResult)
	err := c.cc.Invoke(ctx, "/pb.CalculatorService/Multiply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
// All implementations must embed UnimplementedCalculatorServiceServer
// for forward compatibility
type CalculatorServiceServer interface {
	Add(context.Context, *CalculatorPayload) (*CalculatorResult, error)
	Substract(context.Context, *CalculatorPayload) (*CalculatorResult, error)
	Multiply(context.Context, *CalculatorPayload) (*CalculatorResult, error)
	mustEmbedUnimplementedCalculatorServiceServer()
}

// UnimplementedCalculatorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (UnimplementedCalculatorServiceServer) Add(context.Context, *CalculatorPayload) (*CalculatorResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedCalculatorServiceServer) Substract(context.Context, *CalculatorPayload) (*CalculatorResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Substract not implemented")
}
func (UnimplementedCalculatorServiceServer) Multiply(context.Context, *CalculatorPayload) (*CalculatorResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiply not implemented")
}
func (UnimplementedCalculatorServiceServer) mustEmbedUnimplementedCalculatorServiceServer() {}

// UnsafeCalculatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatorServiceServer will
// result in compilation errors.
type UnsafeCalculatorServiceServer interface {
	mustEmbedUnimplementedCalculatorServiceServer()
}

func RegisterCalculatorServiceServer(s grpc.ServiceRegistrar, srv CalculatorServiceServer) {
	s.RegisterService(&CalculatorService_ServiceDesc, srv)
}

func _CalculatorService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculatorPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CalculatorService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Add(ctx, req.(*CalculatorPayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Substract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculatorPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Substract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CalculatorService/Substract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Substract(ctx, req.(*CalculatorPayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculatorPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CalculatorService/Multiply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Multiply(ctx, req.(*CalculatorPayload))
	}
	return interceptor(ctx, in, info, handler)
}

// CalculatorService_ServiceDesc is the grpc.ServiceDesc for CalculatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalculatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _CalculatorService_Add_Handler,
		},
		{
			MethodName: "Substract",
			Handler:    _CalculatorService_Substract_Handler,
		},
		{
			MethodName: "Multiply",
			Handler:    _CalculatorService_Multiply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/calculator.proto",
}