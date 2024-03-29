package hashmap

import (
	"context"
	"fmt"

	"github.com/hollowdll/kvdb/cmd/kvdb-cli/client"
	"github.com/hollowdll/kvdb/internal/common"
	"github.com/hollowdll/kvdb/proto/kvdbserver"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdDeleteHashMapFields = &cobra.Command{
	Use:   "delete [key] [field ...]",
	Short: "Remove fields from a HashMap",
	Long:  "Remove fields from a HashMap using a key",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		deleteHashMapFields(args[0], args[1:])
	},
}

func init() {
	cmdDeleteHashMapFields.Flags().StringVarP(&dbName, "database", "d", "", "database to use")
}

func deleteHashMapFields(key string, fields []string) {
	md := client.GetBaseGrpcMetadata()
	if len(dbName) > 0 {
		md.Set(common.GrpcMetadataKeyDbName, dbName)
	}
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, cancel := context.WithTimeout(ctx, client.CtxTimeout)
	defer cancel()

	res, err := client.GrpcStorageClient.DeleteHashMapFields(ctx, &kvdbserver.DeleteHashMapFieldsRequest{Key: key, Fields: fields})
	client.CheckGrpcError(err)

	if res.Ok {
		fmt.Printf("%d\n", res.FieldsRemoved)
	} else {
		fmt.Println(client.ValueNone)
	}
}
