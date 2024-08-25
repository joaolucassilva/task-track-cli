package cmd

import (
	"github.com/joaolucassilva/task-track-cli/internal/entities"
	"github.com/joaolucassilva/task-track-cli/internal/handlers"
	"github.com/spf13/cobra"
)

func Execute() error {
	rootCmd := &cobra.Command{}

	rootCmd.AddCommand(addTasks)
	rootCmd.AddCommand(changeDoneStatus)
	rootCmd.AddCommand(changeInProgressStatus)
	rootCmd.AddCommand(updateDescriptionCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(listCmd)

	return rootCmd.Execute()
}

var addTasks = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		handlers.Add(args[0])
	},
}

var changeDoneStatus = &cobra.Command{
	Use:   "mark-done",
	Short: "Mark a task as completed",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		handlers.UpdateStatus(args[0], entities.StatusDone)
	},
}
var changeInProgressStatus = &cobra.Command{
	Use:   "mark-in-progress",
	Short: "Mark a task as in progress",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		handlers.UpdateStatus(args[0], entities.StatusInProgress)
	},
}

var updateDescriptionCmd = &cobra.Command{
	Use:   "update",
	Short: "Updating Description Task",
	Args:  cobra.MatchAll(cobra.ExactArgs(2), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		handlers.UpdateDescription(args[0], args[1])
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete Task By ID",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		handlers.DeleteById(args[0])
	},
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
