// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpcweb

import (
	context "context"
	link "github.com/batazor/shortlink/internal/services/api/domain/link"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LinkClient is the client API for Link service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LinkClient interface {
	GetLinks(ctx context.Context, in *LinkRequest, opts ...grpc.CallOption) (*link.Links, error)
	GetLink(ctx context.Context, in *link.Link, opts ...grpc.CallOption) (*link.Link, error)
	CreateLink(ctx context.Context, in *link.Link, opts ...grpc.CallOption) (*link.Link, error)
	DeleteLink(ctx context.Context, in *link.Link, opts ...grpc.CallOption) (*empty.Empty, error)
}

type linkClient struct {
	cc grpc.ClientConnInterface
}

func NewLinkClient(cc grpc.ClientConnInterface) LinkClient {
	return &linkClient{cc}
}

func (c *linkClient) GetLinks(ctx context.Context, in *LinkRequest, opts ...grpc.CallOption) (*link.Links, error) {
	out := new(link.Links)
	err := c.cc.Invoke(ctx, "/grpcweb.Link/GetLinks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkClient) GetLink(ctx context.Context, in *link.Link, opts ...grpc.CallOption) (*link.Link, error) {
	out := new(link.Link)
	err := c.cc.Invoke(ctx, "/grpcweb.Link/GetLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkClient) CreateLink(ctx context.Context, in *link.Link, opts ...grpc.CallOption) (*link.Link, error) {
	out := new(link.Link)
	err := c.cc.Invoke(ctx, "/grpcweb.Link/CreateLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkClient) DeleteLink(ctx context.Context, in *link.Link, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/grpcweb.Link/DeleteLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LinkServer is the server API for Link service.
// All implementations must embed UnimplementedLinkServer
// for forward compatibility
type LinkServer interface {
	GetLinks(context.Context, *LinkRequest) (*link.Links, error)
	GetLink(context.Context, *link.Link) (*link.Link, error)
	CreateLink(context.Context, *link.Link) (*link.Link, error)
	DeleteLink(context.Context, *link.Link) (*empty.Empty, error)
	mustEmbedUnimplementedLinkServer()
}

// UnimplementedLinkServer must be embedded to have forward compatible implementations.
type UnimplementedLinkServer struct {
}

func (UnimplementedLinkServer) GetLinks(context.Context, *LinkRequest) (*link.Links, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLinks not implemented")
}
func (UnimplementedLinkServer) GetLink(context.Context, *link.Link) (*link.Link, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLink not implemented")
}
func (UnimplementedLinkServer) CreateLink(context.Context, *link.Link) (*link.Link, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLink not implemented")
}
func (UnimplementedLinkServer) DeleteLink(context.Context, *link.Link) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLink not implemented")
}
func (UnimplementedLinkServer) mustEmbedUnimplementedLinkServer() {}

// UnsafeLinkServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LinkServer will
// result in compilation errors.
type UnsafeLinkServer interface {
	mustEmbedUnimplementedLinkServer()
}

func RegisterLinkServer(s grpc.ServiceRegistrar, srv LinkServer) {
	s.RegisterService(&Link_ServiceDesc, srv)
}

func _Link_GetLinks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkServer).GetLinks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcweb.Link/GetLinks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkServer).GetLinks(ctx, req.(*LinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Link_GetLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(link.Link)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkServer).GetLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcweb.Link/GetLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkServer).GetLink(ctx, req.(*link.Link))
	}
	return interceptor(ctx, in, info, handler)
}

func _Link_CreateLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(link.Link)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkServer).CreateLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcweb.Link/CreateLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkServer).CreateLink(ctx, req.(*link.Link))
	}
	return interceptor(ctx, in, info, handler)
}

func _Link_DeleteLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(link.Link)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkServer).DeleteLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcweb.Link/DeleteLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkServer).DeleteLink(ctx, req.(*link.Link))
	}
	return interceptor(ctx, in, info, handler)
}

// Link_ServiceDesc is the grpc.ServiceDesc for Link service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Link_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcweb.Link",
	HandlerType: (*LinkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLinks",
			Handler:    _Link_GetLinks_Handler,
		},
		{
			MethodName: "GetLink",
			Handler:    _Link_GetLink_Handler,
		},
		{
			MethodName: "CreateLink",
			Handler:    _Link_CreateLink_Handler,
		},
		{
			MethodName: "DeleteLink",
			Handler:    _Link_DeleteLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
