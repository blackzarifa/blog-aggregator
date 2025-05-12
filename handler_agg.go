package main

import (
	"context"
	"fmt"
	"gator/internal/rss"
)

func handlerAgg(s *state, cmd command) error {
	ctx := context.Background()
	feedURL := "https://www.wagslane.dev/index.xml"

	feed, err := rss.FetchFeed(ctx, feedURL)
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}

	fmt.Printf("%+v\n", feed)
	return nil
}
