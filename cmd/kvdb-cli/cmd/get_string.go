package cmd

import (
	"context"
	"fmt"

	"github.com/hollowdll/kvdb/api/v0/kvpb"
	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdGetString = &cobra.Command{
	Use:   "get [key]",
	Short: "Get a String value",
	Long:  "Gets the value a String key is holding.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		getString(args[0])
	},
}

func init() {
	cmdGetString.Flags().StringVarP(&dbName, "database", "d", "", "database to use")
}

func getString(key string) {
	md := client.GetBaseGrpcMetadata()
	if len(dbName) > 0 {
		md.Set(common.GrpcMetadataKeyDbName, dbName)
	}
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, cancel := context.WithTimeout(ctx, client.CtxTimeout)
	defer cancel()

	response, err := client.GrpcStringKVClient.GetString(ctx, &kvpb.GetStringRequest{Key: key})
	client.CheckGrpcError(err)

	if response.Ok {
		fmt.Printf("\"%s\"\n", response.Value)
	} else {
		fmt.Println(client.ValueNone)
	}
}
