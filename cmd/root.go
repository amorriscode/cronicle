package cmd

import (
	"cronicle/cmd/brag"
	"cronicle/cmd/daily"
	"cronicle/cmd/todo"
	"cronicle/config"
	"cronicle/ui"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cronicle",
	Short: "A journal in your terminal",
	Long: `As developers, it's easy to get lost in the weeds. cronicle helps you keep
track of your todo list, your daily work log, and your brag doc.

Keep track of your developer journey in the command line.`,
	Run: run,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) {
	p := tea.NewProgram(ui.New(), tea.WithAltScreen(), tea.WithMouseCellMotion())

	if err := p.Start(); err != nil {
		log.Printf("Darn, something went wrong: %v", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(config.InitConfig)

	rootCmd.AddCommand(daily.New())
	rootCmd.AddCommand(brag.New())
	rootCmd.AddCommand(todo.New())
}
