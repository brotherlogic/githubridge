// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: githubridge.proto

package proto

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

// GithubridgeServiceClient is the client API for GithubridgeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GithubridgeServiceClient interface {
	AddLabel(ctx context.Context, in *AddLabelRequest, opts ...grpc.CallOption) (*AddLabelResponse, error)
	CreateIssue(ctx context.Context, in *CreateIssueRequest, opts ...grpc.CallOption) (*CreateIssueResponse, error)
	GetIssue(ctx context.Context, in *GetIssueRequest, opts ...grpc.CallOption) (*GetIssueResponse, error)
	GetIssues(ctx context.Context, in *GetIssuesRequest, opts ...grpc.CallOption) (*GetIssuesResponse, error)
	CloseIssue(ctx context.Context, in *CloseIssueRequest, opts ...grpc.CallOption) (*CloseIssueResponse, error)
	CommentOnIssue(ctx context.Context, in *CommentOnIssueRequest, opts ...grpc.CallOption) (*CommentOnIssueResponse, error)
	GetComments(ctx context.Context, in *GetCommentsRequest, opts ...grpc.CallOption) (*GetCommentsResponse, error)
	GetLabels(ctx context.Context, in *GetLabelsRequest, opts ...grpc.CallOption) (*GetLabelsResponse, error)
}

type githubridgeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGithubridgeServiceClient(cc grpc.ClientConnInterface) GithubridgeServiceClient {
	return &githubridgeServiceClient{cc}
}

