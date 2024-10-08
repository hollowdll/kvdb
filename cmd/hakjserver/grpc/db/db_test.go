package db

import (
	"context"
	"strings"
	"testing"

	"github.com/hollowdll/hakjdb"
	"github.com/hollowdll/hakjdb/api/v1/dbpb"
	"github.com/hollowdll/hakjdb/cmd/hakjserver/config"
	"github.com/hollowdll/hakjdb/cmd/hakjserver/server"
	"github.com/hollowdll/hakjdb/internal/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCreateDB(t *testing.T) {
	cfg := config.DefaultConfig()

	t.Run("DatabaseNotExists", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewDBServiceServer(s)
		dbName := "test"

		req := &dbpb.CreateDBRequest{DbName: dbName}
		resp, err := gs.CreateDB(context.Background(), req)

		require.NoErrorf(t, err, "expected no error; error = %s", err)
		require.NotNil(t, resp)
		assert.Equalf(t, dbName, resp.DbName, "expected DbName = %s; got = %s", dbName, resp.DbName)
	})

	t.Run("DatabaseAlreadyExists", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewDBServiceServer(s)
		dbName := "test"

		req := &dbpb.CreateDBRequest{DbName: dbName}
		_, err := gs.CreateDB(context.Background(), req)
		require.NoErrorf(t, err, "expected no error; error = %s", err)

		resp, err := gs.CreateDB(context.Background(), req)
		require.Error(t, err)
		require.Nil(t, resp)

		st, ok := status.FromError(err)
		require.NotNil(t, st)
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.AlreadyExists, st.Code(), "expected status = %s; got = %s", codes.AlreadyExists, st.Code())
	})

	t.Run("InvalidDatabaseName", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewDBServiceServer(s)
		dbName := "   "

		req := &dbpb.CreateDBRequest{DbName: dbName}
		resp, err := gs.CreateDB(context.Background(), req)
		require.Error(t, err)
		require.Nil(t, resp)

		st, ok := status.FromError(err)
		require.NotNil(t, st)
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.InvalidArgument, st.Code(), "expected status = %s; got = %s", codes.InvalidArgument, st.Code())
	})

	t.Run("InvalidDatabaseDescription", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewDBServiceServer(s)
		name := "db123"
		desc := strings.Repeat("a", 256)

		req := &dbpb.CreateDBRequest{DbName: name, Description: desc}
		resp, err := gs.CreateDB(context.Background(), req)
		require.Error(t, err)
		require.Nil(t, resp)

		st, ok := status.FromError(err)
		require.NotNil(t, st)
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.InvalidArgument, st.Code(), "expected status = %s; got = %s", codes.InvalidArgument, st.Code())
	})
}

func TestGetAllDBs(t *testing.T) {
	cfg := config.DefaultConfig()

	t.Run("NoDatabases", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewDBServiceServer(s)
		expected := 0
		req := &dbpb.GetAllDBsRequest{}
		resp, err := gs.GetAllDBs(context.Background(), req)

		require.NoErrorf(t, err, "expected no error; error = %s", err)
		require.NotNil(t, resp)
		assert.Equalf(t, expected, len(resp.DbNames), "expected databases = %d; got = %d", expected, len(resp.DbNames))
	})

	t.Run("MultipleDatabases", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewDBServiceServer(s)

		dbs := []string{"db0", "db1", "db2"}
		for _, db := range dbs {
			req := &dbpb.CreateDBRequest{DbName: db}
			_, err := gs.CreateDB(context.Background(), req)
			require.NoErrorf(t, err, "expected no error; error = %s", err)
		}

		req := &dbpb.GetAllDBsRequest{}
		resp, err := gs.GetAllDBs(context.Background(), req)
		require.NoErrorf(t, err, "expected no error; error = %s", err)
		require.NotNil(t, resp)
		assert.Equalf(t, len(dbs), len(resp.DbNames), "expected databases = %d; got = %d", len(dbs), len(resp.DbNames))

		for _, db := range resp.DbNames {
			assert.Equalf(t, true, common.StringInSlice(db, dbs), "expected db name %s to be in %v", db, dbs)
		}
	})
}

