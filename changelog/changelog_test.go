package changelog_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/acco32/changelog/changelog"

	"github.com/dchest/uniuri"
	"github.com/stretchr/testify/assert"
)

func TestErrorWhenTitleIsNull(t *testing.T) {
	entry := changelog.Entry{Title: "", Author: "", Type: changelog.Fixed}
	err := changelog.CreateChangelogEntry(entry, "")
	assert.NotNil(t, err)
}

func TestErrorWhenTitleIsTooLong(t *testing.T) {
	title := "asdf asdf asdf adsf asdf asdf asdf asdf asdf adsf adsf asdf asdf asdf asdf asdf adsf asdf asdf asdf asdf asdf adsf adsf asdf asdf"
	entry := changelog.Entry{Title: title, Author: "", Type: changelog.Changed}
	err := changelog.CreateChangelogEntry(entry, "")
	assert.NotNil(t, err)
}

func TestErrorWhenAuthorIsNull(t *testing.T) {
	entry := changelog.Entry{Title: "some title", Author: "", Type: changelog.Security}
	err := changelog.CreateChangelogEntry(entry, "")
	assert.NotNil(t, err)
}

func TestErrorWhenFileAlreadyExists(t *testing.T) {
	title := "same title"
	entry := changelog.Entry{Title: title, Author: "me", Type: changelog.Fixed}

	unreleasedFolder := uniuri.NewLen(10)
	err := changelog.CreateChangelogEntry(entry, unreleasedFolder)
	assert.Nil(t, err)

	err = changelog.CreateChangelogEntry(entry, unreleasedFolder)
	assert.NotNil(t, err)

	os.RemoveAll(unreleasedFolder)
}

func TestErrorWhenUnreleasedFolderNotFound(t *testing.T) {
	unreleasedFolder := uniuri.NewLen(10)
	unreleasedDataFile := fmt.Sprintf("%s.yml", uniuri.NewLen(10))
	err := changelog.CreateChangelog(unreleasedFolder, unreleasedDataFile)
	assert.Error(t, err)
}

func TestErrorWhenUnreleasedFolderEmpty(t *testing.T) {

	unreleasedFolder := uniuri.NewLen(10)

	os.Mkdir(unreleasedFolder, 0644)

	err := changelog.CreateChangelog(unreleasedFolder, "")
	assert.Error(t, err)
	os.RemoveAll(unreleasedFolder)
}

func TestCreateUnreleasedChangelogDataFile(t *testing.T) {

	unreleasedFolder := uniuri.NewLen(10)
	unreleasedDataFile := fmt.Sprintf("%s.yml", uniuri.NewLen(10))

	entry1 := changelog.Entry{Title: "entry1", Author: "me", Type: changelog.Fixed}
	err := changelog.CreateChangelogEntry(entry1, unreleasedFolder)
	assert.Nil(t, err)

	entry2 := changelog.Entry{Title: "entry2", Author: "me2", Type: changelog.Deprecated}
	err = changelog.CreateChangelogEntry(entry2, unreleasedFolder)
	assert.Nil(t, err)

	changelog.CreateChangelog(unreleasedFolder, unreleasedDataFile)

	fi, err := os.Stat(unreleasedDataFile)
	assert.NotNil(t, fi)
	assert.Nil(t, err)

	os.RemoveAll(unreleasedFolder)
	os.RemoveAll(unreleasedDataFile)
}

func TestErrorWhenChangelogDataFileHasBadDataFormat(t *testing.T) {

	unreleasedDataFile := fmt.Sprintf("%s.yml", uniuri.NewLen(10))

	data := []byte{25, 85, 6, 89, 54, 13, 52, 10, 52}

	ioutil.WriteFile(unreleasedDataFile, data, 0644)
	str, err := changelog.Text(unreleasedDataFile)
	assert.Empty(t, str)
	assert.Error(t, err)

	os.RemoveAll(unreleasedDataFile)
}

