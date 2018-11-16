package changelog_test

import (
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
	err := changelog.CreateChangelog(unreleasedFolder)
	assert.Error(t, err)
}

func TestCreateUnreleasedChangelogTemplate(t *testing.T) {

	unreleasedFolder := uniuri.NewLen(10)
	entry1 := changelog.Entry{Title: "entry1", Author: "me", Type: changelog.Fixed}
	err := changelog.CreateChangelogEntry(entry1, unreleasedFolder)
	assert.Nil(t, err)

	entry2 := changelog.Entry{Title: "entry2", Author: "me2", Type: changelog.Deprecated}
	err = changelog.CreateChangelogEntry(entry2, unreleasedFolder)
	assert.Nil(t, err)

	changelog.CreateChangelog(unreleasedFolder)
	os.RemoveAll(unreleasedFolder)
	os.RemoveAll("changelog.yml")
}
