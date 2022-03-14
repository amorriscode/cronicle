package create

import (
	"cronicle/utils"
	"log"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd :=
		&cobra.Command{
			Use:   "create",
			Short: "Create a new todo",
			Long:  "Create a new todo in your cronicle journal.",
			Run:   run,
		}

	cmd.Flags().StringP("message", "m", "", "content of your todo")
	cmd.MarkFlagRequired("message")

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	message, _ := cmd.Flags().GetString("message")
	// load storage directory from config
	d := utils.GetStorageDir()
	fn := filepath.Join(d, uuid.NewString()+".txt")

	f, err := os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := f.Write([]byte(message)); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

}
