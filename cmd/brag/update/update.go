package update

import (
	"log"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a brag entry",
		Long:  "Update a brag entry in your cronicle journal.",
		Run:   run,
	}

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	log.Println("Updating a brag doc...")
}
