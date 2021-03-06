package delete

import (
	"cronicle/utils"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete [ID!]",
		Short: "delete a brag entry",
		Long:  "delete a brag entry",
		Run:   run,
	}

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	files := utils.GetAllFiles("brag")

	n, err := strconv.Atoi(args[0])
	if err != nil || n == 0 || n > len(files) {
		fmt.Printf("Invalid argument")
		return
	}

	utils.DeleteFile(files[n-1].Name(), "brag")
	utils.ListFiles("brag")
}
