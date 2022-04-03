package update

import (
	"cronicle/utils/brag"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update <ID>",
		Short: "update a brag entry",
		Long:  "update a brag entry",
		Run:   run,
	}

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	brag.EditBrag()
}
