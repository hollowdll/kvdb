package db

import (
	"context"
	"fmt"

	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"github.com/spf13/cobra"
)

var cmdDbCreate = &cobra.Command{
	Use:   "create [flags]",
	Short: "Create a new database",
	Long:  "Create a new database",
	Run: func(cmd *cobra.Command, args []string) {
		createDatabase()
	},
}

func init() {
	cmdDbCreate.Flags().StringVarP(&dbName, "name", "n", "", "name of the database (required)")
	cmdDbCreate.MarkFlagRequired("name")
}

func createDatabase() {
	ctx, cancel := context.WithTimeout(context.Background(), client.ClientCtxTimeout)
	defer cancel()
	_, err := client.GrpcDatabaseClient.CreateDatabase(ctx, &kvdbserver.CreateDatabaseRequest{DbName: dbName})
	client.CheckGrpcError(err)

	fmt.Println("OK")
}
