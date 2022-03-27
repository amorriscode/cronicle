package today

import (
	"cronicle/cmd/today/create"
	"cronicle/cmd/today/delete"
	"cronicle/cmd/today/update"
	"cronicle/utils"
	"path/filepath"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "today",
		Short: "Manage your log entries",
		Long:  "Manage the log entries in your cronicle journal.",
	}
	// Ensure todo storage dir exists on execute
	cobra.OnInitialize(CreateStorageDir)

	cmd.AddCommand(create.New())
	cmd.AddCommand(update.New())
	cmd.AddCommand(delete.New())

	return cmd
}

func CreateStorageDir() {
	d := utils.GetStorageDir()
	utils.CreateDirIfNotExist(filepath.Join(d, "today"))
}
