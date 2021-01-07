// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package carshop

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// BackOfficeServiceClient is the client API for BackOfficeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BackOfficeServiceClient interface {
	RegisterOwner(ctx context.Context, in *Owner, opts ...grpc.CallOption) (*empty.Empty, error)
	RegisterCar(ctx context.Context, in *Car, opts ...grpc.CallOption) (*empty.Empty, error)
}

type backOfficeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBackOfficeServiceClient(cc grpc.ClientConnInterface) BackOfficeServiceClient {
	return &backOfficeServiceClient{cc}
}

func (c *backOfficeServiceClient) RegisterOwner(ctx context.Context, in *Owner, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/com.aviebrantz.carshop.BackOfficeService/RegisterOwner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backOfficeServiceClient) RegisterCar(ctx context.Context, in *Car, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/com.aviebrantz.carshop.BackOfficeService/RegisterCar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackOfficeServiceServer is the server API for BackOfficeService service.
// All implementations should embed UnimplementedBackOfficeServiceServer
// for forward compatibility
type BackOfficeServiceServer interface {
	RegisterOwner(context.Context, *Owner) (*empty.Empty, error)
	RegisterCar(context.Context, *Car) (*empty.Empty, error)
}

// UnimplementedBackOfficeServiceServer should be embedded to have forward compatible implementations.
type UnimplementedBackOfficeServiceServer struct {
}

func (UnimplementedBackOfficeServiceServer) RegisterOwner(context.Context, *Owner) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterOwner not implemented")
}
func (UnimplementedBackOfficeServiceServer) RegisterCar(context.Context, *Car) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterCar not implemented")
}

// UnsafeBackOfficeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BackOfficeServiceServer will
// result in compilation errors.
type UnsafeBackOfficeServiceServer interface {
	mustEmbedUnimplementedBackOfficeServiceServer()
}

func RegisterBackOfficeServiceServer(s *grpc.Server, srv BackOfficeServiceServer) {
	s.RegisterService(&_BackOfficeService_serviceDesc, srv)
}

func _BackOfficeService_RegisterOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Owner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackOfficeServiceServer).RegisterOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.aviebrantz.carshop.BackOfficeService/RegisterOwner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackOfficeServiceServer).RegisterOwner(ctx, req.(*Owner))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackOfficeService_RegisterCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Car)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackOfficeServiceServer).RegisterCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.aviebrantz.carshop.BackOfficeService/RegisterCar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackOfficeServiceServer).RegisterCar(ctx, req.(*Car))
	}
	return interceptor(ctx, in, info, handler)
}

var _BackOfficeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.aviebrantz.carshop.BackOfficeService",
	HandlerType: (*BackOfficeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterOwner",
			Handler:    _BackOfficeService_RegisterOwner_Handler,
		},
		{
			MethodName: "RegisterCar",
			Handler:    _BackOfficeService_RegisterCar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backoffice.proto",
}
