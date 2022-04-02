package complete

import (
	"cronicle/utils"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "complete",
		Short: "complete a todo entry",
		Long:  "complete a todo entry in your cronicle journal.",
		Run:   run,
	}

	cmd.Flags().IntP("number", "n", 0, "number on ordered list to complete")
	cmd.MarkFlagRequired("number")

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	n, _ := cmd.Flags().GetInt("number")
	files := utils.GetAllFiles("todo")

	if n == 0 || n > len(files) {
		return
	}

	utils.MarkCompleted(files[n-1])
	utils.ListTodo()
}
