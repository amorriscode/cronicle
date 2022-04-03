package complete

import (
	"cronicle/utils"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "complete",
		Short: "complete a todo entry with number on ordered list as arg",
		Long:  "complete a todo entry in your cronicle journal.",
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

	utils.MarkCompleted(files[n-1])
	utils.ListTodos()
}
