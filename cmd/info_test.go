package cmd_test

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	"github.com/acco32/changelog/cmd"
)

func TestSummaryDisplaysCopyright(t *testing.T) {
	infoRootCmd := &cobra.Command{Use: "changelog"}
	infoRootCmd.AddCommand(cmd.InfoCmd)

	_, out, err := executeCommandC(infoRootCmd, "info")
	assert.Contains(t, out, "© 2018")
	assert.Nil(t, err)
}

func TestErrorWithInvalidArgNumber(t *testing.T) {
	infoRootCmd := &cobra.Command{Use: "changelog"}
	infoRootCmd.AddCommand(cmd.InfoCmd)

	_, out, err := executeCommandC(infoRootCmd, "info", "blah")
	assert.Contains(t, out, "Error")
	assert.NotNil(t, err)
}
