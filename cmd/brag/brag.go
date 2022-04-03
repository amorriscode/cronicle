package brag

import (
	"cronicle/cmd/brag/create"
	"cronicle/cmd/brag/delete"
	"cronicle/cmd/brag/list"
	"cronicle/cmd/brag/update"
	"cronicle/utils"
	"path/filepath"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "brag",
		Short: "manage your brag doc",
		Long:  "manage the brag doc",
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
	utils.CreateDirIfNotExist(filepath.Join(d, "brag"))
}
