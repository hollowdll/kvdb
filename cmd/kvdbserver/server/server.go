package server

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"time"

	kvdb "github.com/hollowdll/kvdb"
	kvdberrors "github.com/hollowdll/kvdb/errors"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"github.com/hollowdll/kvdb/version"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	kvdbserver.UnimplementedDatabaseServiceServer
	kvdbserver.UnimplementedServerServiceServer
	kvdbserver.UnimplementedStorageServiceServer
	startTime       time.Time
	databases       map[string]*kvdb.Database
	CredentialStore InMemoryCredentialStore
	// True if the server is password protected.
	passwordEnabled bool
	logger          kvdb.Logger
	logFilePath     string
	logFileEnabled  bool
	// The maximum number of keys a database can hold.
	maxKeysPerDb uint32
	mutex        sync.RWMutex
}

// ServerOptions contains options that can be passed to the server when creating it.
type ServerOptions struct {
	MaxKeysPerDb uint32
}

// portInUse is the TCP/IP port the server uses.
var portInUse uint16 = common.ServerDefaultPort

func NewServer() *Server {
	return &Server{
		startTime:       time.Now(),
		databases:       make(map[string]*kvdb.Database),
		CredentialStore: *NewInMemoryCredentialStore(),
		passwordEnabled: false,
		logger:          kvdb.NewDefaultLogger(),
		logFilePath:     "",
		logFileEnabled:  false,
		maxKeysPerDb:    common.DbMaxKeyCount,
	}
}

func NewServerWithOptions(options *ServerOptions) *Server {
	return &Server{
		startTime:       time.Now(),
		databases:       make(map[string]*kvdb.Database),
		CredentialStore: *NewInMemoryCredentialStore(),
		passwordEnabled: false,
		logger:          kvdb.NewDefaultLogger(),
		logFilePath:     "",
		logFileEnabled:  false,
		maxKeysPerDb:    options.MaxKeysPerDb,
	}
}

// DisableLogger disables all log outputs from this server.
func (s *Server) DisableLogger() {
	s.logger.Disable()
}

// EnableDebugLogs enables server debug logs.
func (s *Server) EnableDebugLogs() {
	s.logger.EnableDebug()
}

// SetLogFilePath sets the file path to the log file.
func (s *Server) SetLogFilePath(filePath string) {
	s.logFilePath = filePath
}

// EnableLogFile enables logger to write logs to the log file.
func (s *Server) EnableLogFile() {
	err := s.logger.EnableLogFile(s.logFilePath)
	if err != nil {
		s.logger.Fatalf("Failed to enable log file: %v", err)
	}
	s.logFileEnabled = true
}

// CloseLogger closes logger and releases its possible resources.
func (s *Server) CloseLogger() {
	err := s.logger.CloseLogFile()
	if err != nil {
		s.logger.Fatalf("Failed to close log file: %v", err)
	}
}

// EnablePasswordProtection enables server password protection and sets the password.
func (s *Server) EnablePasswordProtection(password string) {
	if err := s.CredentialStore.SetServerPassword([]byte(password)); err != nil {
		s.logger.Fatalf("Failed to set server password: %v", err)
	}
	s.passwordEnabled = true
	s.logger.Infof("Password protection is enabled. Clients need to authenticate using password.")
}

// getTotalDataSize returns the total amount of stored data on this server in bytes.
func (s *Server) getTotalDataSize() uint64 {
	var sum uint64
	for _, db := range s.databases {
		sum += db.GetStoredSizeBytes()
	}

	return sum
}

// CreateDefaultDatabase creates an empty default database.
func (s *Server) CreateDefaultDatabase(name string) {
	if err := kvdb.ValidateDatabaseName(name); err != nil {
		s.logger.Fatalf("Failed to create default database: %v", err)
	}
	db := kvdb.CreateDatabase(name)
	s.databases[db.Name] = db
	s.logger.Infof("Created default database '%s'", db.Name)
}

// DbMaxKeysReached returns true if a database has reached or exceeded the maximum key limit.
func (s *Server) DbMaxKeysReached(db *kvdb.Database) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return db.GetKeyCount() >= s.maxKeysPerDb
}

