package hashmap

import (
	"github.com/spf13/cobra"
)

var (
	dbName     string
	CmdHashMap = &cobra.Command{
		Use:   "hashmap",
		Short: "Manage HashMap keys",
		Long:  "Manage HashMap keys.",
	}
)

func init() {
	CmdHashMap.AddCommand(cmdSetHashMap)
	CmdHashMap.AddCommand(cmdGetHashMapFieldValues)
	CmdHashMap.AddCommand(cmdDeleteHashMapFields)
	CmdHashMap.AddCommand(cmdGetAllHashMapFieldsAndValues)
}
