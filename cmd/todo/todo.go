package todo

import (
	"cronicle/cmd/todo/create"
	"cronicle/cmd/todo/delete"
	"cronicle/cmd/todo/update"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "todo",
		Short: "Manage your todos",
		Long:  "Manage the todos in your cronicle journal.",
	}
	cmd.AddCommand(create.New())
	cmd.AddCommand(update.New())
	cmd.AddCommand(delete.New())

	return cmd
}
