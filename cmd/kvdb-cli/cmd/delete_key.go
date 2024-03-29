package cmd

import (
	"context"
	"fmt"

	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdDeleteKey = &cobra.Command{
	Use:   "delete [key]",
	Short: "Delete a key and its value",
	Long:  "Delete a key and the value it's holding",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		deleteKey(args[0])
	},
}

func init() {
	cmdDeleteKey.Flags().StringVarP(&dbName, "database", "d", "", "database to use")
}

func deleteKey(key string) {
	md := client.GetBaseGrpcMetadata()
	if len(dbName) > 0 {
		md.Set(common.GrpcMetadataKeyDbName, dbName)
	}
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, cancel := context.WithTimeout(ctx, client.CtxTimeout)
	defer cancel()

	response, err := client.GrpcStorageClient.DeleteKey(ctx, &kvdbserver.DeleteKeyRequest{Key: key})
	client.CheckGrpcError(err)

	if response.Ok {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
