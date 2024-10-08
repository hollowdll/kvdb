package cmd

import (
	"context"
	"fmt"

	"github.com/hollowdll/hakjdb/api/v1/kvpb"
	"github.com/hollowdll/hakjdb/cmd/hakjctl/client"
	"github.com/hollowdll/hakjdb/cmd/hakjctl/config"
	"github.com/hollowdll/hakjdb/internal/common"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdGetKeyType = &cobra.Command{
	Use:   "keytype KEY",
	Short: "Get the data type of a key",
	Long:  "Get the data type of a key. Returns (None) if the key doesn't exist.",
	Example: `# Use the default database
hakjctl keytype mykey

# Specify the database to use
hakjctl keytype mykey -d default`,
	Args: cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		getKeyType(args[0])
	},
}

func init() {
	cmdGetKeyType.Flags().StringVarP(&dbName, "database", "d", "", client.DBFlagMsg)
}

func getKeyType(key string) {
	md := client.GetBaseGrpcMetadata()
	if len(dbName) > 0 {
		md.Set(common.GrpcMetadataKeyDbName, dbName)
	}
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, cancel := context.WithTimeout(ctx, config.GetCmdTimeout())
	defer cancel()

	response, err := client.GrpcGeneralKVClient.GetKeyType(ctx, &kvpb.GetKeyTypeRequest{Key: key})
	client.CheckGrpcError(err)

	if response.Ok {
		fmt.Printf("\"%s\"\n", response.KeyType)
	} else {
		fmt.Println(client.ValueNone)
	}
}
