// feeds/parser.go
package feeds

import (
	"fmt"
	"time"

	"github.com/mmcdole/gofeed"
)

// fetchAndParse fetches and parses the RSS feed from the given URL.
func fetchAndParse(feedURL string) (*gofeed.Feed, error) {
	parser := gofeed.NewParser()
	feed, err := parser.ParseURL(feedURL)
	if err != nil {
		return nil, fmt.Errorf("error fetching and parsing feed: %v", err)
	}
	return feed, nil
}

// handleFeedError prints a user-friendly error message based on the encountered error.
func handleFeedError(err error) {
	fmt.Println("Error:", err)
	// You can add more specific error handling here if needed
}

// displayFeedContent displays the content of the parsed feed.
func displayFeedContent(feed *gofeed.Feed) {
	fmt.Printf("Title: %s\n", feed.Title)
	fmt.Printf("Description: %s\n", feed.Description)
	fmt.Println("Items:")
	for _, item := range feed.Items {
		fmt.Printf("- Title: %s\n", item.Title)
		fmt.Printf("  Date: %s\n", formatDate(item.PublishedParsed))
		fmt.Printf("  Link: %s\n", item.Link)
	}
}

// formatDate formats the date to a human-readable string.
func formatDate(date *time.Time) string {
	if date != nil {
		return date.Format("2006-01-02 15:04:05")
	}
	return "N/A"
}

