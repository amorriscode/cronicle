package daily

import (
	"cronicle/cmd/daily/create"
	"cronicle/cmd/daily/delete"
	"cronicle/cmd/daily/list"
	"cronicle/cmd/daily/update"
	"cronicle/utils"
	"path/filepath"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "daily",
		Short: "manage your daily entries",
		Long:  "manage the daily entries",
	}
	// Ensure todo storage dir exists on execute
	cobra.OnInitialize(CreateStorageDir)

	cmd.AddCommand(create.New())
	cmd.AddCommand(update.New())
	cmd.AddCommand(delete.New())
	cmd.AddCommand(list.New())

	return cmd
}

func CreateStorageDir() {
	d := utils.GetStorageDir()
	utils.CreateDirIfNotExist(filepath.Join(d, "daily"))
}
