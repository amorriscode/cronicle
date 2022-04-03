package list

import (
	"cronicle/utils"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list daily files",
		Long:  "list daily files",
		Run:   run,
	}

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	utils.ListFiles("daily")
}
