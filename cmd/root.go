package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(AddCmd)
	rootCmd.AddCommand(InfoCmd)
	rootCmd.AddCommand(ReleaseCmd)
}

var rootCmd = &cobra.Command{
	Use:   "changelog",
	Short: "Generate changelog data",
}

// Execute Entry Point
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		displayErrorAndExit(rootCmd, err)
	}
}
