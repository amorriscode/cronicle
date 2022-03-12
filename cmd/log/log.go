package log

import (
	"cronicle/cmd/log/create"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "log",
		Short: "Manage your log entries",
		Long:  "Manage the log entries in your cronicle journal.",
	}

	cmd.AddCommand(create.New())

	return cmd
}
