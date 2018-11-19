package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)


var version = "0.1.0"
var copyright = "© 2018 Matthew Moore"
var licence = "MIT Licence"

var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Software Summary",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		msg, _ := info()
		cmd.Println(msg)
	},
}

func info() (string, error) {
	return fmt.Sprintf("changelog %s\n\nVersion: %s\nLicence: %s \n", copyright, version, licence), nil
}
