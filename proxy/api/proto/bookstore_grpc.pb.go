// Copyright 2016 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
////////////////////////////////////////////////////////////////////////////////

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: bookstore.proto

package bookstore

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Bookstore_ListShelves_FullMethodName = "/endpoints.examples.bookstore.Bookstore/ListShelves"
	Bookstore_CreateShelf_FullMethodName = "/endpoints.examples.bookstore.Bookstore/CreateShelf"
	Bookstore_GetShelf_FullMethodName    = "/endpoints.examples.bookstore.Bookstore/GetShelf"
	Bookstore_DeleteShelf_FullMethodName = "/endpoints.examples.bookstore.Bookstore/DeleteShelf"
	Bookstore_ListBooks_FullMethodName   = "/endpoints.examples.bookstore.Bookstore/ListBooks"
	Bookstore_CreateBook_FullMethodName  = "/endpoints.examples.bookstore.Bookstore/CreateBook"
	Bookstore_GetBook_FullMethodName     = "/endpoints.examples.bookstore.Bookstore/GetBook"
	Bookstore_DeleteBook_FullMethodName  = "/endpoints.examples.bookstore.Bookstore/DeleteBook"
)

// BookstoreClient is the client API for Bookstore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookstoreClient interface {
	// Returns a list of all shelves in the bookstore.
	ListShelves(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListShelvesResponse, error)
	// Creates a new shelf in the bookstore.
	CreateShelf(ctx context.Context, in *CreateShelfRequest, opts ...grpc.CallOption) (*Shelf, error)
	// Returns a specific bookstore shelf.
	GetShelf(ctx context.Context, in *GetShelfRequest, opts ...grpc.CallOption) (*Shelf, error)
	// Deletes a shelf, including all books that are stored on the shelf.
	DeleteShelf(ctx context.Context, in *DeleteShelfRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Returns a list of books on a shelf.
	ListBooks(ctx context.Context, in *ListBooksRequest, opts ...grpc.CallOption) (*ListBooksResponse, error)
	// Creates a new book.
	CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*Book, error)
	// Returns a specific book.
	GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*Book, error)
	// Deletes a book from a shelf.
	DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type bookstoreClient struct {
	cc grpc.ClientConnInterface
}

func NewBookstoreClient(cc grpc.ClientConnInterface) BookstoreClient {
	return &bookstoreClient{cc}
}

