package update

import (
	"cronicle/utils/entries"
	"log"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update <ID>",
		Short: "update a todo entry",
		Long:  "update a todo entry",
		Run:   run,
	}

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	log.Println(args)
	entries.EditEntry(args, "todo")
}
