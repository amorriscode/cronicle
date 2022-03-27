package todo

import (
	"cronicle/cmd/todo/create"
	"cronicle/cmd/todo/delete"
	"cronicle/cmd/todo/list"
	"cronicle/cmd/todo/update"
	"cronicle/utils"
	"path/filepath"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "todo",
		Short: "Manage your todos",
		Long:  "Manage the todos in your cronicle journal.",
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
	utils.CreateDirIfNotExist(filepath.Join(d, "todo"))
}
