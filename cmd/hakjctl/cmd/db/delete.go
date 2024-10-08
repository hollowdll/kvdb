package db

import (
	"context"
	"fmt"

	"github.com/hollowdll/hakjdb/api/v1/dbpb"
	"github.com/hollowdll/hakjdb/cmd/hakjctl/client"
	"github.com/hollowdll/hakjdb/cmd/hakjctl/config"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdDbDelete = &cobra.Command{
	Use:   "delete NAME",
	Short: "Delete a database",
	Long:  "Delete the database with the specified name. Returns the name of the deleted database.",
	Example: `# Delete database 'mydb'
hakjctl db delete mydb`,
	Args: cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		deleteDatabase(args[0])
	},
}

func deleteDatabase(name string) {
	if !client.PromptConfirm(fmt.Sprintf("Delete database '%s' and all its data? Yes/No: ", name)) {
		return
	}
	ctx := metadata.NewOutgoingContext(context.Background(), client.GetBaseGrpcMetadata())
	ctx, cancel := context.WithTimeout(ctx, config.GetCmdTimeout())
	defer cancel()

	res, err := client.GrpcDBClient.DeleteDB(ctx, &dbpb.DeleteDBRequest{DbName: name})
	client.CheckGrpcError(err)
	fmt.Println(res.DbName)
}
