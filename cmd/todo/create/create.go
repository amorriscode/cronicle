package create

import (
	"log"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd :=
		&cobra.Command{
			Use:   "create",
			Short: "Create a new todo",
			Long:  "Create a new todo in your cronicle journal.",
			Run:   run,
		}

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	log.Println("Creating a new todo...")
}
