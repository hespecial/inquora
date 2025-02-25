// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: like.proto

package pb

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
	Like_ThumbUp_FullMethodName   = "/service.Like/ThumbUp"
	Like_IsThumbUp_FullMethodName = "/service.Like/IsThumbUp"
)

// LikeClient is the client API for Like service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LikeClient interface {
	ThumbUp(ctx context.Context, in *ThumbUpRequest, opts ...grpc.CallOption) (*ThumbUpResponse, error)
	IsThumbUp(ctx context.Context, in *IsThumbUpRequest, opts ...grpc.CallOption) (*IsThumbUpResponse, error)
}

type likeClient struct {
	cc grpc.ClientConnInterface
}

func NewLikeClient(cc grpc.ClientConnInterface) LikeClient {
	return &likeClient{cc}
}

func (c *likeClient) ThumbUp(ctx context.Context, in *ThumbUpRequest, opts ...grpc.CallOption) (*ThumbUpResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ThumbUpResponse)
	err := c.cc.Invoke(ctx, Like_ThumbUp_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likeClient) IsThumbUp(ctx context.Context, in *IsThumbUpRequest, opts ...grpc.CallOption) (*IsThumbUpResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IsThumbUpResponse)
	err := c.cc.Invoke(ctx, Like_IsThumbUp_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LikeServer is the server API for Like service.
// All implementations must embed UnimplementedLikeServer
// for forward compatibility.
type LikeServer interface {
	ThumbUp(context.Context, *ThumbUpRequest) (*ThumbUpResponse, error)
	IsThumbUp(context.Context, *IsThumbUpRequest) (*IsThumbUpResponse, error)
	mustEmbedUnimplementedLikeServer()
}

// UnimplementedLikeServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLikeServer struct{}

func (UnimplementedLikeServer) ThumbUp(context.Context, *ThumbUpRequest) (*ThumbUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ThumbUp not implemented")
}
func (UnimplementedLikeServer) IsThumbUp(context.Context, *IsThumbUpRequest) (*IsThumbUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsThumbUp not implemented")
}
func (UnimplementedLikeServer) mustEmbedUnimplementedLikeServer() {}
func (UnimplementedLikeServer) testEmbeddedByValue()              {}

// UnsafeLikeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LikeServer will
// result in compilation errors.
type UnsafeLikeServer interface {
	mustEmbedUnimplementedLikeServer()
}

func RegisterLikeServer(s grpc.ServiceRegistrar, srv LikeServer) {
	// If the following call pancis, it indicates UnimplementedLikeServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Like_ServiceDesc, srv)
}

func _Like_ThumbUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThumbUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServer).ThumbUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Like_ThumbUp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServer).ThumbUp(ctx, req.(*ThumbUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Like_IsThumbUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsThumbUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServer).IsThumbUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Like_IsThumbUp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServer).IsThumbUp(ctx, req.(*IsThumbUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Like_ServiceDesc is the grpc.ServiceDesc for Like service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Like_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.Like",
	HandlerType: (*LikeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ThumbUp",
			Handler:    _Like_ThumbUp_Handler,
		},
		{
			MethodName: "IsThumbUp",
			Handler:    _Like_IsThumbUp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "like.proto",
}
