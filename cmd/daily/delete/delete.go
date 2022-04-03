package delete

import (
	"cronicle/utils/daily"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <ID>",
		Short: "delete a daily file",
		Long:  "delete a daily file",
		Run:   run,
	}

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	daily.DeleteDaily()
}
