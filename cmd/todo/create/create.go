package create

import (
	"cronicle/utils"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd :=
		&cobra.Command{
			Use:   "create",
			Short: "create a new todo",
			Long:  "create a new todo in your cronicle journal.",
			Run:   run,
		}

	cmd.Flags().StringP("message", "m", "", "content of your todo")
	cmd.Flags().StringP("date", "d", "", "due date YYYY-MM-DD")
	cmd.Flags().StringP("tags", "t", "", "comma separated tags of your todo")
	cmd.MarkFlagRequired("message")

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	m, _ := cmd.Flags().GetString("message")
	d, _ := cmd.Flags().GetString("date")
	t, _ := cmd.Flags().GetString("tags")
	todo := utils.ComposeTodo(utils.WriteParams{Message: m, Date: d, Tags: t})

	utils.WriteToFile(todo, utils.GetPath([]string{"todo", uuid.NewString() + ".md"}))
	utils.ListTodos()
}
