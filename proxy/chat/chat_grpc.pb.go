// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: chat.proto

package chat

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

const (
	ChatService_CreateUser_FullMethodName           = "/chat.ChatService/CreateUser"
	ChatService_GetUser_FullMethodName              = "/chat.ChatService/GetUser"
	ChatService_SendMessage_FullMethodName          = "/chat.ChatService/SendMessage"
	ChatService_GetMultipleUsers_FullMethodName     = "/chat.ChatService/GetMultipleUsers"
	ChatService_SendMultipleMessages_FullMethodName = "/chat.ChatService/SendMultipleMessages"
	ChatService_GetMessages_FullMethodName          = "/chat.ChatService/GetMessages"
)

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*Empty, error)
	GetUser(ctx context.Context, in *GetById, opts ...grpc.CallOption) (*User, error)
	SendMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Empty, error)
	GetMultipleUsers(ctx context.Context, in *Pagination, opts ...grpc.CallOption) (ChatService_GetMultipleUsersClient, error)
	SendMultipleMessages(ctx context.Context, opts ...grpc.CallOption) (ChatService_SendMultipleMessagesClient, error)
	GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (ChatService_GetMessagesClient, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, ChatService_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetUser(ctx context.Context, in *GetById, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, ChatService_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) SendMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, ChatService_SendMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetMultipleUsers(ctx context.Context, in *Pagination, opts ...grpc.CallOption) (ChatService_GetMultipleUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[0], ChatService_GetMultipleUsers_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceGetMultipleUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatService_GetMultipleUsersClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type chatServiceGetMultipleUsersClient struct {
	grpc.ClientStream
}

func (x *chatServiceGetMultipleUsersClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) SendMultipleMessages(ctx context.Context, opts ...grpc.CallOption) (ChatService_SendMultipleMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[1], ChatService_SendMultipleMessages_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceSendMultipleMessagesClient{stream}
	return x, nil
}

type ChatService_SendMultipleMessagesClient interface {
	Send(*Message) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type chatServiceSendMultipleMessagesClient struct {
	grpc.ClientStream
}

func (x *chatServiceSendMultipleMessagesClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceSendMultipleMessagesClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (ChatService_GetMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[2], ChatService_GetMessages_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceGetMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatService_GetMessagesClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type chatServiceGetMessagesClient struct {
	grpc.ClientStream
}

func (x *chatServiceGetMessagesClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*Empty, error)
	GetUser(context.Context, *GetById) (*User, error)
	SendMessage(context.Context, *Message) (*Empty, error)
	GetMultipleUsers(*Pagination, ChatService_GetMultipleUsersServer) error
	SendMultipleMessages(ChatService_SendMultipleMessagesServer) error
	GetMessages(*GetMessagesRequest, ChatService_GetMessagesServer) error
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) CreateUser(context.Context, *CreateUserRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedChatServiceServer) GetUser(context.Context, *GetById) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedChatServiceServer) SendMessage(context.Context, *Message) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedChatServiceServer) GetMultipleUsers(*Pagination, ChatService_GetMultipleUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMultipleUsers not implemented")
}
func (UnimplementedChatServiceServer) SendMultipleMessages(ChatService_SendMultipleMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method SendMultipleMessages not implemented")
}
func (UnimplementedChatServiceServer) GetMessages(*GetMessagesRequest, ChatService_GetMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetUser(ctx, req.(*GetById))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SendMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetMultipleUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Pagination)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).GetMultipleUsers(m, &chatServiceGetMultipleUsersServer{stream})
}

type ChatService_GetMultipleUsersServer interface {
	Send(*User) error
	grpc.ServerStream
}

type chatServiceGetMultipleUsersServer struct {
	grpc.ServerStream
}

func (x *chatServiceGetMultipleUsersServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func _ChatService_SendMultipleMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).SendMultipleMessages(&chatServiceSendMultipleMessagesServer{stream})
}

type ChatService_SendMultipleMessagesServer interface {
	SendAndClose(*Empty) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type chatServiceSendMultipleMessagesServer struct {
	grpc.ServerStream
}

func (x *chatServiceSendMultipleMessagesServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceSendMultipleMessagesServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ChatService_GetMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetMessagesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).GetMessages(m, &chatServiceGetMessagesServer{stream})
}

type ChatService_GetMessagesServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type chatServiceGetMessagesServer struct {
	grpc.ServerStream
}

func (x *chatServiceGetMessagesServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _ChatService_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _ChatService_GetUser_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _ChatService_SendMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetMultipleUsers",
			Handler:       _ChatService_GetMultipleUsers_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendMultipleMessages",
			Handler:       _ChatService_SendMultipleMessages_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetMessages",
			Handler:       _ChatService_GetMessages_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chat.proto",
}