// getOsInfo returns information about the server's operating system.
func getOsInfo() (string, error) {
	osInfo := runtime.GOOS

	switch osInfo {
	case "linux":
		cmd := exec.Command("uname", "-r", "-m")
		output, err := cmd.Output()
		if err != nil {
			return "", err
		}
		return "Linux " + strings.TrimSpace(string(output)), nil
	case "windows":
		cmd := exec.Command("cmd", "/c", "ver")
		output, err := cmd.Output()
		if err != nil {
			return "", err
		}
		return strings.TrimSpace(string(output)), nil
	default:
		return osInfo, nil
	}
}

// GetServerInfo is the implementation of RPC GetServerInfo.
func (s *Server) GetServerInfo(ctx context.Context, req *kvdbserver.GetServerInfoRequest) (res *kvdbserver.GetServerInfoResponse, err error) {
	logPrefix := "GetServerInfo"
	s.logger.Debugf("%s: attempt to get server info", logPrefix)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: failed to get server info: %v", logPrefix, err)
		} else {
			s.logger.Debugf("%s: get server info success", logPrefix)
		}
	}()

	osInfo, err := getOsInfo()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	s.mutex.RLock()
	defer s.mutex.RUnlock()

	info := &kvdbserver.ServerInfo{
		KvdbVersion:   version.Version,
		GoVersion:     runtime.Version(),
		DbCount:       uint32(len(s.databases)),
		TotalDataSize: s.getTotalDataSize(),
		Os:            osInfo,
		Arch:          runtime.GOARCH,
		ProcessId:     uint32(os.Getpid()),
		UptimeSeconds: uint64(time.Since(s.startTime).Seconds()),
		TcpPort:       uint32(portInUse),
	}

	return &kvdbserver.GetServerInfoResponse{Data: info}, nil
}

// GetLogs is the implementation of RPC GetLogs.
func (s *Server) GetLogs(ctx context.Context, req *kvdbserver.GetLogsRequest) (res *kvdbserver.GetLogsResponse, err error) {
	logPrefix := "GetLogs"
	s.logger.Debugf("%s: attempt to get server logs", logPrefix)
	defer func() {
		if err != nil {
			s.logger.Errorf("%s: failed to get server logs: %v", logPrefix, err)
		} else {
			s.logger.Debugf("%s: get server logs success", logPrefix)
		}
	}()

	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if !s.logFileEnabled {
		return nil, status.Errorf(codes.FailedPrecondition, "%s: enable server log file to get logs", kvdberrors.ErrLogFileNotEnabled.Error())
	}
	s.logger.Debugf("%s: log file is enabled", logPrefix)

	lines, err := common.ReadFileLines(s.logFilePath)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &kvdbserver.GetLogsResponse{Logs: lines, LogfileEnabled: true}, nil
}

// initServer initializes the server.
// Returns the initialized Server and grpc.Server.
func initServer() (*Server, *grpc.Server) {
	server := NewServer()
	initConfig(server)
	server.logger.ClearFlags()
	server.logger.Infof("Starting kvdb v%s server ...", version.Version)

	if viper.GetBool(ConfigKeyLogFileEnabled) {
		server.EnableLogFile()
		server.logger.Infof("Log file is enabled. Logs will be written to the log file. The file is located at %s", server.logFilePath)
	}

	if viper.GetBool(ConfigKeyDebugEnabled) {
		server.EnableDebugLogs()
		server.logger.Info("Debug mode is enabled. Debug messages will be logged.")
	}

	password, present := os.LookupEnv(EnvVarPassword)
	if present {
		server.EnablePasswordProtection(password)
	} else {
		server.logger.Warningf("Password protection is disabled.")
	}

	server.CreateDefaultDatabase(viper.GetString(ConfigKeyDefaultDatabase))

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(server.authInterceptor))
	kvdbserver.RegisterDatabaseServiceServer(grpcServer, server)
	kvdbserver.RegisterServerServiceServer(grpcServer, server)
	kvdbserver.RegisterStorageServiceServer(grpcServer, server)

	return server, grpcServer
}

// StartServer initializes and starts the server.
func StartServer() {
	server, grpcServer := initServer()
	defer server.CloseLogger()

	portInUse = viper.GetUint16(ConfigKeyPort)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", portInUse))
	if err != nil {
		server.logger.Fatalf("Failed to listen: %v", err)
	}
	server.logger.Infof("Server listening at %v", listener.Addr())

	if err := grpcServer.Serve(listener); err != nil {
		server.logger.Fatalf("Failed to serve gRPC: %v", err)
	}
}
