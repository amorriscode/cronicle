package update

import (
	"cronicle/utils/entries"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update <ID>",
		Short: "update a daily entry",
		Long:  "update a daily entry",
		Run:   run,
	}

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	entries.EditEntry(args, "daily")
}
