package delete

import (
	"log"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a brag entry",
		Long:  "Delete a brag entry in your cronicle journal.",
		Run:   run,
	}

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	log.Println("Deleting a brag doc...")
}
