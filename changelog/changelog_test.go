package changelog_test

import (
	"github.com/acco32/changelog/changelog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorWhenTitleIsNull(t *testing.T) {
	entry := changelog.Entry{Title: "", Author: "", Type: changelog.Fixed}
	err := changelog.CreateChangelogEntry(entry)
	assert.NotNil(t, err)
}

func TestErrorWhenTitleIsTooLong(t *testing.T) {
	title := "asdf asdf asdf adsf asdf asdf asdf asdf asdf adsf adsf asdf asdf asdf asdf asdf adsf asdf asdf asdf asdf asdf adsf adsf asdf asdf"
	entry := changelog.Entry{Title: title, Author: "", Type: changelog.Changed}
	err := changelog.CreateChangelogEntry(entry)
	assert.NotNil(t, err)
}

func TestErrorWhenAuthorIsNull(t *testing.T) {
	entry := changelog.Entry{Title: "some title", Author: "", Type: changelog.Security}
	err := changelog.CreateChangelogEntry(entry)
	assert.NotNil(t, err)
}

func TestErrorWhenFileAlreadyExists(t *testing.T) {
	title := "same title"
	entry := changelog.Entry{Title: title, Author: "me", Type: changelog.Fixed}

	err := changelog.CreateChangelogEntry(entry)
	assert.Nil(t, err)

	err = changelog.CreateChangelogEntry(entry)
	assert.NotNil(t, err)

	os.RemoveAll(changelog.DefaultChangelogFolder)
}

func TestCreateUnreleasedChangelogTemplate (t *testing.T){

	entry1 := changelog.Entry{Title: "entry1", Author: "me", Type: changelog.Fixed}
	err := changelog.CreateChangelogEntry(entry1)
	assert.Nil(t, err)

	entry2 := changelog.Entry{Title: "entry2", Author: "me2", Type: changelog.Deprecated}
	err = changelog.CreateChangelogEntry(entry2)
	assert.Nil(t, err)

  changelog.CreateChangelog()
	os.RemoveAll(changelog.DefaultChangelogFolder)
}
