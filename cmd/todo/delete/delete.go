package delete

import (
	"cronicle/utils"
	"cronicle/utils/todo"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete [ID!]",
		Short: "delete a todo entry",
		Long:  "delete a todo entry",
		Run:   run,
	}

	return cmd
}

func run(cmd *cobra.Command, args []string) {

	files := utils.GetAllFiles("todo")

	n, err := strconv.Atoi(args[0])
	if err != nil || n == 0 || n > len(files) {
		fmt.Printf("Invalid argument")
		return
	}

	utils.DeleteFile(files[n-1].Name(), "todo")
	todo.ListTodos()
}
