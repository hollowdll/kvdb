package client

import (
	"crypto/x509"
	"fmt"
	"os"

	"github.com/hollowdll/kvdb/api/v0/authpb"
	"github.com/hollowdll/kvdb/api/v0/dbpb"
	"github.com/hollowdll/kvdb/api/v0/echopb"
	"github.com/hollowdll/kvdb/api/v0/kvpb"
	"github.com/hollowdll/kvdb/api/v0/serverpb"
	"github.com/hollowdll/kvdb/cmd/kvdbctl/config"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

const (
	// ValueNone is a special value for values that do not exist.
	ValueNone string = "(None)"
)

var (
	GrpcEchoClient      echopb.EchoServiceClient
	GrpcAuthClient      authpb.AuthServiceClient
	GrpcServerClient    serverpb.ServerServiceClient
	GrpcDBClient        dbpb.DBServiceClient
	GrpcGeneralKVClient kvpb.GeneralKVServiceClient
	GrpcStringKVClient  kvpb.StringKVServiceClient
	GrpcHashMapKVClient kvpb.HashMapKVServiceClient
	grpcClientConn      *grpc.ClientConn
)

// InitClient initializes the client and connections.
func InitClient() {
	var dialOpts []grpc.DialOption
	dialOpts = append(dialOpts, grpc.WithTransportCredentials(getTransportCreds()))
	createEmptyTokenCache()

	address := fmt.Sprintf("%s:%d", viper.GetString("host"), viper.GetUint16("port"))
	conn, err := grpc.NewClient(address, dialOpts...)
	if err != nil {
		cobra.CheckErr(fmt.Sprintf("failed to connect to the server: %s", err))
	}

	GrpcEchoClient = echopb.NewEchoServiceClient(conn)
	GrpcAuthClient = authpb.NewAuthServiceClient(conn)
	GrpcServerClient = serverpb.NewServerServiceClient(conn)
	GrpcDBClient = dbpb.NewDBServiceClient(conn)
	GrpcGeneralKVClient = kvpb.NewGeneralKVServiceClient(conn)
	GrpcStringKVClient = kvpb.NewStringKVServiceClient(conn)
	GrpcHashMapKVClient = kvpb.NewHashMapKVServiceClient(conn)
	grpcClientConn = conn
}

func getTransportCreds() credentials.TransportCredentials {
	if viper.GetBool(config.ConfigKeyTlsEnabled) {
		certBytes, err := os.ReadFile(viper.GetString(config.ConfigKeyTlsCertPath))
		if err != nil {
			cobra.CheckErr(fmt.Sprintf("failed to read TLS certificate: %v", err))
		}
		certPool := x509.NewCertPool()
		if !certPool.AppendCertsFromPEM(certBytes) {
			cobra.CheckErr("failed to parse TLS certificate")
		}

		return credentials.NewClientTLSFromCert(certPool, "")
	} else {
		return insecure.NewCredentials()
	}
}

// CloseConnections closes all connections to the server.
func CloseConnections() {
	if grpcClientConn != nil {
		if err := grpcClientConn.Close(); err != nil {
			cobra.CheckErr(fmt.Sprintf("failed to close connections: %v", err))
		}
	}
}

// ReadPasswordFromEnv reads password from environment variable.
// The returned bool is true if it is present.
func ReadPasswordFromEnv() (string, bool) {
	return os.LookupEnv(config.EnvVarPassword)
}

// GetBaseGrpcMetadata returns base gRPC metadata for all requests.
// It can be overwritten or extended.
func GetBaseGrpcMetadata() metadata.MD {
	md := metadata.Pairs()
	file, err := GetTokenCacheFilePath()
	cobra.CheckErr(err)
	token, err := ReadTokenFromCache(file)
	cobra.CheckErr(err)
	if token != "" {
		md.Set(common.GrpcMetadataKeyAuthToken, token)
	}

	dbName := viper.GetString(config.ConfigKeyDatabase)
	md.Set(common.GrpcMetadataKeyDbName, dbName)

	return md
}
