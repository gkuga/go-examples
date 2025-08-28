package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "A simple CLI app",
}

func Execute() error {
	return rootCmd.Execute()
}
