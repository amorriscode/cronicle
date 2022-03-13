package create

import (
	"log"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new brag entry",
		Long:  "Create a new brag entry in your cronicle journal.",
		Run:   run,
	}
	return cmd
}

func run(cmd *cobra.Command, args []string) {
	log.Println("Creating a new brag entry...")
}
