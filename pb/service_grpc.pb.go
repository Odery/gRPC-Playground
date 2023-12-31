// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: service.proto

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

const (
	CustomerCreator_CreateCustomer_FullMethodName = "/CustomerCreator/CreateCustomer"
)

// CustomerCreatorClient is the client API for CustomerCreator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerCreatorClient interface {
	CreateCustomer(ctx context.Context, in *NewCustomer, opts ...grpc.CallOption) (*Customer, error)
}

type customerCreatorClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerCreatorClient(cc grpc.ClientConnInterface) CustomerCreatorClient {
	return &customerCreatorClient{cc}
}

func (c *customerCreatorClient) CreateCustomer(ctx context.Context, in *NewCustomer, opts ...grpc.CallOption) (*Customer, error) {
	out := new(Customer)
	err := c.cc.Invoke(ctx, CustomerCreator_CreateCustomer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerCreatorServer is the server API for CustomerCreator service.
// All implementations must embed UnimplementedCustomerCreatorServer
// for forward compatibility
type CustomerCreatorServer interface {
	CreateCustomer(context.Context, *NewCustomer) (*Customer, error)
	mustEmbedUnimplementedCustomerCreatorServer()
}

// UnimplementedCustomerCreatorServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerCreatorServer struct {
}

func (UnimplementedCustomerCreatorServer) CreateCustomer(context.Context, *NewCustomer) (*Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCustomer not implemented")
}
func (UnimplementedCustomerCreatorServer) mustEmbedUnimplementedCustomerCreatorServer() {}

// UnsafeCustomerCreatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerCreatorServer will
// result in compilation errors.
type UnsafeCustomerCreatorServer interface {
	mustEmbedUnimplementedCustomerCreatorServer()
}

func RegisterCustomerCreatorServer(s grpc.ServiceRegistrar, srv CustomerCreatorServer) {
	s.RegisterService(&CustomerCreator_ServiceDesc, srv)
}

func _CustomerCreator_CreateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewCustomer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerCreatorServer).CreateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomerCreator_CreateCustomer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerCreatorServer).CreateCustomer(ctx, req.(*NewCustomer))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerCreator_ServiceDesc is the grpc.ServiceDesc for CustomerCreator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerCreator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CustomerCreator",
	HandlerType: (*CustomerCreatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCustomer",
			Handler:    _CustomerCreator_CreateCustomer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
