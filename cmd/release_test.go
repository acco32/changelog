package cmd_test

import (
	"testing"

	"github.com/spf13/cobra"

	"github.com/acco32/changelog/cmd"
)

func TestSummaryDispljjaysCopyright(t *testing.T) {
	releaseRootCmd := &cobra.Command{Use: "changelog"}
	releaseRootCmd.AddCommand(cmd.ReleaseCmd)
}
