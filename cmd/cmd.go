package cmd

import (
	"github.com/spf13/cobra"
)

func displayErrorAndExit(cmd *cobra.Command, e error) {
	cmd.Printf("Error: %s", e.Error())
}
