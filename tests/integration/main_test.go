package integration

import (
	"crypto/x509"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/hollowdll/kvdb/cmd/kvdbserver/config"
	grpcserver "github.com/hollowdll/kvdb/cmd/kvdbserver/grpc"
	"github.com/hollowdll/kvdb/cmd/kvdbserver/server"
	"github.com/hollowdll/kvdb/internal/testutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

const ctxTimeout = time.Second * 5

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func startTestServer(cfg config.ServerConfig) (*server.KvdbServer, *grpc.Server, int) {
	fmt.Fprint(os.Stderr, "creating test server ...\n")
	s := server.NewKvdbServer(cfg, testutil.DisabledLogger())
	s.CreateDefaultDatabase(cfg.DefaultDB)
	gs := grpcserver.SetupGrpcServer(s)

	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to listen: %v\n", err)
	}
	port := lis.Addr().(*net.TCPAddr).Port
	s.Cfg.PortInUse = uint16(port)
	connLis := server.NewClientConnListener(lis, s, cfg.MaxClientConnections)
	s.ClientConnListener = connLis
	fmt.Fprintf(os.Stderr, "test server listening at %v\n", lis.Addr())

	go func() {
		if err := gs.Serve(connLis); err != nil {
			fmt.Fprintf(os.Stderr, "failed to serve gRPC: %v\n", err)
		}
	}()

	return s, gs, port
}

func defaultConfig() config.ServerConfig {
	return config.DefaultConfig()
}

func tlsConfig() config.ServerConfig {
	tlsCertPath, err := filepath.Abs("../../tls/test-cert/kvdbserver.crt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get TLS certificate path: %v\n", err)
	}
	tlsPrivKeyPath, err := filepath.Abs("../../tls/test-cert/kvdbserver.key")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get TLS private key path: %v\n", err)
	}
	cfg := config.DefaultConfig()
	cfg.TLSEnabled = true
	cfg.TLSCertPath = tlsCertPath
	cfg.TLSPrivKeyPath = tlsPrivKeyPath
	return cfg
}

func getServerAddress(port int) string {
	return fmt.Sprintf("localhost:%d", port)
}

func insecureConnection(address string) (*grpc.ClientConn, error) {
	return grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func secureConnection(address string) (*grpc.ClientConn, error) {
	certBytes, err := os.ReadFile("../../tls/test-cert/kvdbserver.crt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read TLS certificate: %v\n", err)
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(certBytes) {
		fmt.Fprint(os.Stderr, "failed to parse certificate\n")
	}

	creds := credentials.NewClientTLSFromCert(certPool, "")
	return grpc.NewClient(address, grpc.WithTransportCredentials(creds))
}
