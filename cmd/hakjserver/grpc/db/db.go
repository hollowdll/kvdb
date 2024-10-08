package db

import (
	"context"

	"github.com/hollowdll/hakjdb/api/v1/dbpb"
	grpcerrors "github.com/hollowdll/hakjdb/cmd/hakjserver/grpc/errors"
	"github.com/hollowdll/hakjdb/cmd/hakjserver/server"
)

type DBServiceServer struct {
	srv server.DBService
	dbpb.UnimplementedDBServiceServer
}

func NewDBServiceServer(s *server.HakjServer) dbpb.DBServiceServer {
	return &DBServiceServer{srv: s}
}

// CreateDB is the implementation of RPC CreateDB.
func (s *DBServiceServer) CreateDB(ctx context.Context, req *dbpb.CreateDBRequest) (res *dbpb.CreateDBResponse, err error) {
	res, err = s.srv.CreateDB(ctx, req)
	if err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}

	return res, nil
}

// DeleteDB is the implementation of RPC DeleteDB.
func (s *DBServiceServer) DeleteDB(ctx context.Context, req *dbpb.DeleteDBRequest) (res *dbpb.DeleteDBResponse, err error) {
	res, err = s.srv.DeleteDB(ctx, req)
	if err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}

	return res, nil
}

// GetAllDBs is the implementation of RPC GetAllDBs.
func (s *DBServiceServer) GetAllDBs(ctx context.Context, req *dbpb.GetAllDBsRequest) (res *dbpb.GetAllDBsResponse, err error) {
	res, err = s.srv.GetAllDBs(ctx, req)
	if err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}

	return res, nil
}

// GetDBInfo is the implementation of RPC GetDBInfo.
func (s *DBServiceServer) GetDBInfo(ctx context.Context, req *dbpb.GetDBInfoRequest) (res *dbpb.GetDBInfoResponse, err error) {
	res, err = s.srv.GetDBInfo(ctx, req)
	if err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}

	return res, nil
}

// ChangeDB is the implementation of RPC ChangeDB.
func (s *DBServiceServer) ChangeDB(ctx context.Context, req *dbpb.ChangeDBRequest) (res *dbpb.ChangeDBResponse, err error) {
	res, err = s.srv.ChangeDB(ctx, req)
	if err != nil {
		return nil, grpcerrors.ToGrpcError(err)
	}

	return res, nil
}
