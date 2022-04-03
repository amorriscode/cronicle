package delete

import (
	"cronicle/utils"
	"fmt"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete a todo entry",
		Long:  "delete a todo entry in your cronicle journal.",
		Run:   run,
	}

	cmd.Flags().IntP("number", "n", 0, "number on ordered list to delete")
	cmd.MarkFlagRequired("number")

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	n, _ := cmd.Flags().GetInt("number")

	files := utils.GetAllFiles("todo")

	if n == 0 || n > len(files) {
		fmt.Printf("Number is not valid")
		return
	}

	utils.DeleteTodo(files[n-1].Name())
	utils.ListTodos()
}
