package update

import (
	"cronicle/utils"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "update a todo entry in vim",
		Long:  "update a todo entry with number on ordered list as arg",
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
	path := utils.GetPath([]string{"todo", files[n-1].Name()})

	utils.EditFile(path)
}
