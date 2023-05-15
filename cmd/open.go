/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/atolix/tassel/bookmark"
	"github.com/manifoldco/promptui"
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

		templates := &promptui.SelectTemplates{
			Label:    "{{ . }}?",
			Active:   "→ {{ .Name | cyan }} ({{ .Url | red }})",
			Inactive: "  {{ .Name | cyan }} ({{ .Url | red }})",
			Selected: "→ {{ .Name | red | cyan }}",
		}

		prompt := promptui.Select{
			Label:     "Select Bookmark",
			Items:     bookmarks,
			Templates: templates,
		}

		i, _, err := prompt.Run()

		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		fmt.Println(bookmarks[i].Url)
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}
