package main

import (
	"context"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"sync"

	kvdb "github.com/hollowdll/kvdb"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"github.com/hollowdll/kvdb/version"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	kvdbserver.UnimplementedDatabaseServiceServer
	kvdbserver.UnimplementedServerServiceServer
	kvdbserver.UnimplementedStorageServiceServer
	databases map[string]*kvdb.Database
	logger    kvdb.Logger
	mutex     sync.RWMutex
}

func newServer() *server {
	return &server{
		databases: make(map[string]*kvdb.Database),
		logger:    *kvdb.NewLogger(),
	}
}

// getTotalDataSize returns the total amount of stored data on this server in bytes.
func (s *server) getTotalDataSize() uint64 {
	var sum uint64
	for _, db := range s.databases {
		sum += db.GetStoredSizeBytes()
	}

	return sum
}

// getOsInfo returns information about the server's operating system.
func getOsInfo() (string, error) {
	osInfo := runtime.GOOS

	switch osInfo {
	case "linux":
		cmd := exec.Command("uname", "-o", "-s", "-r")
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

// GetServerInfo returns information about the server.
func (s *server) GetServerInfo(ctx context.Context, req *kvdbserver.GetServerInfoRequest) (*kvdbserver.GetServerInfoResponse, error) {
	osInfo, err := getOsInfo()
	if err != nil {
		errMsg := fmt.Sprintf("%s", err)
		return nil, status.Error(codes.Internal, errMsg)
	}

	s.mutex.RLock()
	defer s.mutex.RUnlock()

	info := &kvdbserver.ServerInfo{
		Version:       version.Version,
		GoVersion:     runtime.Version(),
		DbCount:       uint32(len(s.databases)),
		TotalDataSize: s.getTotalDataSize(),
		Os:            osInfo,
		Arch:          runtime.GOARCH,
	}

	return &kvdbserver.GetServerInfoResponse{Info: info}, nil
}