func TestGetDBInfo(t *testing.T) {
	cfg := config.DefaultConfig()

	t.Run("DatabaseNotFound", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewDBServiceServer(s)
		dbName := "db0"

		req := &dbpb.GetDBInfoRequest{DbName: dbName}
		resp, err := gs.GetDBInfo(context.Background(), req)
		require.Error(t, err)
		require.Nil(t, resp)

		st, ok := status.FromError(err)
		require.NotNil(t, st)
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.NotFound, st.Code(), "expected status = %s; got = %s", codes.NotFound, st.Code())
	})

	t.Run("DatabaseExists", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewDBServiceServer(s)
		name := "db0"
		desc := "Test database."

		reqCreate := &dbpb.CreateDBRequest{DbName: name, Description: desc}
		_, err := gs.CreateDB(context.Background(), reqCreate)
		require.NoErrorf(t, err, "expected no error; error = %s", err)

		reqGet := &dbpb.GetDBInfoRequest{DbName: name}
		resp, err := gs.GetDBInfo(context.Background(), reqGet)
		expectedKeyCount := uint32(0)
		expectedDataSize := uint64(0)

		require.NoErrorf(t, err, "expected no error; error = %s", err)
		require.NotNil(t, resp)
		assert.Equalf(t, name, resp.Data.Name, "expected db name = %s; got = %s", name, resp.Data.Name)
		assert.Equalf(t, desc, resp.Data.Description, "expected description = %s; got = %s", desc, resp.Data.Description)
		assert.Equalf(t, expectedKeyCount, resp.Data.KeyCount, "expected keys = %d; got = %d", expectedKeyCount, resp.Data.KeyCount)
		assert.Equalf(t, expectedDataSize, resp.Data.DataSize, "expected data size = %d; got = %d", expectedDataSize, resp.Data.DataSize)
	})
}

func TestDeleteDB(t *testing.T) {
	cfg := config.DefaultConfig()

	t.Run("DatabaseExists", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewDBServiceServer(s)
		dbName := "default"
		s.CreateDefaultDatabase(dbName)

		req := &dbpb.DeleteDBRequest{DbName: dbName}
		resp, err := gs.DeleteDB(context.Background(), req)

		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, resp)
		assert.Equalf(t, dbName, resp.DbName, "expected db name = &s; got = %s", dbName, resp.DbName)
	})

	t.Run("DatabaseNotFound", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewDBServiceServer(s)
		dbName := "default"

		req := &dbpb.DeleteDBRequest{DbName: dbName}
		resp, err := gs.DeleteDB(context.Background(), req)
		require.Error(t, err)
		require.Nil(t, resp)

		st, ok := status.FromError(err)
		require.NotNil(t, st)
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.NotFound, st.Code(), "expected status = %s; got = %s", codes.NotFound, st.Code())
	})
}