func TestErrorWhenChangelogDataFileIsNotArrayInRootNode(t *testing.T) {

	unreleasedDataFile := fmt.Sprintf("%s.yml", uniuri.NewLen(10))

	data := []byte("name: author\nage: 99\nfloat: 3.14159")
	ioutil.WriteFile(unreleasedDataFile, data, 0644)
	str, err := changelog.Text(unreleasedDataFile)
	assert.Empty(t, str)
	assert.Error(t, err)

	os.RemoveAll(unreleasedDataFile)
}

func TestErrorWhenChangelogDataFileHasArrayInChildNode(t *testing.T) {

	unreleasedDataFile := fmt.Sprintf("%s.yml", uniuri.NewLen(10))

	data := []byte(`name: author
age: 99
float: 3.14159
data:
  -
    title: some feature 1
    author: some author
    type: deprecated
  -
    title: some feature 2
    author: some author
    type: fixed
`)

	ioutil.WriteFile(unreleasedDataFile, data, 0644)
	str, err := changelog.Text(unreleasedDataFile)
	assert.Empty(t, str)
	assert.Error(t, err)

	os.RemoveAll(unreleasedDataFile)
}

func TestCreateUnreleasedChangelogText(t *testing.T) {

	unreleasedDataFile := fmt.Sprintf("%s.yml", uniuri.NewLen(10))

	data := []byte(`-
  title: Some bug fix
  author: author1
  type: fixed
-
  title: Some new feature
  name: author2
  type: new
`)

	ioutil.WriteFile(unreleasedDataFile, data, 0644)
	os.Stat(unreleasedDataFile)

	str, err := changelog.Text(unreleasedDataFile)
	assert.NotEmpty(t, str)
	assert.NoError(t, err)

	os.RemoveAll(unreleasedDataFile)
}

func TestCreateUnreleasedChangelogMarkdown(t *testing.T) {

	unreleasedDataFile := fmt.Sprintf("%s.yml", uniuri.NewLen(10))

	data := []byte(`-
  title: Some bug fix
  author: author1
  type: fixed
-
  title: Some new feature
  name: author2
  type: new
`)

	ioutil.WriteFile(unreleasedDataFile, data, 0644)
	os.Stat(unreleasedDataFile)

	str, err := changelog.Markdown(unreleasedDataFile)
	assert.NotEmpty(t, str)
	assert.NoError(t, err)

	os.RemoveAll(unreleasedDataFile)
}

func TestErrorWhenNoChangelogFile(t *testing.T) {
	changelogFile := fmt.Sprintf("%s.md", uniuri.NewLen(10))
	err := changelog.AppendChanges(changelogFile, "")
	assert.Error(t, err)
}

func TestErrorWhenChangelogHeadingNotFoundInFile(t *testing.T) {
	changelogFile := fmt.Sprintf("%s.md", uniuri.NewLen(10))
	data := []byte("# Markdown")
	ioutil.WriteFile(changelogFile, data, 0644)
	err := changelog.AppendChanges(changelogFile, "")
	assert.Error(t, err)
	os.Remove(changelogFile)
}

func TestAppendChanges(t *testing.T) {

	changelogFile := fmt.Sprintf("%s.md", uniuri.NewLen(10))
	data := []byte("# Changelog")
	ioutil.WriteFile(changelogFile, data, 0644)

	unreleasedDataFile := fmt.Sprintf("%s.yml", uniuri.NewLen(10))

	data = []byte(`-
  title: Some bug fix
  author: author1
  type: fixed
-
  title: Some new feature
  name: author2
  type: new
`)

	ioutil.WriteFile(unreleasedDataFile, data, 0644)
	os.Stat(unreleasedDataFile)

	str, err := changelog.Markdown(unreleasedDataFile)
	assert.NotEmpty(t, str)
	assert.NoError(t, err)

	err = changelog.AppendChanges(changelogFile, str)
	assert.NoError(t, err)
	os.Remove(changelogFile)
	os.Remove(unreleasedDataFile)
}
