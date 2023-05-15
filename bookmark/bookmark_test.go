package bookmark_test

import (
	"os"
	"testing"

	"github.com/atolix/tassel/bookmark"
)

func TestReadBookmarks(t *testing.T) {
	path := os.Getenv("BOOKMARK_PATH")
	bookmarks, err := bookmark.ReadBookmarks(path)

	if err != nil {
		t.Errorf("failed to read bookmarks: %v", err)
	}

	if len(bookmarks) == 0 {
		t.Errorf("No Bookmark was found.")
	}
}
