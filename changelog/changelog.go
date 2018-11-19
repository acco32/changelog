package changelog

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/smallfish/simpleyaml"
)

//Changelog Template for type of change being applied
type Changelog struct {
	Name        string
	Description string
}

// ToString Prints friendly description
func (c *Changelog) ToString() string {
	return fmt.Sprintf("%12s - %s", c.Name, c.Description)
}

//Entry A changelog entry
type Entry struct {
	Title  string
	Author string
	Type   Changelog
}

var (
	//New New Feature
	New = Changelog{Name: "new", Description: "New feature"}
	//Fixed Bug Fix
	Fixed = Changelog{Name: "fixed", Description: "Bug fix"}
	//Changed Feature change
	Changed = Changelog{Name: "changed", Description: "Feature change"}
	//Deprecated New deprecation
	Deprecated = Changelog{Name: "deprecated", Description: "Feature deprecation"}
	//Removed Feature removal
	Removed = Changelog{Name: "removed", Description: "Feature removal"}
	//Security Security Fix
	Security = Changelog{Name: "security", Description: "Security fix"}
	//Performance Performance Improvement
	Performance = Changelog{Name: "performance", Description: "Performance improvement"}
	//Other Unknown Changelog type
	Other = Changelog{Name: "other", Description: "Other"}
)

const (
	//DefaultChangelogFolder All entries are stored in an unreleased folder.
	DefaultChangelogFolder = "unreleased"

	// DefaultChangelogFile The name of the file to append rendered content
	DefaultChangelogFile = "CHANGELOG.md"

	// DefaultChangelogDataFile The name of the file to append unreleased data
	DefaultChangelogDataFile = "changelog.yml"

	// DefaultChangelogHeader The header text to seek which is where we append changes
	DefaultChangelogHeader = "# Changelog"
)

//CreateChangelogEntry Create a new entry for the changelog. Usually done on a new branch.
func CreateChangelogEntry(file Entry, unreleasedFolder string) error {

	if !validTitle(file.Title) {
		return fmt.Errorf("Title must be between 1 and 120 characters (including spaces). Current Length %d", len(file.Title))
	}

	if !validAuthor(file.Author) {
		return fmt.Errorf("Author must be between 1 and 80 characters. Current Length %d", len(file.Author))
	}

	filename := path.Join(unreleasedFolder, strings.ToLower(strings.Replace(file.Title, " ", "_", -1))+".yml")
	if fileExists(filename) {
		return errors.New("File already exists")
	}

	os.Mkdir(unreleasedFolder, 0777)

	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	yaml := fmt.Sprintf("-\n  title: %s\n  author: %s\n  type: %s\n", file.Title, file.Author, file.Type.Name)

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
func CreateChangelog(unreleasedFolder string, outputFilename string) error {

	_, err := os.Stat(unreleasedFolder)
	if err != nil {
		return fmt.Errorf("\"%s\" folder does not exist", unreleasedFolder)
	}

	files, err := ioutil.ReadDir(unreleasedFolder)
	if err != nil {
		return fmt.Errorf("Unable to read files in unreleased folder: %s", err.Error())
	}

	if len(files) == 0 {
		return fmt.Errorf("At least one file must exist")
	}

	var buffer bytes.Buffer

	for _, f := range files {
		dat, err := ioutil.ReadFile(path.Join(unreleasedFolder, f.Name()))
		if err != nil {
			return fmt.Errorf("Error reading file %s", f.Name())
		}

		buffer.WriteString(string(dat))
	}

	cl, _ := os.Create(outputFilename)
	defer cl.Close()

	cl.Write(buffer.Bytes())

	return nil
}

// Text Generate a text representation of latest changes
func Text(version string, outputFilename string) (string, error) {

	if version == "" {
		return "", errors.New("version string cannot be empty")
	}

	data, err := ioutil.ReadFile(outputFilename)
	if err != nil {
		return "", err
	}

	y, err := simpleyaml.NewYaml(data)
	if err != nil {
		return "", fmt.Errorf("Problem reading data file: %s", err.Error())
	}

	if !y.IsArray() {
		return "", fmt.Errorf("input yaml has incorrect schema layout. Expected only an list of items")
	}

	text := fmt.Sprintf("%s:\n", version)
	totalElements, _ := y.GetArraySize()
	for cl := 0; cl < totalElements; cl++ {
		featureTitle, _ := y.GetIndex(cl).Get("title").String()
		featureType, _ := y.GetIndex(cl).Get("type").String()
		text = fmt.Sprintf("%s%12s | %s\n", text, featureType, featureTitle)
	}

	return text, nil
}

// Markdown Generate markdown representation of latest changes
func Markdown(version string, dataFilename string) (string, error) {

	if version == "" {
		return "", errors.New("version string cannot be empty")
	}

	data, err := ioutil.ReadFile(dataFilename)
	if err != nil {
		return "", err
	}

	y, err := simpleyaml.NewYaml(data)
	if err != nil {
		return "", fmt.Errorf("Problem reading data file: %s", err.Error())
	}

	if !y.IsArray() {
		return "", fmt.Errorf("input yaml has incorrect schema layout. Expected only an list of items")
	}

	text := fmt.Sprintf("## %s\n\n", version)
	totalElements, _ := y.GetArraySize()
	for cl := 0; cl < totalElements; cl++ {
		featureTitle, _ := y.GetIndex(cl).Get("title").String()
		featureType, _ := y.GetIndex(cl).Get("type").String()
		text = fmt.Sprintf("%s**%s**  %s  \n", text, strings.ToUpper(featureType), featureTitle)
	}
	return text, nil
}

// AppendChanges Add the latest changes to current changelog file
func AppendChanges(changelogFile string, changes string) error {
	_, err := os.Stat(changelogFile)
	if err != nil {
		return fmt.Errorf("%s not found", changelogFile)
	}

	clf, err := ioutil.ReadFile(changelogFile)
	if err != nil {
		return fmt.Errorf("Problem reading %s file", changelogFile)
	}

	text := string(clf)
	if !strings.Contains(text, DefaultChangelogHeader) {
		return fmt.Errorf("Missing Changelog Header, \"%s\", in %s file", DefaultChangelogHeader, changelogFile)
	}

	newChanges := strings.Replace(text, DefaultChangelogHeader, fmt.Sprintf("%s\n\n%s", DefaultChangelogHeader, changes), 1)
	return ioutil.WriteFile(changelogFile, []byte(newChanges), 0644)
}
