// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: book_pro_service.proto

package book_pro_service

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	BookProService_CreateBook_FullMethodName  = "/book_pro_service.BookProService/CreateBook"
	BookProService_UpdateBook_FullMethodName  = "/book_pro_service.BookProService/UpdateBook"
	BookProService_GetBookList_FullMethodName = "/book_pro_service.BookProService/GetBookList"
	BookProService_GetBook_FullMethodName     = "/book_pro_service.BookProService/GetBook"
	BookProService_DeleteBook_FullMethodName  = "/book_pro_service.BookProService/DeleteBook"
)

// BookProServiceClient is the client API for BookProService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookProServiceClient interface {
	CreateBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*empty.Empty, error)
	GetBookList(ctx context.Context, in *GetBookListRequest, opts ...grpc.CallOption) (*GetBookListResponse, error)
	GetBook(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Book, error)
	DeleteBook(ctx context.Context, in *ById, opts ...grpc.CallOption) (*empty.Empty, error)
}

type bookProServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookProServiceClient(cc grpc.ClientConnInterface) BookProServiceClient {
	return &bookProServiceClient{cc}
}

func (c *bookProServiceClient) CreateBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, BookProService_CreateBook_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookProServiceClient) UpdateBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, BookProService_UpdateBook_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookProServiceClient) GetBookList(ctx context.Context, in *GetBookListRequest, opts ...grpc.CallOption) (*GetBookListResponse, error) {
	out := new(GetBookListResponse)
	err := c.cc.Invoke(ctx, BookProService_GetBookList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookProServiceClient) GetBook(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, BookProService_GetBook_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookProServiceClient) DeleteBook(ctx context.Context, in *ById, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, BookProService_DeleteBook_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookProServiceServer is the server API for BookProService service.
// All implementations must embed UnimplementedBookProServiceServer
// for forward compatibility
type BookProServiceServer interface {
	CreateBook(context.Context, *Book) (*empty.Empty, error)
	UpdateBook(context.Context, *Book) (*empty.Empty, error)
	GetBookList(context.Context, *GetBookListRequest) (*GetBookListResponse, error)
	GetBook(context.Context, *ById) (*Book, error)
	DeleteBook(context.Context, *ById) (*empty.Empty, error)
	mustEmbedUnimplementedBookProServiceServer()
}

// UnimplementedBookProServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBookProServiceServer struct {
}

func (UnimplementedBookProServiceServer) CreateBook(context.Context, *Book) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}
func (UnimplementedBookProServiceServer) UpdateBook(context.Context, *Book) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBook not implemented")
}
func (UnimplementedBookProServiceServer) GetBookList(context.Context, *GetBookListRequest) (*GetBookListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookList not implemented")
}
func (UnimplementedBookProServiceServer) GetBook(context.Context, *ById) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (UnimplementedBookProServiceServer) DeleteBook(context.Context, *ById) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBook not implemented")
}
func (UnimplementedBookProServiceServer) mustEmbedUnimplementedBookProServiceServer() {}

// UnsafeBookProServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookProServiceServer will
// result in compilation errors.
type UnsafeBookProServiceServer interface {
	mustEmbedUnimplementedBookProServiceServer()
}

func RegisterBookProServiceServer(s grpc.ServiceRegistrar, srv BookProServiceServer) {
	s.RegisterService(&BookProService_ServiceDesc, srv)
}

func _BookProService_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookProServiceServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookProService_CreateBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookProServiceServer).CreateBook(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookProService_UpdateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookProServiceServer).UpdateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookProService_UpdateBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookProServiceServer).UpdateBook(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookProService_GetBookList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookProServiceServer).GetBookList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookProService_GetBookList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookProServiceServer).GetBookList(ctx, req.(*GetBookListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookProService_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookProServiceServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookProService_GetBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookProServiceServer).GetBook(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookProService_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookProServiceServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookProService_DeleteBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookProServiceServer).DeleteBook(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

// BookProService_ServiceDesc is the grpc.ServiceDesc for BookProService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookProService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "book_pro_service.BookProService",
	HandlerType: (*BookProServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBook",
			Handler:    _BookProService_CreateBook_Handler,
		},
		{
			MethodName: "UpdateBook",
			Handler:    _BookProService_UpdateBook_Handler,
		},
		{
			MethodName: "GetBookList",
			Handler:    _BookProService_GetBookList_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _BookProService_GetBook_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _BookProService_DeleteBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "book_pro_service.proto",
}
