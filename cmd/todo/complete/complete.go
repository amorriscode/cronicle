package complete

import (
	"cronicle/utils/todo"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "complete <ID>",
		Short: "complete a todo entry",
		Long:  "complete a todo entry",
		Run:   run,
	}

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	todo.CompleteTodo()
}
