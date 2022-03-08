// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: wireFormat.proto

package dasrpc

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

// DASServiceImplClient is the client API for DASServiceImpl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DASServiceImplClient interface {
	Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreResponse, error)
	Retrieve(ctx context.Context, in *RetrieveRequest, opts ...grpc.CallOption) (*RetrieveResponse, error)
}

type dASServiceImplClient struct {
	cc grpc.ClientConnInterface
}

func NewDASServiceImplClient(cc grpc.ClientConnInterface) DASServiceImplClient {
	return &dASServiceImplClient{cc}
}

func (c *dASServiceImplClient) Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreResponse, error) {
	out := new(StoreResponse)
	err := c.cc.Invoke(ctx, "/das.DASServiceImpl/Store", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dASServiceImplClient) Retrieve(ctx context.Context, in *RetrieveRequest, opts ...grpc.CallOption) (*RetrieveResponse, error) {
	out := new(RetrieveResponse)
	err := c.cc.Invoke(ctx, "/das.DASServiceImpl/Retrieve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DASServiceImplServer is the server API for DASServiceImpl service.
// All implementations must embed UnimplementedDASServiceImplServer
// for forward compatibility
type DASServiceImplServer interface {
	Store(context.Context, *StoreRequest) (*StoreResponse, error)
	Retrieve(context.Context, *RetrieveRequest) (*RetrieveResponse, error)
	mustEmbedUnimplementedDASServiceImplServer()
}

// UnimplementedDASServiceImplServer must be embedded to have forward compatible implementations.
type UnimplementedDASServiceImplServer struct {
}

func (UnimplementedDASServiceImplServer) Store(context.Context, *StoreRequest) (*StoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Store not implemented")
}
func (UnimplementedDASServiceImplServer) Retrieve(context.Context, *RetrieveRequest) (*RetrieveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Retrieve not implemented")
}
func (UnimplementedDASServiceImplServer) mustEmbedUnimplementedDASServiceImplServer() {}

// UnsafeDASServiceImplServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DASServiceImplServer will
// result in compilation errors.
type UnsafeDASServiceImplServer interface {
	mustEmbedUnimplementedDASServiceImplServer()
}

func RegisterDASServiceImplServer(s grpc.ServiceRegistrar, srv DASServiceImplServer) {
	s.RegisterService(&DASServiceImpl_ServiceDesc, srv)
}

func _DASServiceImpl_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DASServiceImplServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/das.DASServiceImpl/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DASServiceImplServer).Store(ctx, req.(*StoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DASServiceImpl_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DASServiceImplServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/das.DASServiceImpl/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DASServiceImplServer).Retrieve(ctx, req.(*RetrieveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DASServiceImpl_ServiceDesc is the grpc.ServiceDesc for DASServiceImpl service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DASServiceImpl_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "das.DASServiceImpl",
	HandlerType: (*DASServiceImplServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Store",
			Handler:    _DASServiceImpl_Store_Handler,
		},
		{
			MethodName: "Retrieve",
			Handler:    _DASServiceImpl_Retrieve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wireFormat.proto",
}