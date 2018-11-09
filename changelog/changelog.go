package changelog

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

//Changelog Template for type of change being applied
type Changelog struct {
	Name        string
	Description string
}

//Entry A changelog entry
type Entry struct {
	Title  string
	Author string
	Type   Changelog
}

var (
	//Added New Feature
	Added = Changelog{Name: "added", Description: "New feature"}
	//Fixed Bug Fix
	Fixed = Changelog{Name: "fixed", Description: "Bug fix"}
	//Changed Feature change
	Changed = Changelog{Name: "changed", Description: "Feature change"}
	//Deprecated New deprecation
	Deprecated = Changelog{Name: "deprecated", Description: "New deprecation"}
	//Removed Feature removal
	Removed = Changelog{Name: "removed", Description: "Feature removal"}
	//Security Security Fix
	Security = Changelog{Name: "security", Description: "Security fix"}
	//Performance Performance Improvement
	Performance = Changelog{Name: "performance", Description: "Performance improvement"}
	//Other Unknown Changelog type
	Other = Changelog{Name: "other", Description: "Other"}
)

//DefaultChangelogFolder All entries are stored in an unreleased folder.
const DefaultChangelogFolder = "unreleased"

//CreateChangelogEntry Create a new entry for the changelog. Usually done on a new branch.
func CreateChangelogEntry(file Entry) error {

	if !validTitle(file.Title) {
		return fmt.Errorf("Title must be between 1 and 120 characters (including spaces). Current Length %d", len(file.Title))
	}

	if !validAuthor(file.Author) {
		return fmt.Errorf("Author must be between 1 and 80 characters. Current Length %d", len(file.Author))
	}

	filename := path.Join(DefaultChangelogFolder, strings.ToLower(strings.Replace(file.Title, " ", "_", -1))+".yml")
	if fileExists(filename) {
		return errors.New("File already exists")
	}

	os.Mkdir(DefaultChangelogFolder, 0777)

	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	yaml := fmt.Sprintf("---\ntitle: %s\nauthor: %s\ntype: %s\n", file.Title, file.Author, file.Type.Name)
	if _, err := f.WriteString(yaml); err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	return nil
}

func fileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func validTitle(title string) bool {
	return len(title) > 0 && len(title) < 120
}

func validAuthor(author string) bool {
	return len(author) > 0 && len(author) < 80
}

// CreateChangelog Gather all unreleased entries and create latest log
func CreateChangelog() error {
	
	files, err := ioutil.ReadDir(DefaultChangelogFolder)
	if err != nil {
		return fmt.Errorf("Unable to read files in unreleased folder: %s", err.Error())
	}

	var buffer bytes.Buffer

	for _, f := range files {
		dat, err := ioutil.ReadFile(path.Join(DefaultChangelogFolder, f.Name()))
		if err != nil {
			return fmt.Errorf("Error reading file %s", f.Name())
		}

		buffer.WriteString(string(dat))
	}

	cl, _ := os.Create(path.Join(DefaultChangelogFolder, "changelog.yml"))
	defer cl.Close()

	cl.Write(buffer.Bytes())

	return nil
}
