package cmd

import (
	"fmt"
	"os"

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
		fmt.Println(err)
		os.Exit(1)
	}
}
