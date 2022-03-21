package delete

import (
	"cronicle/ui/constants"
	"cronicle/utils"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a todo entry",
		Long:  "Delete a todo entry in your cronicle journal.",
		Run:   run,
	}

	cmd.Flags().IntP("number", "n", 0, "number on ordered list to delete")
	cmd.MarkFlagRequired("number")

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	n, _ := cmd.Flags().GetInt("number")
	if n == 0 {
		return
	}
	dirPath := utils.GetPath([]string{"todo"})
	files, _ := ioutil.ReadDir(dirPath)
	e := os.Remove(filepath.Join(dirPath, files[n-1].Name()))

	if e != nil {
		log.Fatal(constants.ERROR_DELETE_FILE, e)
	}

}
