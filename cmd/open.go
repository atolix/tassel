/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/atolix/tassel/bookmark"
	"github.com/spf13/cobra"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open the Chrome bookmarks",
	Run: func(cmd *cobra.Command, args []string) {
		path := os.Getenv("BOOKMARK_PATH")
		if path == "" {
			fmt.Println("The environment variable BOOKMARK_PATH is not set.")
			return
		}

		bookmarks, err := bookmark.ReadBookmarks(path)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		for _, bookmark := range bookmarks {
			fmt.Println("Name: %s, URL: %s\n", bookmark.Name, bookmark.Url)
		}
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}
