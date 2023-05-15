package bookmark

import (
	"encoding/json"
	"io/ioutil"
)

type Bookmark struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type BookmarkNode struct {
	Type     string         `json:"type"`
	Name     string         `json:"name"`
	Url      string         `json:"url"`
	Children []BookmarkNode `json:"children"`
}

type BookmarkFile struct {
	Roots BookmarkRoots `json:"roots"`
}

type BookmarkRoots struct {
	BookmarkBar BookmarkNode `json:"bookmark_bar"`
	Other       BookmarkNode `json:"other"`
}

func ReadBookmarks(path string) ([]Bookmark, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var bookmarks BookmarkFile
	if err := json.Unmarshal(data, &bookmarks); err != nil {
		return nil, err
	}

	var result []Bookmark
	var visitNode func(node BookmarkNode)
	visitNode = func(node BookmarkNode) {
		if node.Type == "url" {
			result = append(result, Bookmark{node.Name, node.Url})
		}
		for _, child := range node.Children {
			visitNode(child)
		}
	}

	visitNode(bookmarks.Roots.BookmarkBar)
	visitNode(bookmarks.Roots.Other)

	return result, nil
}
