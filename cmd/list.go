package cmd

import (
	"github.com/joaolucassilva/task-track-cli/internal/handlers"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		filter := ""
		if len(args) != 0 {
			filter = args[0]
		}

		handlers.ListAll(filter)
	},
}
