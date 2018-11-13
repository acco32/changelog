package cmd

import (
	"fmt"

	"github.com/acco32/changelog/changelog"
	"github.com/spf13/cobra"
)

var author string
var changelogType changelog.Changelog
var title string

func init() {
	addCmd.Flags().StringVarP(&author, "author", "a", "", "Individual creating feature.")
	addCmd.MarkFlagRequired("author")

	addCmd.Flags().StringVarP(&title, "title", "t", "", "Feature description. Space will be replaced by underscores.")
	addCmd.MarkFlagRequired("title")

	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add new changelog type",
	Args:  cobra.ExactArgs(1),
	Run:   add,
}

func add(cmd *cobra.Command, args []string) {

	parseEntryType := func(changelogType string) (changelog.Changelog, error) {

		switch changelogType {
		case changelog.New.Name:
			return changelog.New, nil
		case changelog.Fixed.Name:
			return changelog.Fixed, nil
		case changelog.Changed.Name:
			return changelog.Changed, nil
		case changelog.Removed.Name:
			return changelog.Removed, nil
		case changelog.Deprecated.Name:
			return changelog.Deprecated, nil
		case changelog.Security.Name:
			return changelog.Security, nil
		case changelog.Performance.Name:
			return changelog.Performance, nil
		}
		return changelog.Other, fmt.Errorf("Unrecognized format entered: %s", changelogType)
	}

	entryType, err := parseEntryType(args[0])
	if err != nil {
		fmt.Println("Error: %s", err.Error())
	}

	entry := changelog.Entry{Title: title, Author: author, Type: entryType}
	err = changelog.CreateChangelogEntry(entry)
	if err != nil {
		fmt.Println("Error: %s", err.Error())
	}
}