func (c *bookstoreClient) ListShelves(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListShelvesResponse, error) {
	out := new(ListShelvesResponse)
	err := c.cc.Invoke(ctx, Bookstore_ListShelves_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) CreateShelf(ctx context.Context, in *CreateShelfRequest, opts ...grpc.CallOption) (*Shelf, error) {
	out := new(Shelf)
	err := c.cc.Invoke(ctx, Bookstore_CreateShelf_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) GetShelf(ctx context.Context, in *GetShelfRequest, opts ...grpc.CallOption) (*Shelf, error) {
	out := new(Shelf)
	err := c.cc.Invoke(ctx, Bookstore_GetShelf_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) DeleteShelf(ctx context.Context, in *DeleteShelfRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Bookstore_DeleteShelf_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) ListBooks(ctx context.Context, in *ListBooksRequest, opts ...grpc.CallOption) (*ListBooksResponse, error) {
	out := new(ListBooksResponse)
	err := c.cc.Invoke(ctx, Bookstore_ListBooks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, Bookstore_CreateBook_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, Bookstore_GetBook_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Bookstore_DeleteBook_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookstoreServer is the server API for Bookstore service.
// All implementations must embed UnimplementedBookstoreServer
// for forward compatibility
type BookstoreServer interface {
	// Returns a list of all shelves in the bookstore.
	ListShelves(context.Context, *emptypb.Empty) (*ListShelvesResponse, error)
	// Creates a new shelf in the bookstore.
	CreateShelf(context.Context, *CreateShelfRequest) (*Shelf, error)
	// Returns a specific bookstore shelf.
	GetShelf(context.Context, *GetShelfRequest) (*Shelf, error)
	// Deletes a shelf, including all books that are stored on the shelf.
	DeleteShelf(context.Context, *DeleteShelfRequest) (*emptypb.Empty, error)
	// Returns a list of books on a shelf.
	ListBooks(context.Context, *ListBooksRequest) (*ListBooksResponse, error)
	// Creates a new book.
	CreateBook(context.Context, *CreateBookRequest) (*Book, error)
	// Returns a specific book.
	GetBook(context.Context, *GetBookRequest) (*Book, error)
	// Deletes a book from a shelf.
	DeleteBook(context.Context, *DeleteBookRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedBookstoreServer()
}

// UnimplementedBookstoreServer must be embedded to have forward compatible implementations.
type UnimplementedBookstoreServer struct {
}

func (UnimplementedBookstoreServer) ListShelves(context.Context, *emptypb.Empty) (*ListShelvesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListShelves not implemented")
}
func (UnimplementedBookstoreServer) CreateShelf(context.Context, *CreateShelfRequest) (*Shelf, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShelf not implemented")
}
func (UnimplementedBookstoreServer) GetShelf(context.Context, *GetShelfRequest) (*Shelf, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShelf not implemented")
}
func (UnimplementedBookstoreServer) DeleteShelf(context.Context, *DeleteShelfRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteShelf not implemented")
}
func (UnimplementedBookstoreServer) ListBooks(context.Context, *ListBooksRequest) (*ListBooksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBooks not implemented")
}
func (UnimplementedBookstoreServer) CreateBook(context.Context, *CreateBookRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}
func (UnimplementedBookstoreServer) GetBook(context.Context, *GetBookRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (UnimplementedBookstoreServer) DeleteBook(context.Context, *DeleteBookRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBook not implemented")
}
func (UnimplementedBookstoreServer) mustEmbedUnimplementedBookstoreServer() {}

// UnsafeBookstoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookstoreServer will
// result in compilation errors.
type UnsafeBookstoreServer interface {
	mustEmbedUnimplementedBookstoreServer()
}

func RegisterBookstoreServer(s grpc.ServiceRegistrar, srv BookstoreServer) {
	s.RegisterService(&Bookstore_ServiceDesc, srv)
}

func _Bookstore_ListShelves_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).ListShelves(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bookstore_ListShelves_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).ListShelves(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_CreateShelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).CreateShelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bookstore_CreateShelf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).CreateShelf(ctx, req.(*CreateShelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_GetShelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).GetShelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bookstore_GetShelf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).GetShelf(ctx, req.(*GetShelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_DeleteShelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteShelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).DeleteShelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bookstore_DeleteShelf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).DeleteShelf(ctx, req.(*DeleteShelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_ListBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).ListBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bookstore_ListBooks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).ListBooks(ctx, req.(*ListBooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bookstore_CreateBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).CreateBook(ctx, req.(*CreateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bookstore_GetBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).GetBook(ctx, req.(*GetBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bookstore_DeleteBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).DeleteBook(ctx, req.(*DeleteBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Bookstore_ServiceDesc is the grpc.ServiceDesc for Bookstore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bookstore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "endpoints.examples.bookstore.Bookstore",
	HandlerType: (*BookstoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListShelves",
			Handler:    _Bookstore_ListShelves_Handler,
		},
		{
			MethodName: "CreateShelf",
			Handler:    _Bookstore_CreateShelf_Handler,
		},
		{
			MethodName: "GetShelf",
			Handler:    _Bookstore_GetShelf_Handler,
		},
		{
			MethodName: "DeleteShelf",
			Handler:    _Bookstore_DeleteShelf_Handler,
		},
		{
			MethodName: "ListBooks",
			Handler:    _Bookstore_ListBooks_Handler,
		},
		{
			MethodName: "CreateBook",
			Handler:    _Bookstore_CreateBook_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _Bookstore_GetBook_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _Bookstore_DeleteBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bookstore.proto",
}
