package list

import (
	"cronicle/utils"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list todo entries",
		Long:  "list uncompleted todo entries with id",
		Run:   run,
	}

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	utils.ListTodos()
}
