// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: api/proto/option/v1/option.proto

package optionv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	OptionService_GetDex_FullMethodName          = "/jax.v1.OptionService/GetDex"
	OptionService_GetDexByStrikes_FullMethodName = "/jax.v1.OptionService/GetDexByStrikes"
	OptionService_GetGex_FullMethodName          = "/jax.v1.OptionService/GetGex"
	OptionService_GetGexByStrikes_FullMethodName = "/jax.v1.OptionService/GetGexByStrikes"
)

// OptionServiceClient is the client API for OptionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// OptionService provides endpoints for delta and gamma exposure calculations
type OptionServiceClient interface {
	// GetDex returns the delta exposure calculations for given parameters
	GetDex(ctx context.Context, in *GetDexRequest, opts ...grpc.CallOption) (*GetDexResponse, error)
	// GetDexByStrikes returns the delta exposure calculations for a specified number of strikes around the spot price
	GetDexByStrikes(ctx context.Context, in *GetDexByStrikesRequest, opts ...grpc.CallOption) (*GetDexResponse, error)
	// GetGex returns the gamma exposure calculations for given parameters
	GetGex(ctx context.Context, in *GetDexRequest, opts ...grpc.CallOption) (*GetDexResponse, error)
	// GetGexByStrikes returns the gamma exposure calculations for a specified number of strikes around the spot price
	GetGexByStrikes(ctx context.Context, in *GetDexByStrikesRequest, opts ...grpc.CallOption) (*GetDexResponse, error)
}

type optionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOptionServiceClient(cc grpc.ClientConnInterface) OptionServiceClient {
	return &optionServiceClient{cc}
}

func (c *optionServiceClient) GetDex(ctx context.Context, in *GetDexRequest, opts ...grpc.CallOption) (*GetDexResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDexResponse)
	err := c.cc.Invoke(ctx, OptionService_GetDex_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *optionServiceClient) GetDexByStrikes(ctx context.Context, in *GetDexByStrikesRequest, opts ...grpc.CallOption) (*GetDexResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDexResponse)
	err := c.cc.Invoke(ctx, OptionService_GetDexByStrikes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *optionServiceClient) GetGex(ctx context.Context, in *GetDexRequest, opts ...grpc.CallOption) (*GetDexResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDexResponse)
	err := c.cc.Invoke(ctx, OptionService_GetGex_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *optionServiceClient) GetGexByStrikes(ctx context.Context, in *GetDexByStrikesRequest, opts ...grpc.CallOption) (*GetDexResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDexResponse)
	err := c.cc.Invoke(ctx, OptionService_GetGexByStrikes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OptionServiceServer is the server API for OptionService service.
// All implementations must embed UnimplementedOptionServiceServer
// for forward compatibility.
//
// OptionService provides endpoints for delta and gamma exposure calculations
type OptionServiceServer interface {
	// GetDex returns the delta exposure calculations for given parameters
	GetDex(context.Context, *GetDexRequest) (*GetDexResponse, error)
	// GetDexByStrikes returns the delta exposure calculations for a specified number of strikes around the spot price
	GetDexByStrikes(context.Context, *GetDexByStrikesRequest) (*GetDexResponse, error)
	// GetGex returns the gamma exposure calculations for given parameters
	GetGex(context.Context, *GetDexRequest) (*GetDexResponse, error)
	// GetGexByStrikes returns the gamma exposure calculations for a specified number of strikes around the spot price
	GetGexByStrikes(context.Context, *GetDexByStrikesRequest) (*GetDexResponse, error)
	mustEmbedUnimplementedOptionServiceServer()
}

// UnimplementedOptionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOptionServiceServer struct{}

func (UnimplementedOptionServiceServer) GetDex(context.Context, *GetDexRequest) (*GetDexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDex not implemented")
}
func (UnimplementedOptionServiceServer) GetDexByStrikes(context.Context, *GetDexByStrikesRequest) (*GetDexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDexByStrikes not implemented")
}
func (UnimplementedOptionServiceServer) GetGex(context.Context, *GetDexRequest) (*GetDexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGex not implemented")
}
func (UnimplementedOptionServiceServer) GetGexByStrikes(context.Context, *GetDexByStrikesRequest) (*GetDexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGexByStrikes not implemented")
}
func (UnimplementedOptionServiceServer) mustEmbedUnimplementedOptionServiceServer() {}
func (UnimplementedOptionServiceServer) testEmbeddedByValue()                       {}

// UnsafeOptionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OptionServiceServer will
// result in compilation errors.
type UnsafeOptionServiceServer interface {
	mustEmbedUnimplementedOptionServiceServer()
}

func RegisterOptionServiceServer(s grpc.ServiceRegistrar, srv OptionServiceServer) {
	// If the following call pancis, it indicates UnimplementedOptionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OptionService_ServiceDesc, srv)
}

func _OptionService_GetDex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OptionServiceServer).GetDex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OptionService_GetDex_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OptionServiceServer).GetDex(ctx, req.(*GetDexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OptionService_GetDexByStrikes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDexByStrikesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OptionServiceServer).GetDexByStrikes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OptionService_GetDexByStrikes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OptionServiceServer).GetDexByStrikes(ctx, req.(*GetDexByStrikesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OptionService_GetGex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OptionServiceServer).GetGex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OptionService_GetGex_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OptionServiceServer).GetGex(ctx, req.(*GetDexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OptionService_GetGexByStrikes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDexByStrikesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OptionServiceServer).GetGexByStrikes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OptionService_GetGexByStrikes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OptionServiceServer).GetGexByStrikes(ctx, req.(*GetDexByStrikesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OptionService_ServiceDesc is the grpc.ServiceDesc for OptionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OptionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jax.v1.OptionService",
	HandlerType: (*OptionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDex",
			Handler:    _OptionService_GetDex_Handler,
		},
		{
			MethodName: "GetDexByStrikes",
			Handler:    _OptionService_GetDexByStrikes_Handler,
		},
		{
			MethodName: "GetGex",
			Handler:    _OptionService_GetGex_Handler,
		},
		{
			MethodName: "GetGexByStrikes",
			Handler:    _OptionService_GetGexByStrikes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/option/v1/option.proto",
}