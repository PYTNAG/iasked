// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.0
// source: comment/service_comment.proto

package pb

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

// CommentAPIClient is the client API for CommentAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentAPIClient interface {
	CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error)
	GetRFCComments(ctx context.Context, in *GetRFCCommentsRequest, opts ...grpc.CallOption) (CommentAPI_GetRFCCommentsClient, error)
	CountRFCComments(ctx context.Context, in *CountRFCCommentsRequest, opts ...grpc.CallOption) (*CountRFCCommentsResponse, error)
	DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type commentAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentAPIClient(cc grpc.ClientConnInterface) CommentAPIClient {
	return &commentAPIClient{cc}
}

func (c *commentAPIClient) CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error) {
	out := new(CreateCommentResponse)
	err := c.cc.Invoke(ctx, "/pb.CommentAPI/CreateComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentAPIClient) GetRFCComments(ctx context.Context, in *GetRFCCommentsRequest, opts ...grpc.CallOption) (CommentAPI_GetRFCCommentsClient, error) {
	stream, err := c.cc.NewStream(ctx, &CommentAPI_ServiceDesc.Streams[0], "/pb.CommentAPI/GetRFCComments", opts...)
	if err != nil {
		return nil, err
	}
	x := &commentAPIGetRFCCommentsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CommentAPI_GetRFCCommentsClient interface {
	Recv() (*Comment, error)
	grpc.ClientStream
}

type commentAPIGetRFCCommentsClient struct {
	grpc.ClientStream
}

func (x *commentAPIGetRFCCommentsClient) Recv() (*Comment, error) {
	m := new(Comment)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *commentAPIClient) CountRFCComments(ctx context.Context, in *CountRFCCommentsRequest, opts ...grpc.CallOption) (*CountRFCCommentsResponse, error) {
	out := new(CountRFCCommentsResponse)
	err := c.cc.Invoke(ctx, "/pb.CommentAPI/CountRFCComments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentAPIClient) DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.CommentAPI/DeleteComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentAPIServer is the server API for CommentAPI service.
// All implementations must embed UnimplementedCommentAPIServer
// for forward compatibility
type CommentAPIServer interface {
	CreateComment(context.Context, *CreateCommentRequest) (*CreateCommentResponse, error)
	GetRFCComments(*GetRFCCommentsRequest, CommentAPI_GetRFCCommentsServer) error
	CountRFCComments(context.Context, *CountRFCCommentsRequest) (*CountRFCCommentsResponse, error)
	DeleteComment(context.Context, *DeleteCommentRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedCommentAPIServer()
}

// UnimplementedCommentAPIServer must be embedded to have forward compatible implementations.
type UnimplementedCommentAPIServer struct {
}

func (UnimplementedCommentAPIServer) CreateComment(context.Context, *CreateCommentRequest) (*CreateCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedCommentAPIServer) GetRFCComments(*GetRFCCommentsRequest, CommentAPI_GetRFCCommentsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetRFCComments not implemented")
}
func (UnimplementedCommentAPIServer) CountRFCComments(context.Context, *CountRFCCommentsRequest) (*CountRFCCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountRFCComments not implemented")
}
func (UnimplementedCommentAPIServer) DeleteComment(context.Context, *DeleteCommentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedCommentAPIServer) mustEmbedUnimplementedCommentAPIServer() {}

// UnsafeCommentAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentAPIServer will
// result in compilation errors.
type UnsafeCommentAPIServer interface {
	mustEmbedUnimplementedCommentAPIServer()
}

func RegisterCommentAPIServer(s grpc.ServiceRegistrar, srv CommentAPIServer) {
	s.RegisterService(&CommentAPI_ServiceDesc, srv)
}

func _CommentAPI_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentAPIServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CommentAPI/CreateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentAPIServer).CreateComment(ctx, req.(*CreateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentAPI_GetRFCComments_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetRFCCommentsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CommentAPIServer).GetRFCComments(m, &commentAPIGetRFCCommentsServer{stream})
}

type CommentAPI_GetRFCCommentsServer interface {
	Send(*Comment) error
	grpc.ServerStream
}

type commentAPIGetRFCCommentsServer struct {
	grpc.ServerStream
}

func (x *commentAPIGetRFCCommentsServer) Send(m *Comment) error {
	return x.ServerStream.SendMsg(m)
}

func _CommentAPI_CountRFCComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountRFCCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentAPIServer).CountRFCComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CommentAPI/CountRFCComments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentAPIServer).CountRFCComments(ctx, req.(*CountRFCCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentAPI_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentAPIServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CommentAPI/DeleteComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentAPIServer).DeleteComment(ctx, req.(*DeleteCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommentAPI_ServiceDesc is the grpc.ServiceDesc for CommentAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommentAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CommentAPI",
	HandlerType: (*CommentAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateComment",
			Handler:    _CommentAPI_CreateComment_Handler,
		},
		{
			MethodName: "CountRFCComments",
			Handler:    _CommentAPI_CountRFCComments_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _CommentAPI_DeleteComment_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetRFCComments",
			Handler:       _CommentAPI_GetRFCComments_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "comment/service_comment.proto",
}
