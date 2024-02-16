package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdGetKeys = &cobra.Command{
	Use:   "getkeys",
	Short: "Get all the keys of a database",
	Long:  "Get all the keys of a database",
	Run: func(cmd *cobra.Command, args []string) {
		getKeys()
	},
}

func init() {
	cmdGetKeys.Flags().StringVarP(&dbName, "db", "d", "", "database to use")
}

func getKeys() {
	md := client.GetBaseGrpcMetadata()
	if len(dbName) > 0 {
		md.Set(common.GrpcMetadataKeyDbName, dbName)
	}
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, cancel := context.WithTimeout(ctx, client.CtxTimeout)
	defer cancel()

	res, err := client.GrpcStorageClient.GetKeys(ctx, &kvdbserver.GetKeysRequest{})
	client.CheckGrpcError(err)

	for i, key := range res.Keys {
		res.Keys[i] = fmt.Sprintf("%d) %s", i+1, key)
	}
	if len(res.Keys) > 0 {
		fmt.Println(strings.Join(res.Keys, "\n"))
	}
}
