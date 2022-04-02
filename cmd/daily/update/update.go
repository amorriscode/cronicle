package update

import (
	"log"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a daily file",
		Long:  "Update a daily file in your cronicle journal.",
		Run:   run,
	}

	cmd.Flags().IntP("number", "n", 0, "number on ordered list to update")
	cmd.MarkFlagRequired("number")
	cmd.Flags().StringP("message", "m", "", "content of your todo")

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	log.Println("Updating a log doc...")
}
