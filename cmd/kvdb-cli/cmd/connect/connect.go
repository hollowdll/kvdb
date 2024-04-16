package connect

import (
	"github.com/spf13/cobra"
)

var CmdConnect = &cobra.Command{
	Use:   "connect",
	Short: "Manage connection settings",
	Long:  "Manages connection settings used to connect to a kvdb server.",
}

func init() {
	CmdConnect.AddCommand(cmdConnectShow)
	CmdConnect.AddCommand(cmdConnectSet)
}
