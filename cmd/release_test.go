package cmd_test

import (
	"testing"

	"github.com/acco32/changelog/cmd"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestErrorWhenNoChangelogFile(t *testing.T) {
	releaseRootCmd := &cobra.Command{Use: "changelog"}
	releaseRootCmd.AddCommand(cmd.ReleaseCmd)

	_, out, err := executeCommandC(releaseRootCmd, "release")
	assert.Contains(t, out, "Error")
	assert.Nil(t, err)
}
