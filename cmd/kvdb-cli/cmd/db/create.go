package db

import (
	"context"
	"fmt"

	"github.com/hollowdll/kvdb/api/v0/dbpb"
	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var (
	cmdDbCreate = &cobra.Command{
		Use:   "create NAME",
		Short: "Create a new database",
		Long:  "Creates a new database.",
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
		Run: func(cmd *cobra.Command, args []string) {
			createDatabase(args[0])
		},
	}
	dbDesc string
)

func init() {
	cmdDbCreate.Flags().StringVarP(&dbDesc, "description", "d", "", "description of the database")
}

func createDatabase(name string) {
	ctx := metadata.NewOutgoingContext(context.Background(), client.GetBaseGrpcMetadata())
	ctx, cancel := context.WithTimeout(ctx, client.CtxTimeout)
	defer cancel()
	res, err := client.GrpcDBClient.CreateDB(ctx, &dbpb.CreateDBRequest{DbName: name, Description: dbDesc})
	client.CheckGrpcError(err)
	fmt.Println(res.DbName)
}
