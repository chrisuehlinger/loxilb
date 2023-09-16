// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: loxinet/xsync.proto

package loxinet

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

// XSyncClient is the client API for XSync service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type XSyncClient interface {
	DpWorkOnCtGetGRPC(ctx context.Context, in *ConnGet, opts ...grpc.CallOption) (*XSyncReply, error)
	DpWorkOnCtModGRPC(ctx context.Context, in *CtInfoMod, opts ...grpc.CallOption) (*XSyncReply, error)
	DpWorkOnBlockCtModGRPC(ctx context.Context, in *BlockCtInfoMod, opts ...grpc.CallOption) (*XSyncReply, error)
}

type xSyncClient struct {
	cc grpc.ClientConnInterface
}

func NewXSyncClient(cc grpc.ClientConnInterface) XSyncClient {
	return &xSyncClient{cc}
}

func (c *xSyncClient) DpWorkOnCtGetGRPC(ctx context.Context, in *ConnGet, opts ...grpc.CallOption) (*XSyncReply, error) {
	out := new(XSyncReply)
	err := c.cc.Invoke(ctx, "/XSync/DpWorkOnCtGetGRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xSyncClient) DpWorkOnCtModGRPC(ctx context.Context, in *CtInfoMod, opts ...grpc.CallOption) (*XSyncReply, error) {
	out := new(XSyncReply)
	err := c.cc.Invoke(ctx, "/XSync/DpWorkOnCtModGRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xSyncClient) DpWorkOnBlockCtModGRPC(ctx context.Context, in *BlockCtInfoMod, opts ...grpc.CallOption) (*XSyncReply, error) {
	out := new(XSyncReply)
	err := c.cc.Invoke(ctx, "/XSync/DpWorkOnBlockCtModGRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// XSyncServer is the server API for XSync service.
// All implementations must embed UnimplementedXSyncServer
// for forward compatibility
type XSyncServer interface {
	DpWorkOnCtGetGRPC(context.Context, *ConnGet) (*XSyncReply, error)
	DpWorkOnCtModGRPC(context.Context, *CtInfoMod) (*XSyncReply, error)
	DpWorkOnBlockCtModGRPC(context.Context, *BlockCtInfoMod) (*XSyncReply, error)
	mustEmbedUnimplementedXSyncServer()
}

// UnimplementedXSyncServer must be embedded to have forward compatible implementations.
type UnimplementedXSyncServer struct {
}

func (UnimplementedXSyncServer) DpWorkOnCtGetGRPC(context.Context, *ConnGet) (*XSyncReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DpWorkOnCtGetGRPC not implemented")
}
func (UnimplementedXSyncServer) DpWorkOnCtModGRPC(context.Context, *CtInfoMod) (*XSyncReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DpWorkOnCtModGRPC not implemented")
}
func (UnimplementedXSyncServer) DpWorkOnBlockCtModGRPC(context.Context, *BlockCtInfoMod) (*XSyncReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DpWorkOnBlockCtModGRPC not implemented")
}
func (UnimplementedXSyncServer) mustEmbedUnimplementedXSyncServer() {}

// UnsafeXSyncServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to XSyncServer will
// result in compilation errors.
type UnsafeXSyncServer interface {
	mustEmbedUnimplementedXSyncServer()
}

func RegisterXSyncServer(s grpc.ServiceRegistrar, srv XSyncServer) {
	s.RegisterService(&XSync_ServiceDesc, srv)
}

func _XSync_DpWorkOnCtGetGRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnGet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XSyncServer).DpWorkOnCtGetGRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/XSync/DpWorkOnCtGetGRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XSyncServer).DpWorkOnCtGetGRPC(ctx, req.(*ConnGet))
	}
	return interceptor(ctx, in, info, handler)
}

func _XSync_DpWorkOnCtModGRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CtInfoMod)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XSyncServer).DpWorkOnCtModGRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/XSync/DpWorkOnCtModGRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XSyncServer).DpWorkOnCtModGRPC(ctx, req.(*CtInfoMod))
	}
	return interceptor(ctx, in, info, handler)
}

func _XSync_DpWorkOnBlockCtModGRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockCtInfoMod)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XSyncServer).DpWorkOnBlockCtModGRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/XSync/DpWorkOnBlockCtModGRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XSyncServer).DpWorkOnBlockCtModGRPC(ctx, req.(*BlockCtInfoMod))
	}
	return interceptor(ctx, in, info, handler)
}

// XSync_ServiceDesc is the grpc.ServiceDesc for XSync service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var XSync_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "XSync",
	HandlerType: (*XSyncServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DpWorkOnCtGetGRPC",
			Handler:    _XSync_DpWorkOnCtGetGRPC_Handler,
		},
		{
			MethodName: "DpWorkOnCtModGRPC",
			Handler:    _XSync_DpWorkOnCtModGRPC_Handler,
		},
		{
			MethodName: "DpWorkOnBlockCtModGRPC",
			Handler:    _XSync_DpWorkOnBlockCtModGRPC_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "loxinet/xsync.proto",
}