// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: api/v1/serverpb/server.proto

package serverpb

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

// ServerServiceClient is the client API for ServerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerServiceClient interface {
	// GetServerInfo returns information about the server.
	GetServerInfo(ctx context.Context, in *GetServerInfoRequest, opts ...grpc.CallOption) (*GetServerInfoResponse, error)
	// GetLogs returns the server logs if the log file is enabled.
	GetLogs(ctx context.Context, in *GetLogsRequest, opts ...grpc.CallOption) (*GetLogsResponse, error)
	// ReloadConfig reloads the server configurations.
	// Not all configurations can be reloaded at runtime.
	ReloadConfig(ctx context.Context, in *ReloadConfigRequest, opts ...grpc.CallOption) (*ReloadConfigResponse, error)
}

type serverServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServerServiceClient(cc grpc.ClientConnInterface) ServerServiceClient {
	return &serverServiceClient{cc}
}

func (c *serverServiceClient) GetServerInfo(ctx context.Context, in *GetServerInfoRequest, opts ...grpc.CallOption) (*GetServerInfoResponse, error) {
	out := new(GetServerInfoResponse)
	err := c.cc.Invoke(ctx, "/api.v1.serverpb.ServerService/GetServerInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverServiceClient) GetLogs(ctx context.Context, in *GetLogsRequest, opts ...grpc.CallOption) (*GetLogsResponse, error) {
	out := new(GetLogsResponse)
	err := c.cc.Invoke(ctx, "/api.v1.serverpb.ServerService/GetLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverServiceClient) ReloadConfig(ctx context.Context, in *ReloadConfigRequest, opts ...grpc.CallOption) (*ReloadConfigResponse, error) {
	out := new(ReloadConfigResponse)
	err := c.cc.Invoke(ctx, "/api.v1.serverpb.ServerService/ReloadConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerServiceServer is the server API for ServerService service.
// All implementations must embed UnimplementedServerServiceServer
// for forward compatibility
type ServerServiceServer interface {
	// GetServerInfo returns information about the server.
	GetServerInfo(context.Context, *GetServerInfoRequest) (*GetServerInfoResponse, error)
	// GetLogs returns the server logs if the log file is enabled.
	GetLogs(context.Context, *GetLogsRequest) (*GetLogsResponse, error)
	// ReloadConfig reloads the server configurations.
	// Not all configurations can be reloaded at runtime.
	ReloadConfig(context.Context, *ReloadConfigRequest) (*ReloadConfigResponse, error)
	mustEmbedUnimplementedServerServiceServer()
}

// UnimplementedServerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServerServiceServer struct {
}

func (UnimplementedServerServiceServer) GetServerInfo(context.Context, *GetServerInfoRequest) (*GetServerInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerInfo not implemented")
}
func (UnimplementedServerServiceServer) GetLogs(context.Context, *GetLogsRequest) (*GetLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogs not implemented")
}
func (UnimplementedServerServiceServer) ReloadConfig(context.Context, *ReloadConfigRequest) (*ReloadConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReloadConfig not implemented")
}
func (UnimplementedServerServiceServer) mustEmbedUnimplementedServerServiceServer() {}

// UnsafeServerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerServiceServer will
// result in compilation errors.
type UnsafeServerServiceServer interface {
	mustEmbedUnimplementedServerServiceServer()
}

func RegisterServerServiceServer(s grpc.ServiceRegistrar, srv ServerServiceServer) {
	s.RegisterService(&ServerService_ServiceDesc, srv)
}

func _ServerService_GetServerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).GetServerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.serverpb.ServerService/GetServerInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).GetServerInfo(ctx, req.(*GetServerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerService_GetLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).GetLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.serverpb.ServerService/GetLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).GetLogs(ctx, req.(*GetLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerService_ReloadConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReloadConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServiceServer).ReloadConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.serverpb.ServerService/ReloadConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServiceServer).ReloadConfig(ctx, req.(*ReloadConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServerService_ServiceDesc is the grpc.ServiceDesc for ServerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.serverpb.ServerService",
	HandlerType: (*ServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServerInfo",
			Handler:    _ServerService_GetServerInfo_Handler,
		},
		{
			MethodName: "GetLogs",
			Handler:    _ServerService_GetLogs_Handler,
		},
		{
			MethodName: "ReloadConfig",
			Handler:    _ServerService_ReloadConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/serverpb/server.proto",
}