func (c *githubridgeServiceClient) AddLabel(ctx context.Context, in *AddLabelRequest, opts ...grpc.CallOption) (*AddLabelResponse, error) {
	out := new(AddLabelResponse)
	err := c.cc.Invoke(ctx, "/githubridge.GithubridgeService/AddLabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *githubridgeServiceClient) CreateIssue(ctx context.Context, in *CreateIssueRequest, opts ...grpc.CallOption) (*CreateIssueResponse, error) {
	out := new(CreateIssueResponse)
	err := c.cc.Invoke(ctx, "/githubridge.GithubridgeService/CreateIssue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *githubridgeServiceClient) GetIssue(ctx context.Context, in *GetIssueRequest, opts ...grpc.CallOption) (*GetIssueResponse, error) {
	out := new(GetIssueResponse)
	err := c.cc.Invoke(ctx, "/githubridge.GithubridgeService/GetIssue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *githubridgeServiceClient) GetIssues(ctx context.Context, in *GetIssuesRequest, opts ...grpc.CallOption) (*GetIssuesResponse, error) {
	out := new(GetIssuesResponse)
	err := c.cc.Invoke(ctx, "/githubridge.GithubridgeService/GetIssues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *githubridgeServiceClient) CloseIssue(ctx context.Context, in *CloseIssueRequest, opts ...grpc.CallOption) (*CloseIssueResponse, error) {
	out := new(CloseIssueResponse)
	err := c.cc.Invoke(ctx, "/githubridge.GithubridgeService/CloseIssue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *githubridgeServiceClient) CommentOnIssue(ctx context.Context, in *CommentOnIssueRequest, opts ...grpc.CallOption) (*CommentOnIssueResponse, error) {
	out := new(CommentOnIssueResponse)
	err := c.cc.Invoke(ctx, "/githubridge.GithubridgeService/CommentOnIssue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *githubridgeServiceClient) GetComments(ctx context.Context, in *GetCommentsRequest, opts ...grpc.CallOption) (*GetCommentsResponse, error) {
	out := new(GetCommentsResponse)
	err := c.cc.Invoke(ctx, "/githubridge.GithubridgeService/GetComments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *githubridgeServiceClient) GetLabels(ctx context.Context, in *GetLabelsRequest, opts ...grpc.CallOption) (*GetLabelsResponse, error) {
	out := new(GetLabelsResponse)
	err := c.cc.Invoke(ctx, "/githubridge.GithubridgeService/GetLabels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GithubridgeServiceServer is the server API for GithubridgeService service.
// All implementations should embed UnimplementedGithubridgeServiceServer
// for forward compatibility
type GithubridgeServiceServer interface {
	AddLabel(context.Context, *AddLabelRequest) (*AddLabelResponse, error)
	CreateIssue(context.Context, *CreateIssueRequest) (*CreateIssueResponse, error)
	GetIssue(context.Context, *GetIssueRequest) (*GetIssueResponse, error)
	GetIssues(context.Context, *GetIssuesRequest) (*GetIssuesResponse, error)
	CloseIssue(context.Context, *CloseIssueRequest) (*CloseIssueResponse, error)
	CommentOnIssue(context.Context, *CommentOnIssueRequest) (*CommentOnIssueResponse, error)
	GetComments(context.Context, *GetCommentsRequest) (*GetCommentsResponse, error)
	GetLabels(context.Context, *GetLabelsRequest) (*GetLabelsResponse, error)
}

// UnimplementedGithubridgeServiceServer should be embedded to have forward compatible implementations.
type UnimplementedGithubridgeServiceServer struct {
}

func (UnimplementedGithubridgeServiceServer) AddLabel(context.Context, *AddLabelRequest) (*AddLabelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddLabel not implemented")
}
func (UnimplementedGithubridgeServiceServer) CreateIssue(context.Context, *CreateIssueRequest) (*CreateIssueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIssue not implemented")
}
func (UnimplementedGithubridgeServiceServer) GetIssue(context.Context, *GetIssueRequest) (*GetIssueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIssue not implemented")
}
func (UnimplementedGithubridgeServiceServer) GetIssues(context.Context, *GetIssuesRequest) (*GetIssuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIssues not implemented")
}
func (UnimplementedGithubridgeServiceServer) CloseIssue(context.Context, *CloseIssueRequest) (*CloseIssueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseIssue not implemented")
}
func (UnimplementedGithubridgeServiceServer) CommentOnIssue(context.Context, *CommentOnIssueRequest) (*CommentOnIssueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentOnIssue not implemented")
}
func (UnimplementedGithubridgeServiceServer) GetComments(context.Context, *GetCommentsRequest) (*GetCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComments not implemented")
}
func (UnimplementedGithubridgeServiceServer) GetLabels(context.Context, *GetLabelsRequest) (*GetLabelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLabels not implemented")
}

// UnsafeGithubridgeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GithubridgeServiceServer will
// result in compilation errors.
type UnsafeGithubridgeServiceServer interface {
	mustEmbedUnimplementedGithubridgeServiceServer()
}

func RegisterGithubridgeServiceServer(s grpc.ServiceRegistrar, srv GithubridgeServiceServer) {
	s.RegisterService(&GithubridgeService_ServiceDesc, srv)
}

func _GithubridgeService_AddLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GithubridgeServiceServer).AddLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/githubridge.GithubridgeService/AddLabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GithubridgeServiceServer).AddLabel(ctx, req.(*AddLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GithubridgeService_CreateIssue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIssueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GithubridgeServiceServer).CreateIssue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/githubridge.GithubridgeService/CreateIssue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GithubridgeServiceServer).CreateIssue(ctx, req.(*CreateIssueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GithubridgeService_GetIssue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIssueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GithubridgeServiceServer).GetIssue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/githubridge.GithubridgeService/GetIssue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GithubridgeServiceServer).GetIssue(ctx, req.(*GetIssueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GithubridgeService_GetIssues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIssuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GithubridgeServiceServer).GetIssues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/githubridge.GithubridgeService/GetIssues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GithubridgeServiceServer).GetIssues(ctx, req.(*GetIssuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GithubridgeService_CloseIssue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseIssueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GithubridgeServiceServer).CloseIssue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/githubridge.GithubridgeService/CloseIssue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GithubridgeServiceServer).CloseIssue(ctx, req.(*CloseIssueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GithubridgeService_CommentOnIssue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentOnIssueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GithubridgeServiceServer).CommentOnIssue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/githubridge.GithubridgeService/CommentOnIssue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GithubridgeServiceServer).CommentOnIssue(ctx, req.(*CommentOnIssueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GithubridgeService_GetComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GithubridgeServiceServer).GetComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/githubridge.GithubridgeService/GetComments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GithubridgeServiceServer).GetComments(ctx, req.(*GetCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GithubridgeService_GetLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GithubridgeServiceServer).GetLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/githubridge.GithubridgeService/GetLabels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GithubridgeServiceServer).GetLabels(ctx, req.(*GetLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GithubridgeService_ServiceDesc is the grpc.ServiceDesc for GithubridgeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GithubridgeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "githubridge.GithubridgeService",
	HandlerType: (*GithubridgeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddLabel",
			Handler:    _GithubridgeService_AddLabel_Handler,
		},
		{
			MethodName: "CreateIssue",
			Handler:    _GithubridgeService_CreateIssue_Handler,
		},
		{
			MethodName: "GetIssue",
			Handler:    _GithubridgeService_GetIssue_Handler,
		},
		{
			MethodName: "GetIssues",
			Handler:    _GithubridgeService_GetIssues_Handler,
		},
		{
			MethodName: "CloseIssue",
			Handler:    _GithubridgeService_CloseIssue_Handler,
		},
		{
			MethodName: "CommentOnIssue",
			Handler:    _GithubridgeService_CommentOnIssue_Handler,
		},
		{
			MethodName: "GetComments",
			Handler:    _GithubridgeService_GetComments_Handler,
		},
		{
			MethodName: "GetLabels",
			Handler:    _GithubridgeService_GetLabels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "githubridge.proto",
}
