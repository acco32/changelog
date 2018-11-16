package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	// VERSION The version of the tool
	version = "0.0.0"

	// COPYRIGHT Copyright information
	copyright = "© 2018 Matthew Moore"

	// LICENCE Licence information
	licence = "MIT Licence"
)

func init() {
}

var InfoCmd = &cobra.Command{
	Use:   "info",
  Short: "Software Summary",
  Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
    msg, _ := info()
    cmd.Println(msg)
  },
}

func info() (string, error) {
  return fmt.Sprintf("changelog %s\n\nVersion: %s\nLicence: %s \n", copyright, version, licence), nil
}

