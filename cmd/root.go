package cmd

import (
	cronicleLog "cronicle/cmd/log"
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
	rootCmd.AddCommand(cronicleLog.New())

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) {
	p := tea.NewProgram(ui.New(), tea.WithAltScreen())

	if err := p.Start(); err != nil {
		log.Printf("Darn, something went wrong: %v", err)
		os.Exit(1)
	}
}
