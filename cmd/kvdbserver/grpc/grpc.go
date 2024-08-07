package grpc

import (
	"github.com/hollowdll/kvdb/api/v0/dbpb"
	"github.com/hollowdll/kvdb/api/v0/echopb"
	"github.com/hollowdll/kvdb/api/v0/kvpb"
	"github.com/hollowdll/kvdb/api/v0/serverpb"
	dbrpc "github.com/hollowdll/kvdb/cmd/kvdbserver/grpc/db"
	echorpc "github.com/hollowdll/kvdb/cmd/kvdbserver/grpc/echo"
	kvrpc "github.com/hollowdll/kvdb/cmd/kvdbserver/grpc/kv"
	serverrpc "github.com/hollowdll/kvdb/cmd/kvdbserver/grpc/server"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func SetupGrpcServer(s *server.KvdbServer) *grpc.Server {
	logger := s.Logger()
	logger.Infof("Setting up gRPC server ...")
	var opts []grpc.ServerOption
	chainUnaryInterceptors := []grpc.UnaryServerInterceptor{
		newLogUnaryInterceptor(s),
		newAuthUnaryInterceptor(s),
		newHeaderUnaryInterceptor(s),
	}
	opts = append(opts, grpc.ChainUnaryInterceptor(chainUnaryInterceptors...))

	if !s.Cfg.TLSEnabled {
		logger.Warning("TLS is disabled. Connections will not be encrypted")
	} else {
		logger.Info("Attempting to enable TLS ...")
		cert := s.GetTLSCert()
		opts = append(opts, grpc.Creds(credentials.NewServerTLSFromCert(&cert)))
		logger.Info("TLS is enabled. Connections will be encrypted")
	}

	grpcServer := grpc.NewServer(opts...)
	echopb.RegisterEchoServiceServer(grpcServer, echorpc.NewEchoServiceServer())
	serverpb.RegisterServerServiceServer(grpcServer, serverrpc.NewServerServiceServer(s))
	dbpb.RegisterDBServiceServer(grpcServer, dbrpc.NewDBServiceServer(s))
	kvpb.RegisterGeneralKVServiceServer(grpcServer, kvrpc.NewGeneralKVServiceServer(s))
	kvpb.RegisterStringKVServiceServer(grpcServer, kvrpc.NewStringKVServiceServer(s))
	kvpb.RegisterHashMapKVServiceServer(grpcServer, kvrpc.NewHashMapKVServiceServer(s))

	// enable gRPC server reflection in debug mode
	if s.Cfg.DebugEnabled {
		logger.Info("Debug mode detected: enabling gRPC server reflection ...")
		reflection.Register(grpcServer)
		logger.Info("gRPC server reflection enabled")
	}

	return grpcServer
}

func ServeGrpcServer(s *server.KvdbServer, grpcServer *grpc.Server) {
	logger := s.Logger()
	if err := grpcServer.Serve(s.ClientConnListener); err != nil {
		logger.Errorf("Failed to accept incoming connection: %v", err)
	}
}
