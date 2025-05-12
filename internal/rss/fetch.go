package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"
)

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	// Create a request with the provided context
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Set appropriate headers
	req.Header.Set("User-Agent", "Gator")

	// Create an HTTP client with timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Execute the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error fetching feed: %w", err)
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"feed returned non-200 status code: %d",
			resp.StatusCode,
		)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// Parse XML into RSSFeed struct
	var feed RSSFeed
	if err := xml.Unmarshal(body, &feed); err != nil {
		return nil, fmt.Errorf("error parsing XML feed: %w", err)
	}

	// Unescape HTML entities in channel fields
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	// Unescape HTML entities in item fields
	for i := range feed.Channel.Item {
		feed.Channel.Item[i].Title = html.UnescapeString(
			feed.Channel.Item[i].Title,
		)
		feed.Channel.Item[i].Description = html.UnescapeString(
			feed.Channel.Item[i].Description,
		)
	}

	return &feed, nil
}
