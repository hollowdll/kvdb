package db

import (
	"github.com/spf13/cobra"
)

var (
	dbName string
	CmdDb  = &cobra.Command{
		Use:   "db",
		Short: "Manage databases",
		Long:  "Manage databases.",
	}
)

func init() {
	CmdDb.AddCommand(cmdDbCreate)
	CmdDb.AddCommand(cmdDbDelete)
	CmdDb.AddCommand(cmdDbLs)
	CmdDb.AddCommand(cmdDbInfo)
	CmdDb.AddCommand(cmdDbChange)
}
