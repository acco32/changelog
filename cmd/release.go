package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var dryRun bool

func init() {
  releaseCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "Do no delete existing unreleased files.")
  
	rootCmd.AddCommand(releaseCmd)
}

var releaseCmd = &cobra.Command{
	Use:   "release",
  Short: "Generate Changelog",
  Args: cobra.ExactArgs(1),
	Long:  "Use data template and generate/append latest changes to changelog. This will delete data entries in unreleased folder.",
	Run: release,
}

func release(cmd *cobra.Command, args []string) {
  fmt.Println("Release function")
}