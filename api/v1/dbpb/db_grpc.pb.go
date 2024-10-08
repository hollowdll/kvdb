// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: api/v1/dbpb/db.proto

package dbpb

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

// DBServiceClient is the client API for DBService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DBServiceClient interface {
	// CreateDB creates a new database.
	CreateDB(ctx context.Context, in *CreateDBRequest, opts ...grpc.CallOption) (*CreateDBResponse, error)
	// GetAllDBs returns the names of all the databases that exist on the server.
	GetAllDBs(ctx context.Context, in *GetAllDBsRequest, opts ...grpc.CallOption) (*GetAllDBsResponse, error)
	// GetDBInfo returns information about a database.
	GetDBInfo(ctx context.Context, in *GetDBInfoRequest, opts ...grpc.CallOption) (*GetDBInfoResponse, error)
	// DeleteDB deletes a database.
	DeleteDB(ctx context.Context, in *DeleteDBRequest, opts ...grpc.CallOption) (*DeleteDBResponse, error)
	// ChangeDB changes the metadata of a database.
	ChangeDB(ctx context.Context, in *ChangeDBRequest, opts ...grpc.CallOption) (*ChangeDBResponse, error)
}

type dBServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDBServiceClient(cc grpc.ClientConnInterface) DBServiceClient {
	return &dBServiceClient{cc}
}

func (c *dBServiceClient) CreateDB(ctx context.Context, in *CreateDBRequest, opts ...grpc.CallOption) (*CreateDBResponse, error) {
	out := new(CreateDBResponse)
	err := c.cc.Invoke(ctx, "/api.v1.dbpb.DBService/CreateDB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBServiceClient) GetAllDBs(ctx context.Context, in *GetAllDBsRequest, opts ...grpc.CallOption) (*GetAllDBsResponse, error) {
	out := new(GetAllDBsResponse)
	err := c.cc.Invoke(ctx, "/api.v1.dbpb.DBService/GetAllDBs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBServiceClient) GetDBInfo(ctx context.Context, in *GetDBInfoRequest, opts ...grpc.CallOption) (*GetDBInfoResponse, error) {
	out := new(GetDBInfoResponse)
	err := c.cc.Invoke(ctx, "/api.v1.dbpb.DBService/GetDBInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBServiceClient) DeleteDB(ctx context.Context, in *DeleteDBRequest, opts ...grpc.CallOption) (*DeleteDBResponse, error) {
	out := new(DeleteDBResponse)
	err := c.cc.Invoke(ctx, "/api.v1.dbpb.DBService/DeleteDB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBServiceClient) ChangeDB(ctx context.Context, in *ChangeDBRequest, opts ...grpc.CallOption) (*ChangeDBResponse, error) {
	out := new(ChangeDBResponse)
	err := c.cc.Invoke(ctx, "/api.v1.dbpb.DBService/ChangeDB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DBServiceServer is the server API for DBService service.
// All implementations must embed UnimplementedDBServiceServer
// for forward compatibility
type DBServiceServer interface {
	// CreateDB creates a new database.
	CreateDB(context.Context, *CreateDBRequest) (*CreateDBResponse, error)
	// GetAllDBs returns the names of all the databases that exist on the server.
	GetAllDBs(context.Context, *GetAllDBsRequest) (*GetAllDBsResponse, error)
	// GetDBInfo returns information about a database.
	GetDBInfo(context.Context, *GetDBInfoRequest) (*GetDBInfoResponse, error)
	// DeleteDB deletes a database.
	DeleteDB(context.Context, *DeleteDBRequest) (*DeleteDBResponse, error)
	// ChangeDB changes the metadata of a database.
	ChangeDB(context.Context, *ChangeDBRequest) (*ChangeDBResponse, error)
	mustEmbedUnimplementedDBServiceServer()
}

// UnimplementedDBServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDBServiceServer struct {
}

func (UnimplementedDBServiceServer) CreateDB(context.Context, *CreateDBRequest) (*CreateDBResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDB not implemented")
}
func (UnimplementedDBServiceServer) GetAllDBs(context.Context, *GetAllDBsRequest) (*GetAllDBsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllDBs not implemented")
}
func (UnimplementedDBServiceServer) GetDBInfo(context.Context, *GetDBInfoRequest) (*GetDBInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDBInfo not implemented")
}
func (UnimplementedDBServiceServer) DeleteDB(context.Context, *DeleteDBRequest) (*DeleteDBResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDB not implemented")
}
func (UnimplementedDBServiceServer) ChangeDB(context.Context, *ChangeDBRequest) (*ChangeDBResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeDB not implemented")
}
func (UnimplementedDBServiceServer) mustEmbedUnimplementedDBServiceServer() {}

// UnsafeDBServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DBServiceServer will
// result in compilation errors.
type UnsafeDBServiceServer interface {
	mustEmbedUnimplementedDBServiceServer()
}

func RegisterDBServiceServer(s grpc.ServiceRegistrar, srv DBServiceServer) {
	s.RegisterService(&DBService_ServiceDesc, srv)
}

func _DBService_CreateDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDBRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServiceServer).CreateDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.dbpb.DBService/CreateDB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServiceServer).CreateDB(ctx, req.(*CreateDBRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBService_GetAllDBs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllDBsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServiceServer).GetAllDBs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.dbpb.DBService/GetAllDBs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServiceServer).GetAllDBs(ctx, req.(*GetAllDBsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBService_GetDBInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDBInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServiceServer).GetDBInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.dbpb.DBService/GetDBInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServiceServer).GetDBInfo(ctx, req.(*GetDBInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBService_DeleteDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDBRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServiceServer).DeleteDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.dbpb.DBService/DeleteDB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServiceServer).DeleteDB(ctx, req.(*DeleteDBRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBService_ChangeDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeDBRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServiceServer).ChangeDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.dbpb.DBService/ChangeDB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServiceServer).ChangeDB(ctx, req.(*ChangeDBRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DBService_ServiceDesc is the grpc.ServiceDesc for DBService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DBService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.dbpb.DBService",
	HandlerType: (*DBServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDB",
			Handler:    _DBService_CreateDB_Handler,
		},
		{
			MethodName: "GetAllDBs",
			Handler:    _DBService_GetAllDBs_Handler,
		},
		{
			MethodName: "GetDBInfo",
			Handler:    _DBService_GetDBInfo_Handler,
		},
		{
			MethodName: "DeleteDB",
			Handler:    _DBService_DeleteDB_Handler,
		},
		{
			MethodName: "ChangeDB",
			Handler:    _DBService_ChangeDB_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/dbpb/db.proto",
}
