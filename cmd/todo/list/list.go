package list

import (
	"cronicle/ui/constants"
	"cronicle/utils"
	"fmt"
	"io/ioutil"
	"log"

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
	path := utils.GetPath([]string{"todo"})

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(constants.ERROR_LIST_FILE, err)
	}

	for i, f := range files {
		todo := utils.GetTodoFromFile(f)
		fmt.Printf("%v. %s\n", i+1, todo)
	}
}
