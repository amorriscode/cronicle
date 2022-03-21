package delete

import (
	"log"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a todo entry",
		Long:  "Delete a todo entry in your cronicle journal.",
		Run:   run,
	}

	cmd.Flags().StringP("number", "n", "", "number on ordered list to delete")
	cmd.MarkFlagRequired("number")

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	log.Println("Deleting a todo doc...")
}
