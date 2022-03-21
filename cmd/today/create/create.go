package create

import (
	"log"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new today entry",
		Long:  "Create a new today entry in your cronicle journal.",
		Run:   run,
	}

	cmd.Flags().StringP("message", "m", "", "content of your todo")
	cmd.Flags().StringP("date", "d", "", "due date YYYY-MM-DD")
	cmd.Flags().StringP("tags", "t", "", "comma separated tags of your todo")
	cmd.MarkFlagRequired("message")

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	log.Println("Creating a new today entry...")
}
