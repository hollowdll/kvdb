package db

import (
	"context"
	"fmt"
	"time"

	"github.com/hollowdll/kvdb/api/v0/dbpb"
	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/hollowdll/kvdb/cmd/kvdb-cli/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc/metadata"
)

var cmdDbInfo = &cobra.Command{
	Use:   "info",
	Short: "Show information about a database",
	Long:  "Shows information about a database.",
	Run: func(cmd *cobra.Command, args []string) {
		showDbInfo()
	},
}

func init() {
	cmdDbInfo.Flags().StringVarP(&dbName, "name", "n", "", "name of the database")
}

func showDbInfo() {
	ctx := metadata.NewOutgoingContext(context.Background(), client.GetBaseGrpcMetadata())
	ctx, cancel := context.WithTimeout(ctx, client.CtxTimeout)
	defer cancel()
	if len(dbName) < 1 {
		dbName = viper.GetString(config.ConfigKeyDatabase)
	}
	resp, err := client.GrpcDBClient.GetDBInfo(ctx, &dbpb.GetDBInfoRequest{DbName: dbName})
	client.CheckGrpcError(err)

	var info string
	info += fmt.Sprintf("name: %s\n", resp.Data.Name)
	info += fmt.Sprintf("description: %s\n", resp.Data.Description)
	info += fmt.Sprintf("created_at: %s00:00\n", resp.Data.CreatedAt.AsTime().Format(time.RFC3339))
	info += fmt.Sprintf("updated_at: %s00:00\n", resp.Data.UpdatedAt.AsTime().Format(time.RFC3339))
	info += fmt.Sprintf("key_count: %d\n", resp.Data.GetKeyCount())
	info += fmt.Sprintf("data_size: %dB", resp.Data.GetDataSize())

	fmt.Println(info)
}
