// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: api/proto/market/v1/market.proto

package marketv1

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
	MarketService_GetLastTrade_FullMethodName  = "/jax.v1.MarketService/GetLastTrade"
	MarketService_GetAggregates_FullMethodName = "/jax.v1.MarketService/GetAggregates"
)

// MarketServiceClient is the client API for MarketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MarketServiceClient interface {
	GetLastTrade(ctx context.Context, in *GetLastTradeRequest, opts ...grpc.CallOption) (*GetLastTradeResponse, error)
	GetAggregates(ctx context.Context, in *GetAggregatesRequest, opts ...grpc.CallOption) (*GetAggregatesResponse, error)
}

type marketServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMarketServiceClient(cc grpc.ClientConnInterface) MarketServiceClient {
	return &marketServiceClient{cc}
}

func (c *marketServiceClient) GetLastTrade(ctx context.Context, in *GetLastTradeRequest, opts ...grpc.CallOption) (*GetLastTradeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetLastTradeResponse)
	err := c.cc.Invoke(ctx, MarketService_GetLastTrade_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketServiceClient) GetAggregates(ctx context.Context, in *GetAggregatesRequest, opts ...grpc.CallOption) (*GetAggregatesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAggregatesResponse)
	err := c.cc.Invoke(ctx, MarketService_GetAggregates_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MarketServiceServer is the server API for MarketService service.
// All implementations must embed UnimplementedMarketServiceServer
// for forward compatibility.
type MarketServiceServer interface {
	GetLastTrade(context.Context, *GetLastTradeRequest) (*GetLastTradeResponse, error)
	GetAggregates(context.Context, *GetAggregatesRequest) (*GetAggregatesResponse, error)
	mustEmbedUnimplementedMarketServiceServer()
}

// UnimplementedMarketServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMarketServiceServer struct{}

func (UnimplementedMarketServiceServer) GetLastTrade(context.Context, *GetLastTradeRequest) (*GetLastTradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastTrade not implemented")
}
func (UnimplementedMarketServiceServer) GetAggregates(context.Context, *GetAggregatesRequest) (*GetAggregatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAggregates not implemented")
}
func (UnimplementedMarketServiceServer) mustEmbedUnimplementedMarketServiceServer() {}
func (UnimplementedMarketServiceServer) testEmbeddedByValue()                       {}

// UnsafeMarketServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MarketServiceServer will
// result in compilation errors.
type UnsafeMarketServiceServer interface {
	mustEmbedUnimplementedMarketServiceServer()
}

func RegisterMarketServiceServer(s grpc.ServiceRegistrar, srv MarketServiceServer) {
	// If the following call pancis, it indicates UnimplementedMarketServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MarketService_ServiceDesc, srv)
}

func _MarketService_GetLastTrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLastTradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServiceServer).GetLastTrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MarketService_GetLastTrade_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServiceServer).GetLastTrade(ctx, req.(*GetLastTradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketService_GetAggregates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAggregatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServiceServer).GetAggregates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MarketService_GetAggregates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServiceServer).GetAggregates(ctx, req.(*GetAggregatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MarketService_ServiceDesc is the grpc.ServiceDesc for MarketService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MarketService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jax.v1.MarketService",
	HandlerType: (*MarketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLastTrade",
			Handler:    _MarketService_GetLastTrade_Handler,
		},
		{
			MethodName: "GetAggregates",
			Handler:    _MarketService_GetAggregates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/market/v1/market.proto",
}
