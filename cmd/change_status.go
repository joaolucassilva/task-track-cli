package cmd

import (
	"github.com/joaolucassilva/task-track-cli/internal/entities"
	"github.com/joaolucassilva/task-track-cli/internal/handlers"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(changeDoneStatus)
	rootCmd.AddCommand(changeInProgressStatus)
}

var changeDoneStatus = &cobra.Command{
	Use:   "mark-in-progress",
	Short: "Mark a task as in progress",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		handlers.UpdateStatus(args[0], entities.StatusInProgress)
	},
}
var changeInProgressStatus = &cobra.Command{
	Use:   "mark-done",
	Short: "Mark a task as completed",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		handlers.UpdateStatus(args[0], entities.StatusDone)
	},
}
