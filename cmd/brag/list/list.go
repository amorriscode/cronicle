package list

import (
	"cronicle/utils"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list brag files",
		Long:  "list brag files",
		Run:   run,
	}

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	utils.ListFiles("brag")
}
