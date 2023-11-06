// main.go
package main

import (
	"fmt"
	"os"

	"github.com/2200g/rss/feeds"
	"github.com/spf13/cobra"
)

var displayCmd = &cobra.Command{
	Use:   "display",
	Short: "Display the content of an RSS feed",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Usage: rss-reader display <feedURL>")
			return
		}

		feedURL := args[0]

		// Fetch and parse the feed using the feeds package
		feed, err := feeds.fetchAndParse(feedURL)
		if err != nil {
			feeds.HandleFeedError(err)
			return
		}

		// Display the content of the parsed feed using the feeds package
		feeds.displayFeedContent(feed)
	},
}

var rootCmd = &cobra.Command{Use: "rss-reader"}

func init() {
	rootCmd.AddCommand(displayCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