func TestChangeDB(t *testing.T) {
	cfg := config.DefaultConfig()

	t.Run("ChangeNameAndDesc", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewDBServiceServer(s)
		dbName := "default"
		s.CreateDefaultDatabase(dbName)

		newName := "new-name"
		newDesc := "new-desc"
		req := &dbpb.ChangeDBRequest{
			DbName:            dbName,
			NewName:           newName,
			ChangeName:        true,
			NewDescription:    newDesc,
			ChangeDescription: true,
		}
		resp, err := gs.ChangeDB(context.Background(), req)

		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, resp)
		assert.Equalf(t, newName, resp.DbName, "expected db name = &s; got = %s", newName, resp.DbName)
	})

	t.Run("ChangeOnlyDesc", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewDBServiceServer(s)
		dbName := "default"
		s.CreateDefaultDatabase(dbName)

		newDesc := "new-desc"
		req := &dbpb.ChangeDBRequest{
			DbName:            dbName,
			NewDescription:    newDesc,
			ChangeDescription: true,
		}
		resp, err := gs.ChangeDB(context.Background(), req)

		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, resp)
		assert.Equalf(t, dbName, resp.DbName, "expected db name = &s; got = %s", dbName, resp.DbName)
	})

	t.Run("DatabaseNotFound", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewDBServiceServer(s)
		dbName := "test"

		req := &dbpb.ChangeDBRequest{DbName: dbName}
		resp, err := gs.ChangeDB(context.Background(), req)
		require.Error(t, err)
		require.Nil(t, resp)

		st, ok := status.FromError(err)
		require.NotNil(t, st)
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.NotFound, st.Code(), "expected status = %s; got = %s", codes.NotFound, st.Code())
	})

	t.Run("InvalidName", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewDBServiceServer(s)
		dbName := "default"
		s.CreateDefaultDatabase(dbName)

		req := &dbpb.ChangeDBRequest{
			DbName:     dbName,
			NewName:    "???",
			ChangeName: true,
		}
		resp, err := gs.ChangeDB(context.Background(), req)
		require.Error(t, err)
		require.Nil(t, resp)

		st, ok := status.FromError(err)
		require.NotNil(t, st)
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.InvalidArgument, st.Code(), "expected status = %s; got = %s", codes.InvalidArgument, st.Code())
	})

	t.Run("InvalidDescription", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewDBServiceServer(s)
		dbName := "default"
		s.CreateDefaultDatabase(dbName)
		desc := strings.Repeat("a", 256)

		req := &dbpb.ChangeDBRequest{
			DbName:            dbName,
			NewDescription:    desc,
			ChangeDescription: true,
		}
		resp, err := gs.ChangeDB(context.Background(), req)
		require.Error(t, err)
		require.Nil(t, resp)

		st, ok := status.FromError(err)
		require.NotNil(t, st)
		require.Equal(t, true, ok, "expected ok")
		assert.Equal(t, codes.InvalidArgument, st.Code(), "expected status = %s; got = %s", codes.InvalidArgument, st.Code())
	})

	t.Run("ChangeNameDbExists", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewDBServiceServer(s)
		dbName := "default"
		s.CreateDefaultDatabase(dbName)

		newName := "new-name"
		req := &dbpb.ChangeDBRequest{
			DbName:     dbName,
			NewName:    newName,
			ChangeName: true,
		}
		resp, err := gs.ChangeDB(context.Background(), req)
		require.NoErrorf(t, err, "expected no error; error = %v", err)
		require.NotNil(t, resp)
		assert.Equalf(t, newName, resp.DbName, "expected db name = &s; got = %s", newName, resp.DbName)

		reqGet := &dbpb.GetDBInfoRequest{DbName: newName}
		respGet, err := gs.GetDBInfo(context.Background(), reqGet)
		require.NoErrorf(t, err, "expected no error; error = %s", err)
		require.NotNil(t, respGet)
		assert.Equalf(t, newName, respGet.Data.Name, "expected db name = %s; got = %s", newName, respGet.Data.Name)

		reqGetAll := &dbpb.GetAllDBsRequest{}
		respGetAll, err := gs.GetAllDBs(context.Background(), reqGetAll)
		expectedDbCount := 1
		require.NoErrorf(t, err, "expected no error; error = %s", err)
		require.NotNil(t, respGetAll)
		assert.Equalf(t, expectedDbCount, len(respGetAll.DbNames), "expected databases = %d; got = %d", expectedDbCount, len(respGetAll.DbNames))
	})
}

func TestDefaultDatabase(t *testing.T) {
	cfg := config.DefaultConfig()

	t.Run("GetDatabaseInfo", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewDBServiceServer(s)
		dbName := "default"
		s.CreateDefaultDatabase(dbName)

		req := &dbpb.GetDBInfoRequest{DbName: dbName}
		resp, err := gs.GetDBInfo(context.Background(), req)
		require.NoErrorf(t, err, "expected no error; error = %s", err)
		require.NotNil(t, resp)
		assert.Equalf(t, dbName, resp.Data.Name, "expected db name = %s; got = %s", dbName, resp.Data.Name)
	})

	t.Run("GetAllDatabases", func(t *testing.T) {
		s := server.NewHakjServer(cfg, hakjdb.DisabledLogger())
		gs := NewDBServiceServer(s)
		dbName := "default"
		s.CreateDefaultDatabase(dbName)

		req := &dbpb.GetAllDBsRequest{}
		resp, err := gs.GetAllDBs(context.Background(), req)
		expectedDbCount := 1

		require.NoErrorf(t, err, "expected no error; error = %s", err)
		require.NotNil(t, resp)
		assert.Equalf(t, expectedDbCount, len(resp.DbNames), "expected databases = %d; got = %d", expectedDbCount, len(resp.DbNames))
	})
}
