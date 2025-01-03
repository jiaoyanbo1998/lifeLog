// 版本号

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: api/proto/interactive/v1/interactive.proto

// 包名

package interactivev1

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
	InteractiveService_Like_FullMethodName               = "/interactive.v1.interactiveService/Like"
	InteractiveService_UnLike_FullMethodName             = "/interactive.v1.interactiveService/UnLike"
	InteractiveService_Collect_FullMethodName            = "/interactive.v1.interactiveService/Collect"
	InteractiveService_UnCollect_FullMethodName          = "/interactive.v1.interactiveService/UnCollect"
	InteractiveService_GetInteractiveInfo_FullMethodName = "/interactive.v1.interactiveService/GetInteractiveInfo"
)

// InteractiveServiceClient is the client API for InteractiveService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// interactiveService 互动服务
type InteractiveServiceClient interface {
	Like(ctx context.Context, in *LikeRequest, opts ...grpc.CallOption) (*LikeResponse, error)
	UnLike(ctx context.Context, in *UnLikeRequest, opts ...grpc.CallOption) (*UnLikeResponse, error)
	Collect(ctx context.Context, in *CollectRequest, opts ...grpc.CallOption) (*CollectResponse, error)
	UnCollect(ctx context.Context, in *UnCollectRequest, opts ...grpc.CallOption) (*UnCollectResponse, error)
	GetInteractiveInfo(ctx context.Context, in *GetInteractiveInfoRequest, opts ...grpc.CallOption) (*GetInteractiveInfoResponse, error)
}

type interactiveServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInteractiveServiceClient(cc grpc.ClientConnInterface) InteractiveServiceClient {
	return &interactiveServiceClient{cc}
}

func (c *interactiveServiceClient) Like(ctx context.Context, in *LikeRequest, opts ...grpc.CallOption) (*LikeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LikeResponse)
	err := c.cc.Invoke(ctx, InteractiveService_Like_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactiveServiceClient) UnLike(ctx context.Context, in *UnLikeRequest, opts ...grpc.CallOption) (*UnLikeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UnLikeResponse)
	err := c.cc.Invoke(ctx, InteractiveService_UnLike_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactiveServiceClient) Collect(ctx context.Context, in *CollectRequest, opts ...grpc.CallOption) (*CollectResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CollectResponse)
	err := c.cc.Invoke(ctx, InteractiveService_Collect_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactiveServiceClient) UnCollect(ctx context.Context, in *UnCollectRequest, opts ...grpc.CallOption) (*UnCollectResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UnCollectResponse)
	err := c.cc.Invoke(ctx, InteractiveService_UnCollect_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactiveServiceClient) GetInteractiveInfo(ctx context.Context, in *GetInteractiveInfoRequest, opts ...grpc.CallOption) (*GetInteractiveInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetInteractiveInfoResponse)
	err := c.cc.Invoke(ctx, InteractiveService_GetInteractiveInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InteractiveServiceServer is the server API for InteractiveService service.
// All implementations must embed UnimplementedInteractiveServiceServer
// for forward compatibility.
//
// interactiveService 互动服务
type InteractiveServiceServer interface {
	Like(context.Context, *LikeRequest) (*LikeResponse, error)
	UnLike(context.Context, *UnLikeRequest) (*UnLikeResponse, error)
	Collect(context.Context, *CollectRequest) (*CollectResponse, error)
	UnCollect(context.Context, *UnCollectRequest) (*UnCollectResponse, error)
	GetInteractiveInfo(context.Context, *GetInteractiveInfoRequest) (*GetInteractiveInfoResponse, error)
	mustEmbedUnimplementedInteractiveServiceServer()
}

// UnimplementedInteractiveServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedInteractiveServiceServer struct{}

func (UnimplementedInteractiveServiceServer) Like(context.Context, *LikeRequest) (*LikeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Like not implemented")
}
func (UnimplementedInteractiveServiceServer) UnLike(context.Context, *UnLikeRequest) (*UnLikeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnLike not implemented")
}
func (UnimplementedInteractiveServiceServer) Collect(context.Context, *CollectRequest) (*CollectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Collect not implemented")
}
func (UnimplementedInteractiveServiceServer) UnCollect(context.Context, *UnCollectRequest) (*UnCollectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnCollect not implemented")
}
func (UnimplementedInteractiveServiceServer) GetInteractiveInfo(context.Context, *GetInteractiveInfoRequest) (*GetInteractiveInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInteractiveInfo not implemented")
}
func (UnimplementedInteractiveServiceServer) mustEmbedUnimplementedInteractiveServiceServer() {}
func (UnimplementedInteractiveServiceServer) testEmbeddedByValue()                            {}

// UnsafeInteractiveServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InteractiveServiceServer will
// result in compilation errors.
type UnsafeInteractiveServiceServer interface {
	mustEmbedUnimplementedInteractiveServiceServer()
}

func RegisterInteractiveServiceServer(s grpc.ServiceRegistrar, srv InteractiveServiceServer) {
	// If the following call pancis, it indicates UnimplementedInteractiveServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&InteractiveService_ServiceDesc, srv)
}

func _InteractiveService_Like_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractiveServiceServer).Like(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractiveService_Like_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractiveServiceServer).Like(ctx, req.(*LikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractiveService_UnLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnLikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractiveServiceServer).UnLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractiveService_UnLike_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractiveServiceServer).UnLike(ctx, req.(*UnLikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractiveService_Collect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractiveServiceServer).Collect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractiveService_Collect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractiveServiceServer).Collect(ctx, req.(*CollectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractiveService_UnCollect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnCollectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractiveServiceServer).UnCollect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractiveService_UnCollect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractiveServiceServer).UnCollect(ctx, req.(*UnCollectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractiveService_GetInteractiveInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInteractiveInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractiveServiceServer).GetInteractiveInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractiveService_GetInteractiveInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractiveServiceServer).GetInteractiveInfo(ctx, req.(*GetInteractiveInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InteractiveService_ServiceDesc is the grpc.ServiceDesc for InteractiveService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InteractiveService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "interactive.v1.interactiveService",
	HandlerType: (*InteractiveServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Like",
			Handler:    _InteractiveService_Like_Handler,
		},
		{
			MethodName: "UnLike",
			Handler:    _InteractiveService_UnLike_Handler,
		},
		{
			MethodName: "Collect",
			Handler:    _InteractiveService_Collect_Handler,
		},
		{
			MethodName: "UnCollect",
			Handler:    _InteractiveService_UnCollect_Handler,
		},
		{
			MethodName: "GetInteractiveInfo",
			Handler:    _InteractiveService_GetInteractiveInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/interactive/v1/interactive.proto",
}
