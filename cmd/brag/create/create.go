package create

import (
	"cronicle/utils"
	"cronicle/utils/entries"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "create a new brag entry",
		Long:  "create a new brag entry",
		Run:   run,
	}

	cmd.Flags().StringP("message", "m", "", "content of your brag entry")
	cmd.Flags().StringP("tags", "t", "", "comma separated tags of your brag entry")

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	m, _ := cmd.Flags().GetString("message")
	t, _ := cmd.Flags().GetString("tags")
	entries.WriteOrCreateEntry(utils.WriteParams{Message: m, Tags: t}, "brag")
}
