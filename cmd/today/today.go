package today

import (
	"cronicle/cmd/today/create"
	"cronicle/cmd/today/delete"
	"cronicle/cmd/today/update"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "today",
		Short: "Manage your log entries",
		Long:  "Manage the log entries in your cronicle journal.",
	}

	cmd.AddCommand(create.New())
	cmd.AddCommand(update.New())
	cmd.AddCommand(delete.New())

	return cmd
}
