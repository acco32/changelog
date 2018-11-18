package cmd

import (
	"fmt"
	"os"

	"github.com/acco32/changelog/changelog"
	"github.com/spf13/cobra"
)

var preview bool

func init() {
	ReleaseCmd.Flags().BoolVarP(&preview, "preview", "p", false, "Show preview of changelog. Do not delete existing unreleased files.")
}

var ReleaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Generate Changelog",
	Args:  cobra.NoArgs,
	Long:  "Use data template and generate/append latest changes to changelog. This will delete data entries in unreleased folder.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := release(args, changelog.DefaultChangelogFile); err != nil {
			displayErrorAndExit(cmd, err)
		}
	},
}

func release(args []string, changelogFile string) error {


	err := changelog.CreateChangelog(changelog.DefaultChangelogFolder, changelog.DefaultChangelogDataFile)
	if err != nil {
		return err
	}

	clYaml, err := os.Stat(changelog.DefaultChangelogDataFile)
	if err != nil {
		return fmt.Errorf("%s not found", changelog.DefaultChangelogDataFile)
	}

	if preview {
		t, e := changelog.Text(changelog.DefaultChangelogDataFile)
		if e != nil {
			return e
		}
		fmt.Println(t)
		os.Remove(changelog.DefaultChangelogDataFile)
		return nil
	}

  fmt.Printf("Append %s to changelog", clYaml.Name())
  md, err := changelog.Markdown(changelog.DefaultChangelogDataFile)
  if err != nil {
    return err
  }

  err = changelog.AppendChanges(changelog.DefaultChangelogFile, md)
  if err != nil {
    return err
  }

	return os.RemoveAll(changelog.DefaultChangelogFolder)
}
