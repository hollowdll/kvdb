package db

import (
	"context"
	"fmt"

	"github.com/hollowdll/kvdb/api/v0/dbpb"
	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdDbCreate = &cobra.Command{
	Use:   "create",
	Short: "Create a new database",
	Long:  "Creates a new database.",
	Run: func(cmd *cobra.Command, args []string) {
		createDatabase()
	},
}

func init() {
	cmdDbCreate.Flags().StringVarP(&dbName, "name", "n", "", "name of the database (required)")
	cmdDbCreate.MarkFlagRequired("name")
}

func createDatabase() {
	ctx := metadata.NewOutgoingContext(context.Background(), client.GetBaseGrpcMetadata())
	ctx, cancel := context.WithTimeout(ctx, client.CtxTimeout)
	defer cancel()
	res, err := client.GrpcDBClient.CreateDB(ctx, &dbpb.CreateDBRequest{DbName: dbName})
	client.CheckGrpcError(err)

	fmt.Println(res.DbName)
}
