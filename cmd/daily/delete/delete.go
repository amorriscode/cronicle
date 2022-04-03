package delete

import (
	"log"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete a daily file",
		Long:  "delete a daily file in your cronicle journal.",
		Run:   run,
	}

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	log.Println("Deleting a log doc...")
}
