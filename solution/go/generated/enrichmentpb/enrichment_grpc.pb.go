// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.2
// source: enrichment.proto

package enrichmentpb

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

// EnrichmentClient is the client API for Enrichment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EnrichmentClient interface {
	IPToHost(ctx context.Context, in *IP, opts ...grpc.CallOption) (*Host, error)
	HostToIP(ctx context.Context, in *Host, opts ...grpc.CallOption) (*IP, error)
}

type enrichmentClient struct {
	cc grpc.ClientConnInterface
}

func NewEnrichmentClient(cc grpc.ClientConnInterface) EnrichmentClient {
	return &enrichmentClient{cc}
}

func (c *enrichmentClient) IPToHost(ctx context.Context, in *IP, opts ...grpc.CallOption) (*Host, error) {
	out := new(Host)
	err := c.cc.Invoke(ctx, "/enrichment.Enrichment/IPToHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enrichmentClient) HostToIP(ctx context.Context, in *Host, opts ...grpc.CallOption) (*IP, error) {
	out := new(IP)
	err := c.cc.Invoke(ctx, "/enrichment.Enrichment/HostToIP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnrichmentServer is the server API for Enrichment service.
// All implementations must embed UnimplementedEnrichmentServer
// for forward compatibility
type EnrichmentServer interface {
	IPToHost(context.Context, *IP) (*Host, error)
	HostToIP(context.Context, *Host) (*IP, error)
	mustEmbedUnimplementedEnrichmentServer()
}

// UnimplementedEnrichmentServer must be embedded to have forward compatible implementations.
type UnimplementedEnrichmentServer struct {
}

func (UnimplementedEnrichmentServer) IPToHost(context.Context, *IP) (*Host, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IPToHost not implemented")
}
func (UnimplementedEnrichmentServer) HostToIP(context.Context, *Host) (*IP, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HostToIP not implemented")
}
func (UnimplementedEnrichmentServer) mustEmbedUnimplementedEnrichmentServer() {}

// UnsafeEnrichmentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EnrichmentServer will
// result in compilation errors.
type UnsafeEnrichmentServer interface {
	mustEmbedUnimplementedEnrichmentServer()
}

func RegisterEnrichmentServer(s grpc.ServiceRegistrar, srv EnrichmentServer) {
	s.RegisterService(&Enrichment_ServiceDesc, srv)
}

func _Enrichment_IPToHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IP)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnrichmentServer).IPToHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/enrichment.Enrichment/IPToHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnrichmentServer).IPToHost(ctx, req.(*IP))
	}
	return interceptor(ctx, in, info, handler)
}

func _Enrichment_HostToIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Host)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnrichmentServer).HostToIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/enrichment.Enrichment/HostToIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnrichmentServer).HostToIP(ctx, req.(*Host))
	}
	return interceptor(ctx, in, info, handler)
}

// Enrichment_ServiceDesc is the grpc.ServiceDesc for Enrichment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Enrichment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "enrichment.Enrichment",
	HandlerType: (*EnrichmentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IPToHost",
			Handler:    _Enrichment_IPToHost_Handler,
		},
		{
			MethodName: "HostToIP",
			Handler:    _Enrichment_HostToIP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "enrichment.proto",
}
