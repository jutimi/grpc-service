// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: oauth/user.proto

package oauth

import (
	context "context"
	common "github.com/jutimi/grpc-service/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserRouteClient is the client API for UserRoute service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserRouteClient interface {
	GetUserById(ctx context.Context, in *common.GetByIdParams, opts ...grpc.CallOption) (*UserResponse, error)
	GetUserByFilter(ctx context.Context, in *GetUserByFilterParams, opts ...grpc.CallOption) (*UserResponse, error)
	GetUsersByFilter(ctx context.Context, in *GetUserByFilterParams, opts ...grpc.CallOption) (*UsersResponse, error)
	CreateUser(ctx context.Context, in *CreateUserParams, opts ...grpc.CallOption) (*common.SuccessResponse, error)
	BulkCreateUsers(ctx context.Context, opts ...grpc.CallOption) (UserRoute_BulkCreateUsersClient, error)
}

type userRouteClient struct {
	cc grpc.ClientConnInterface
}

func NewUserRouteClient(cc grpc.ClientConnInterface) UserRouteClient {
	return &userRouteClient{cc}
}

func (c *userRouteClient) GetUserById(ctx context.Context, in *common.GetByIdParams, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/oauth.UserRoute/GetUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRouteClient) GetUserByFilter(ctx context.Context, in *GetUserByFilterParams, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/oauth.UserRoute/GetUserByFilter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRouteClient) GetUsersByFilter(ctx context.Context, in *GetUserByFilterParams, opts ...grpc.CallOption) (*UsersResponse, error) {
	out := new(UsersResponse)
	err := c.cc.Invoke(ctx, "/oauth.UserRoute/GetUsersByFilter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRouteClient) CreateUser(ctx context.Context, in *CreateUserParams, opts ...grpc.CallOption) (*common.SuccessResponse, error) {
	out := new(common.SuccessResponse)
	err := c.cc.Invoke(ctx, "/oauth.UserRoute/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRouteClient) BulkCreateUsers(ctx context.Context, opts ...grpc.CallOption) (UserRoute_BulkCreateUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserRoute_ServiceDesc.Streams[0], "/oauth.UserRoute/BulkCreateUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &userRouteBulkCreateUsersClient{stream}
	return x, nil
}

type UserRoute_BulkCreateUsersClient interface {
	Send(*FileChunk) error
	CloseAndRecv() (*UploadResponse, error)
	grpc.ClientStream
}

type userRouteBulkCreateUsersClient struct {
	grpc.ClientStream
}

func (x *userRouteBulkCreateUsersClient) Send(m *FileChunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userRouteBulkCreateUsersClient) CloseAndRecv() (*UploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserRouteServer is the server API for UserRoute service.
// All implementations must embed UnimplementedUserRouteServer
// for forward compatibility
type UserRouteServer interface {
	GetUserById(context.Context, *common.GetByIdParams) (*UserResponse, error)
	GetUserByFilter(context.Context, *GetUserByFilterParams) (*UserResponse, error)
	GetUsersByFilter(context.Context, *GetUserByFilterParams) (*UsersResponse, error)
	CreateUser(context.Context, *CreateUserParams) (*common.SuccessResponse, error)
	BulkCreateUsers(UserRoute_BulkCreateUsersServer) error
	mustEmbedUnimplementedUserRouteServer()
}

// UnimplementedUserRouteServer must be embedded to have forward compatible implementations.
type UnimplementedUserRouteServer struct {
}

func (UnimplementedUserRouteServer) GetUserById(context.Context, *common.GetByIdParams) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedUserRouteServer) GetUserByFilter(context.Context, *GetUserByFilterParams) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByFilter not implemented")
}
func (UnimplementedUserRouteServer) GetUsersByFilter(context.Context, *GetUserByFilterParams) (*UsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersByFilter not implemented")
}
func (UnimplementedUserRouteServer) CreateUser(context.Context, *CreateUserParams) (*common.SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserRouteServer) BulkCreateUsers(UserRoute_BulkCreateUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method BulkCreateUsers not implemented")
}
func (UnimplementedUserRouteServer) mustEmbedUnimplementedUserRouteServer() {}

// UnsafeUserRouteServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserRouteServer will
// result in compilation errors.
type UnsafeUserRouteServer interface {
	mustEmbedUnimplementedUserRouteServer()
}

func RegisterUserRouteServer(s grpc.ServiceRegistrar, srv UserRouteServer) {
	s.RegisterService(&UserRoute_ServiceDesc, srv)
}

func _UserRoute_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.GetByIdParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRouteServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oauth.UserRoute/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRouteServer).GetUserById(ctx, req.(*common.GetByIdParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRoute_GetUserByFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByFilterParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRouteServer).GetUserByFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oauth.UserRoute/GetUserByFilter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRouteServer).GetUserByFilter(ctx, req.(*GetUserByFilterParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRoute_GetUsersByFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByFilterParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRouteServer).GetUsersByFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oauth.UserRoute/GetUsersByFilter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRouteServer).GetUsersByFilter(ctx, req.(*GetUserByFilterParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRoute_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRouteServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oauth.UserRoute/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRouteServer).CreateUser(ctx, req.(*CreateUserParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRoute_BulkCreateUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UserRouteServer).BulkCreateUsers(&userRouteBulkCreateUsersServer{stream})
}

type UserRoute_BulkCreateUsersServer interface {
	SendAndClose(*UploadResponse) error
	Recv() (*FileChunk, error)
	grpc.ServerStream
}

type userRouteBulkCreateUsersServer struct {
	grpc.ServerStream
}

func (x *userRouteBulkCreateUsersServer) SendAndClose(m *UploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userRouteBulkCreateUsersServer) Recv() (*FileChunk, error) {
	m := new(FileChunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserRoute_ServiceDesc is the grpc.ServiceDesc for UserRoute service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserRoute_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "oauth.UserRoute",
	HandlerType: (*UserRouteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserById",
			Handler:    _UserRoute_GetUserById_Handler,
		},
		{
			MethodName: "GetUserByFilter",
			Handler:    _UserRoute_GetUserByFilter_Handler,
		},
		{
			MethodName: "GetUsersByFilter",
			Handler:    _UserRoute_GetUsersByFilter_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserRoute_CreateUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BulkCreateUsers",
			Handler:       _UserRoute_BulkCreateUsers_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "oauth/user.proto",
}
