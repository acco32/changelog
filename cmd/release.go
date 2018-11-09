package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(releaseCmd)
}

var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Generate Changelog",
	Long:  "Use data template and generate/append latest changes to changelog. This will delete data entries in unreleased folder.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Release function")
	},
}
