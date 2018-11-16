package cmd

import (
	"fmt"
	"os"

	"github.com/acco32/changelog/changelog"
	"github.com/spf13/cobra"
)

var dryRun bool

func init() {
	ReleaseCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "Do no delete existing unreleased files.")
}

var ReleaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Generate Changelog",
	Args:  cobra.NoArgs,
	Long:  "Use data template and generate/append latest changes to changelog. This will delete data entries in unreleased folder.",
	Run: func(cmd *cobra.Command, args []string) {
    err := release(args)
    if err != nil {
      cmd.Println(err.Error())
    }
	},
}

func release(args []string) error {

  cl, err := os.Stat("CHANGELOG.md")
	if err != nil {
    return fmt.Errorf("CHANGELOG.md not found")
  }

  err = changelog.CreateChangelog(changelog.DefaultChangelogFolder)
  if err != nil {
    return err
  }

	clYaml, err := os.Stat("changelog.yml")
	if err != nil {
		return fmt.Errorf("changelog.yml not found")
	}

	if dryRun {
		fmt.Println("print output to standard out")
    return nil
	}

  fmt.Printf("Append %s to %s", clYaml.Name(), cl.Name())
  // os.RemoveAll(changelog.DefaultChangelogFolder)
	return nil
}
