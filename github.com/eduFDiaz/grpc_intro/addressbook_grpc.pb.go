// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc_intro

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

// AddressBookServiceClient is the client API for AddressBookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AddressBookServiceClient interface {
	GetAllAddresses(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AddressBook, error)
	GetPersonByName(ctx context.Context, in *ReqGetPerson, opts ...grpc.CallOption) (*Person, error)
}

type addressBookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAddressBookServiceClient(cc grpc.ClientConnInterface) AddressBookServiceClient {
	return &addressBookServiceClient{cc}
}

func (c *addressBookServiceClient) GetAllAddresses(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AddressBook, error) {
	out := new(AddressBook)
	err := c.cc.Invoke(ctx, "/tutorial.AddressBookService/GetAllAddresses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressBookServiceClient) GetPersonByName(ctx context.Context, in *ReqGetPerson, opts ...grpc.CallOption) (*Person, error) {
	out := new(Person)
	err := c.cc.Invoke(ctx, "/tutorial.AddressBookService/GetPersonByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddressBookServiceServer is the server API for AddressBookService service.
// All implementations must embed UnimplementedAddressBookServiceServer
// for forward compatibility
type AddressBookServiceServer interface {
	GetAllAddresses(context.Context, *Empty) (*AddressBook, error)
	GetPersonByName(context.Context, *ReqGetPerson) (*Person, error)
	mustEmbedUnimplementedAddressBookServiceServer()
}

// UnimplementedAddressBookServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAddressBookServiceServer struct {
}

func (UnimplementedAddressBookServiceServer) GetAllAddresses(context.Context, *Empty) (*AddressBook, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllAddresses not implemented")
}
func (UnimplementedAddressBookServiceServer) GetPersonByName(context.Context, *ReqGetPerson) (*Person, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPersonByName not implemented")
}
func (UnimplementedAddressBookServiceServer) mustEmbedUnimplementedAddressBookServiceServer() {}

// UnsafeAddressBookServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AddressBookServiceServer will
// result in compilation errors.
type UnsafeAddressBookServiceServer interface {
	mustEmbedUnimplementedAddressBookServiceServer()
}

func RegisterAddressBookServiceServer(s grpc.ServiceRegistrar, srv AddressBookServiceServer) {
	s.RegisterService(&AddressBookService_ServiceDesc, srv)
}

func _AddressBookService_GetAllAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressBookServiceServer).GetAllAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tutorial.AddressBookService/GetAllAddresses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressBookServiceServer).GetAllAddresses(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressBookService_GetPersonByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetPerson)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressBookServiceServer).GetPersonByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tutorial.AddressBookService/GetPersonByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressBookServiceServer).GetPersonByName(ctx, req.(*ReqGetPerson))
	}
	return interceptor(ctx, in, info, handler)
}

// AddressBookService_ServiceDesc is the grpc.ServiceDesc for AddressBookService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AddressBookService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tutorial.AddressBookService",
	HandlerType: (*AddressBookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllAddresses",
			Handler:    _AddressBookService_GetAllAddresses_Handler,
		},
		{
			MethodName: "GetPersonByName",
			Handler:    _AddressBookService_GetPersonByName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "addressbook.proto",
}
