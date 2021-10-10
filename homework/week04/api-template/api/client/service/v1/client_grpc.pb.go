// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// ClientClient is the client API for Client service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientClient interface {
	GetClient(ctx context.Context, in *GetClientReq, opts ...grpc.CallOption) (*GetClientReply, error)
	CreateClient(ctx context.Context, in *CreateClientReq, opts ...grpc.CallOption) (*CreateClientReply, error)
	UpdateClient(ctx context.Context, in *UpdateClientReq, opts ...grpc.CallOption) (*UpdateClientReply, error)
	ListClient(ctx context.Context, in *ListClientReq, opts ...grpc.CallOption) (*ListClientReply, error)
	DeleteClient(ctx context.Context, in *DeleteClientReq, opts ...grpc.CallOption) (*DeleteClientReply, error)
	SearchClient(ctx context.Context, in *SearchClientReq, opts ...grpc.CallOption) (*SearchClientReply, error)
}

type clientClient struct {
	cc grpc.ClientConnInterface
}

func NewClientClient(cc grpc.ClientConnInterface) ClientClient {
	return &clientClient{cc}
}

func (c *clientClient) GetClient(ctx context.Context, in *GetClientReq, opts ...grpc.CallOption) (*GetClientReply, error) {
	out := new(GetClientReply)
	err := c.cc.Invoke(ctx, "/client.service.v1.Client/GetClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) CreateClient(ctx context.Context, in *CreateClientReq, opts ...grpc.CallOption) (*CreateClientReply, error) {
	out := new(CreateClientReply)
	err := c.cc.Invoke(ctx, "/client.service.v1.Client/CreateClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) UpdateClient(ctx context.Context, in *UpdateClientReq, opts ...grpc.CallOption) (*UpdateClientReply, error) {
	out := new(UpdateClientReply)
	err := c.cc.Invoke(ctx, "/client.service.v1.Client/UpdateClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) ListClient(ctx context.Context, in *ListClientReq, opts ...grpc.CallOption) (*ListClientReply, error) {
	out := new(ListClientReply)
	err := c.cc.Invoke(ctx, "/client.service.v1.Client/ListClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) DeleteClient(ctx context.Context, in *DeleteClientReq, opts ...grpc.CallOption) (*DeleteClientReply, error) {
	out := new(DeleteClientReply)
	err := c.cc.Invoke(ctx, "/client.service.v1.Client/DeleteClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) SearchClient(ctx context.Context, in *SearchClientReq, opts ...grpc.CallOption) (*SearchClientReply, error) {
	out := new(SearchClientReply)
	err := c.cc.Invoke(ctx, "/client.service.v1.Client/SearchClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServer is the server API for Client service.
// All implementations must embed UnimplementedClientServer
// for forward compatibility
type ClientServer interface {
	GetClient(context.Context, *GetClientReq) (*GetClientReply, error)
	CreateClient(context.Context, *CreateClientReq) (*CreateClientReply, error)
	UpdateClient(context.Context, *UpdateClientReq) (*UpdateClientReply, error)
	ListClient(context.Context, *ListClientReq) (*ListClientReply, error)
	DeleteClient(context.Context, *DeleteClientReq) (*DeleteClientReply, error)
	SearchClient(context.Context, *SearchClientReq) (*SearchClientReply, error)
	mustEmbedUnimplementedClientServer()
}

// UnimplementedClientServer must be embedded to have forward compatible implementations.
type UnimplementedClientServer struct {
}

func (UnimplementedClientServer) GetClient(context.Context, *GetClientReq) (*GetClientReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClient not implemented")
}
func (UnimplementedClientServer) CreateClient(context.Context, *CreateClientReq) (*CreateClientReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClient not implemented")
}
func (UnimplementedClientServer) UpdateClient(context.Context, *UpdateClientReq) (*UpdateClientReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClient not implemented")
}
func (UnimplementedClientServer) ListClient(context.Context, *ListClientReq) (*ListClientReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListClient not implemented")
}
func (UnimplementedClientServer) DeleteClient(context.Context, *DeleteClientReq) (*DeleteClientReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClient not implemented")
}
func (UnimplementedClientServer) SearchClient(context.Context, *SearchClientReq) (*SearchClientReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchClient not implemented")
}
func (UnimplementedClientServer) mustEmbedUnimplementedClientServer() {}

// UnsafeClientServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientServer will
// result in compilation errors.
type UnsafeClientServer interface {
	mustEmbedUnimplementedClientServer()
}

func RegisterClientServer(s grpc.ServiceRegistrar, srv ClientServer) {
	s.RegisterService(&Client_ServiceDesc, srv)
}

func _Client_GetClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).GetClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.service.v1.Client/GetClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).GetClient(ctx, req.(*GetClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_CreateClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).CreateClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.service.v1.Client/CreateClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).CreateClient(ctx, req.(*CreateClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_UpdateClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).UpdateClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.service.v1.Client/UpdateClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).UpdateClient(ctx, req.(*UpdateClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_ListClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).ListClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.service.v1.Client/ListClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).ListClient(ctx, req.(*ListClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_DeleteClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).DeleteClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.service.v1.Client/DeleteClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).DeleteClient(ctx, req.(*DeleteClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_SearchClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).SearchClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.service.v1.Client/SearchClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).SearchClient(ctx, req.(*SearchClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Client_ServiceDesc is the grpc.ServiceDesc for Client service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Client_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "client.service.v1.Client",
	HandlerType: (*ClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetClient",
			Handler:    _Client_GetClient_Handler,
		},
		{
			MethodName: "CreateClient",
			Handler:    _Client_CreateClient_Handler,
		},
		{
			MethodName: "UpdateClient",
			Handler:    _Client_UpdateClient_Handler,
		},
		{
			MethodName: "ListClient",
			Handler:    _Client_ListClient_Handler,
		},
		{
			MethodName: "DeleteClient",
			Handler:    _Client_DeleteClient_Handler,
		},
		{
			MethodName: "SearchClient",
			Handler:    _Client_SearchClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/client/service/v1/client.proto",
}