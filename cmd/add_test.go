package cmd_test

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	"github.com/acco32/changelog/changelog"
	"github.com/acco32/changelog/cmd"
)

func TestErrorWhenNoArgs(t *testing.T) {
	addRootCmd := &cobra.Command{Use: "changelog"}
	addRootCmd.AddCommand(cmd.AddCmd)

	_, out, err := executeCommandC(addRootCmd, "add")
	assert.Contains(t, out, "Error")
	assert.NotNil(t, err)
}

func init() {
	os.RemoveAll(changelog.DefaultChangelogFolder)
}

func TestNoTitleError(t *testing.T) {
	addRootCmd := &cobra.Command{Use: "changelog"}
	addRootCmd.AddCommand(cmd.AddCmd)

	_, out, err := executeCommandC(addRootCmd, "add", changelog.Fixed.Name, "-a author")
	assert.Contains(t, out, "Error")
	assert.Contains(t, out, "\"title\" not set")
	assert.NotNil(t, err)
}

func TestNoAuthorError(t *testing.T) {
	t.Skip("Test causing goroutine error. Need to investigate.")
	addRootCmd := &cobra.Command{Use: "changelog"}
	addRootCmd.AddCommand(cmd.AddCmd)

	_, out, err := executeCommandC(addRootCmd, "add", changelog.Fixed.Name, "-t title")
	assert.Contains(t, out, "Error")
	assert.Contains(t, out, "\"author\" not set")
	assert.NotNil(t, err)
}

func TestErrorWithUnknownType(t *testing.T) {
	addRootCmd := &cobra.Command{Use: "changelog", Args: cobra.ExactArgs(1)}
	addRootCmd.AddCommand(cmd.AddCmd)

	_, out, err1 := executeCommandC(addRootCmd, "add", "blah", "-t title", "-a author")
	assert.Contains(t, out, "Unrecognized format entered")
	assert.Contains(t, out, "Unrecognized format entered")
	assert.Nil(t, err1)

	_, err2 := os.Stat(changelog.DefaultChangelogFolder)
	assert.True(t, os.IsNotExist(err2))
}
