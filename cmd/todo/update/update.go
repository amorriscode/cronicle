package update

import (
	"cronicle/utils"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update [ID!]",
		Short: "update a todo entry",
		Long:  "update a todo entry",
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

	path := utils.GetPath([]string{"todo", files[n-1].Name()})

	utils.EditFile(path)
}
