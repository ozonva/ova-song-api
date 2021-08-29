// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ova_song_api

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

// OvaSongApiClient is the client API for OvaSongApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OvaSongApiClient interface {
	CreateSongV1(ctx context.Context, in *CreateSongV1Request, opts ...grpc.CallOption) (*CreateSongV1Response, error)
	DescribeSongV1(ctx context.Context, in *DescribeSongV1Request, opts ...grpc.CallOption) (*DescribeSongV1Response, error)
	ListSongsV1(ctx context.Context, in *ListSongsV1Request, opts ...grpc.CallOption) (*ListSongsV1Response, error)
	RemoveSongV1(ctx context.Context, in *RemoveSongV1Request, opts ...grpc.CallOption) (*RemoveSongV1Response, error)
}

type ovaSongApiClient struct {
	cc grpc.ClientConnInterface
}

func NewOvaSongApiClient(cc grpc.ClientConnInterface) OvaSongApiClient {
	return &ovaSongApiClient{cc}
}

func (c *ovaSongApiClient) CreateSongV1(ctx context.Context, in *CreateSongV1Request, opts ...grpc.CallOption) (*CreateSongV1Response, error) {
	out := new(CreateSongV1Response)
	err := c.cc.Invoke(ctx, "/ova.task.api.OvaSongApi/CreateSongV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ovaSongApiClient) DescribeSongV1(ctx context.Context, in *DescribeSongV1Request, opts ...grpc.CallOption) (*DescribeSongV1Response, error) {
	out := new(DescribeSongV1Response)
	err := c.cc.Invoke(ctx, "/ova.task.api.OvaSongApi/DescribeSongV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ovaSongApiClient) ListSongsV1(ctx context.Context, in *ListSongsV1Request, opts ...grpc.CallOption) (*ListSongsV1Response, error) {
	out := new(ListSongsV1Response)
	err := c.cc.Invoke(ctx, "/ova.task.api.OvaSongApi/ListSongsV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ovaSongApiClient) RemoveSongV1(ctx context.Context, in *RemoveSongV1Request, opts ...grpc.CallOption) (*RemoveSongV1Response, error) {
	out := new(RemoveSongV1Response)
	err := c.cc.Invoke(ctx, "/ova.task.api.OvaSongApi/RemoveSongV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OvaSongApiServer is the server API for OvaSongApi service.
// All implementations must embed UnimplementedOvaSongApiServer
// for forward compatibility
type OvaSongApiServer interface {
	CreateSongV1(context.Context, *CreateSongV1Request) (*CreateSongV1Response, error)
	DescribeSongV1(context.Context, *DescribeSongV1Request) (*DescribeSongV1Response, error)
	ListSongsV1(context.Context, *ListSongsV1Request) (*ListSongsV1Response, error)
	RemoveSongV1(context.Context, *RemoveSongV1Request) (*RemoveSongV1Response, error)
	mustEmbedUnimplementedOvaSongApiServer()
}

// UnimplementedOvaSongApiServer must be embedded to have forward compatible implementations.
type UnimplementedOvaSongApiServer struct {
}

func (UnimplementedOvaSongApiServer) CreateSongV1(context.Context, *CreateSongV1Request) (*CreateSongV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSongV1 not implemented")
}
func (UnimplementedOvaSongApiServer) DescribeSongV1(context.Context, *DescribeSongV1Request) (*DescribeSongV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeSongV1 not implemented")
}
func (UnimplementedOvaSongApiServer) ListSongsV1(context.Context, *ListSongsV1Request) (*ListSongsV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSongsV1 not implemented")
}
func (UnimplementedOvaSongApiServer) RemoveSongV1(context.Context, *RemoveSongV1Request) (*RemoveSongV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveSongV1 not implemented")
}
func (UnimplementedOvaSongApiServer) mustEmbedUnimplementedOvaSongApiServer() {}

// UnsafeOvaSongApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OvaSongApiServer will
// result in compilation errors.
type UnsafeOvaSongApiServer interface {
	mustEmbedUnimplementedOvaSongApiServer()
}

func RegisterOvaSongApiServer(s grpc.ServiceRegistrar, srv OvaSongApiServer) {
	s.RegisterService(&OvaSongApi_ServiceDesc, srv)
}

func _OvaSongApi_CreateSongV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSongV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OvaSongApiServer).CreateSongV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.task.api.OvaSongApi/CreateSongV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OvaSongApiServer).CreateSongV1(ctx, req.(*CreateSongV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OvaSongApi_DescribeSongV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeSongV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OvaSongApiServer).DescribeSongV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.task.api.OvaSongApi/DescribeSongV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OvaSongApiServer).DescribeSongV1(ctx, req.(*DescribeSongV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OvaSongApi_ListSongsV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSongsV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OvaSongApiServer).ListSongsV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.task.api.OvaSongApi/ListSongsV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OvaSongApiServer).ListSongsV1(ctx, req.(*ListSongsV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OvaSongApi_RemoveSongV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveSongV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OvaSongApiServer).RemoveSongV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.task.api.OvaSongApi/RemoveSongV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OvaSongApiServer).RemoveSongV1(ctx, req.(*RemoveSongV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// OvaSongApi_ServiceDesc is the grpc.ServiceDesc for OvaSongApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OvaSongApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ova.task.api.OvaSongApi",
	HandlerType: (*OvaSongApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSongV1",
			Handler:    _OvaSongApi_CreateSongV1_Handler,
		},
		{
			MethodName: "DescribeSongV1",
			Handler:    _OvaSongApi_DescribeSongV1_Handler,
		},
		{
			MethodName: "ListSongsV1",
			Handler:    _OvaSongApi_ListSongsV1_Handler,
		},
		{
			MethodName: "RemoveSongV1",
			Handler:    _OvaSongApi_RemoveSongV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ova-song-api.proto",
}
