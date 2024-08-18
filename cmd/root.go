package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
