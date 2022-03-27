package update

import (
	"cronicle/utils"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a todo entry",
		Long:  "Update a todo entry in your cronicle journal.",
		Run:   run,
	}

	cmd.Flags().IntP("number", "n", 0, "number on ordered list to update")
	cmd.Flags().BoolP("completed", "c", false, "has completed todo")
	cmd.MarkFlagRequired("number")

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	n, _ := cmd.Flags().GetInt("number")
	c, _ := cmd.Flags().GetBool("completed")
	files := utils.GetAllTodos()

	if n == 0 || n > len(files) {
		return
	}

	if c {
		utils.MarkCompleted(files[n-1])
	}
}
