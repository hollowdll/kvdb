package hashmap

import (
	"context"
	"fmt"
	"strings"

	"github.com/hollowdll/hakjdb/api/v1/kvpb"
	"github.com/hollowdll/hakjdb/cmd/hakjctl/client"
	"github.com/hollowdll/hakjdb/cmd/hakjctl/config"
	"github.com/hollowdll/hakjdb/internal/common"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdGetAllHashMapFieldsAndValues = &cobra.Command{
	Use:   "getall KEY",
	Short: "Get all the fields and values of a HashMap key value",
	Long:  "Get all the fields and values of a HashMap key value. Returns (None) if the key doesn't exist.",
	Example: `# Use the default database
hakjctl hashmap getall key1

# Specify the database to use
hakjctl hashmap getall key1 -d default`,
	Args: cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		getAllHashMapFieldsAndValues(args[0])
	},
}

func init() {
	cmdGetAllHashMapFieldsAndValues.Flags().StringVarP(&dbName, "database", "d", "", client.DBFlagMsg)
}

func getAllHashMapFieldsAndValues(key string) {
	md := client.GetBaseGrpcMetadata()
	if len(dbName) > 0 {
		md.Set(common.GrpcMetadataKeyDbName, dbName)
	}
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, cancel := context.WithTimeout(ctx, config.GetCmdTimeout())
	defer cancel()

	res, err := client.GrpcHashMapKVClient.GetAllHashMapFieldsAndValues(ctx, &kvpb.GetAllHashMapFieldsAndValuesRequest{Key: key})
	client.CheckGrpcError(err)

	if res.Ok {
		if len(res.FieldValueMap) > 0 {
			var builder strings.Builder
			element := 0
			for field, value := range res.FieldValueMap {
				element++
				builder.WriteString(fmt.Sprintf("%d) \"%s\": \"%s\"\n", element, field, value))
			}
			fmt.Print(builder.String())
		}
	} else {
		fmt.Println(client.ValueNone)
	}
}
