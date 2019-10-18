// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/api/grpc-web/api.proto

package grpc_web

import (
	context "context"
	fmt "fmt"
	link "github.com/batazor/shortlink/pkg/link"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("pkg/api/grpc-web/api.proto", fileDescriptor_54d0e997a1f3c768) }

var fileDescriptor_54d0e997a1f3c768 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0xc8, 0x4e, 0xd7,
	0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0x2d, 0x4f, 0x4d, 0x02, 0x71, 0xf4, 0x0a,
	0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x38, 0x40, 0x62, 0xf1, 0xe5, 0xa9, 0x49, 0x52, 0x32, 0xe9, 0xf9,
	0xf9, 0xe9, 0x39, 0xa9, 0x60, 0x85, 0x89, 0x79, 0x79, 0xf9, 0x25, 0x89, 0x25, 0x99, 0xf9, 0x79,
	0xc5, 0x10, 0x75, 0x52, 0xd2, 0x50, 0x59, 0x30, 0x2f, 0xa9, 0x34, 0x4d, 0x3f, 0x35, 0xb7, 0xa0,
	0xa4, 0x12, 0x2a, 0x29, 0x0c, 0xb2, 0x20, 0x27, 0x33, 0x2f, 0x1b, 0x4c, 0x40, 0x04, 0x8d, 0x76,
	0x30, 0x72, 0xb1, 0xf8, 0x64, 0xe6, 0x65, 0x0b, 0x99, 0x71, 0xb1, 0xbb, 0xa7, 0x96, 0x80, 0x99,
	0x5c, 0x7a, 0x60, 0x05, 0x20, 0xb6, 0x14, 0x12, 0x5b, 0x49, 0xb8, 0xe9, 0xf2, 0x93, 0xc9, 0x4c,
	0xbc, 0x42, 0xdc, 0x60, 0xdb, 0xab, 0x33, 0x12, 0x8b, 0x33, 0x6a, 0x85, 0x4c, 0xb9, 0xb8, 0x9c,
	0x8b, 0x52, 0x13, 0x4b, 0x52, 0xf1, 0x6a, 0xe5, 0x07, 0x6b, 0xe5, 0x54, 0x62, 0x01, 0x69, 0xb5,
	0x62, 0xd4, 0x12, 0x72, 0xe4, 0xe2, 0x72, 0x49, 0xcd, 0x49, 0xc5, 0xa2, 0x4d, 0x4c, 0x0f, 0xe2,
	0x09, 0x3d, 0x98, 0x27, 0xf4, 0x5c, 0x41, 0x9e, 0x80, 0x19, 0xa1, 0x05, 0x33, 0x22, 0x89, 0x0d,
	0xac, 0xc0, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x92, 0xb5, 0xd4, 0x39, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LinkClient is the client API for Link service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LinkClient interface {
	GetLink(ctx context.Context, in *link.Link, opts ...grpc.CallOption) (*link.Link, error)
	CreateLink(ctx context.Context, in *link.Link, opts ...grpc.CallOption) (*link.Link, error)
	DeleteLink(ctx context.Context, in *link.Link, opts ...grpc.CallOption) (*empty.Empty, error)
}

type linkClient struct {
	cc *grpc.ClientConn
}

func NewLinkClient(cc *grpc.ClientConn) LinkClient {
	return &linkClient{cc}
}

func (c *linkClient) GetLink(ctx context.Context, in *link.Link, opts ...grpc.CallOption) (*link.Link, error) {
	out := new(link.Link)
	err := c.cc.Invoke(ctx, "/grpc_web.Link/GetLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkClient) CreateLink(ctx context.Context, in *link.Link, opts ...grpc.CallOption) (*link.Link, error) {
	out := new(link.Link)
	err := c.cc.Invoke(ctx, "/grpc_web.Link/CreateLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkClient) DeleteLink(ctx context.Context, in *link.Link, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/grpc_web.Link/DeleteLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LinkServer is the server API for Link service.
type LinkServer interface {
	GetLink(context.Context, *link.Link) (*link.Link, error)
	CreateLink(context.Context, *link.Link) (*link.Link, error)
	DeleteLink(context.Context, *link.Link) (*empty.Empty, error)
}

// UnimplementedLinkServer can be embedded to have forward compatible implementations.
type UnimplementedLinkServer struct {
}

func (*UnimplementedLinkServer) GetLink(ctx context.Context, req *link.Link) (*link.Link, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLink not implemented")
}
func (*UnimplementedLinkServer) CreateLink(ctx context.Context, req *link.Link) (*link.Link, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLink not implemented")
}
func (*UnimplementedLinkServer) DeleteLink(ctx context.Context, req *link.Link) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLink not implemented")
}

func RegisterLinkServer(s *grpc.Server, srv LinkServer) {
	s.RegisterService(&_Link_serviceDesc, srv)
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
		FullMethod: "/grpc_web.Link/GetLink",
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
		FullMethod: "/grpc_web.Link/CreateLink",
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
		FullMethod: "/grpc_web.Link/DeleteLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkServer).DeleteLink(ctx, req.(*link.Link))
	}
	return interceptor(ctx, in, info, handler)
}

var _Link_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_web.Link",
	HandlerType: (*LinkServer)(nil),
	Methods: []grpc.MethodDesc{
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
	Metadata: "pkg/api/grpc-web/api.proto",
}
