package cmd

import (
	"github.com/joaolucassilva/task-track-cli/internal/handlers"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateDescriptionCmd)
}

var updateDescriptionCmd = &cobra.Command{
	Use:   "update",
	Short: "Updating Task",
	Args:  cobra.MatchAll(cobra.ExactArgs(2), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		handlers.UpdateDescription(args[0], args[1])
	},
}
