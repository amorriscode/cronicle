package cmd

import (
	cronicleLog "cronicle/cmd/log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cronicle",
	Short: "A journal in your terminal",
	Long: `As developers, it's easy to get lost in the weeds. cronicle helps you keep
track of your todo list, your daily work log, and your brag doc.

Keep track of your developer journey in the command line.`,
}

func Execute() {
	rootCmd.AddCommand(cronicleLog.New())

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
