package client

import (
	"fmt"
	"os"
	"time"

	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	// ClientCtxTimeout specifies how long to wait until operation terminates.
	ClientCtxTimeout = time.Second * 10
	// ValueNone is a special value for values that do not exist.
	ValueNone string = "(None)"
)

var (
	GrpcDatabaseClient   kvdbserver.DatabaseServiceClient
	GrpcStorageClient    kvdbserver.StorageServiceClient
	GrpcServerClient     kvdbserver.ServerServiceClient
	grpcClientConnection *grpc.ClientConn
)

// InitClient initializes the client and connections.
func InitClient() {
	address := fmt.Sprintf("%s:%d", viper.GetString("host"), viper.GetUint16("port"))
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: failed to connect to the server:", err)
		os.Exit(1)
	}

	GrpcDatabaseClient = kvdbserver.NewDatabaseServiceClient(conn)
	GrpcStorageClient = kvdbserver.NewStorageServiceClient(conn)
	GrpcServerClient = kvdbserver.NewServerServiceClient(conn)
	grpcClientConnection = conn
}

// CloseConnections closes all connections to the server.
func CloseConnections() {
	if grpcClientConnection != nil {
		grpcClientConnection.Close()
	}
}
