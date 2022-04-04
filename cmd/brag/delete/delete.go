package delete

import (
	"cronicle/utils/entries"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <ID>",
		Short: "delete a brag entry",
		Long:  "delete a brag entry",
		Run:   run,
	}

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	entries.DeleteEntry(args, "brag")
}
