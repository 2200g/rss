package feeds
import (
	"fmt"
	"time"

	"github.com/mmcdole/gofeed"
)

func fetchAndParse(feedurl string) (*gofeed.Feed, error){
parser := gofeed.NewParser()
	feed, err := parser.ParseURL(feedURL)
	if err != nil {
		return nil, fmt.Errorf("error fetching and parsing feed: %v", err)
	}
	return feed, nil
}

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

func formatDate(date *time.Time) string {
	if date != nil {
		return date.Format("2006-01-02 15:04:05")
	}
	return "N/A"
}
