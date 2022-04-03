package delete

import (
	"cronicle/utils"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete a todo entry",
		Long:  "delete a todo entry in your cronicle journal.",
		Run:   run,
	}

	return cmd
}

func run(cmd *cobra.Command, args []string) {

	files := utils.GetAllFiles("todo")

	n, err := strconv.Atoi(args[0])
	if err != nil || n == 0 || n > len(files) {
		fmt.Printf("Number is not valid")
		return
	}

	utils.DeleteFile(files[n-1].Name(), "todo")
	utils.ListTodos()
}
