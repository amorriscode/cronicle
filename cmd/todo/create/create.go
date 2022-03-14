package create

import (
	"cronicle/utils"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd :=
		&cobra.Command{
			Use:   "create",
			Short: "Create a new todo",
			Long:  "Create a new todo in your cronicle journal.",
			Run:   run,
		}

	cmd.Flags().StringP("message", "m", "", "content of your todo")
	cmd.MarkFlagRequired("message")

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	m, _ := cmd.Flags().GetString("message")
	utils.WriteToFile(m)
}
