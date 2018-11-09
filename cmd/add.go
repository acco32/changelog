package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add new changelog type",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add function")
	},
}

