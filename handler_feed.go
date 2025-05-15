package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func createFeedFollow(
	s *state,
	userID, feedID uuid.UUID,
) (database.CreateFeedFollowRow, error) {
	now := time.Now().UTC()
	return s.db.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: now,
			UpdatedAt: now,
			UserID:    userID,
			FeedID:    feedID,
		},
	)
}

func handlerFollowFeed(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.Name)
	}

	url := cmd.Args[0]
	if s.cfg.CurrentUserName == "" {
		return fmt.Errorf("no user is logged in")
	}

	ctx := context.Background()

	user, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error getting current user: %w", err)
	}

	feed, err := s.db.GetFeedByURL(ctx, url)
	if err != nil {
		return fmt.Errorf("feed not found with URL %s: %w", url, err)
	}

	feedFollow, err := createFeedFollow(s, user.ID, feed.ID)
	if err != nil {
		return fmt.Errorf("error following feed: %w", err)
	}

	fmt.Printf(
		"User %s is now following feed: %s\n",
		feedFollow.UserName,
		feedFollow.FeedName,
	)
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	feed, err := s.db.CreateFeed(
		context.Background(),
		database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			UserID:    user.ID,
			Name:      name,
			Url:       url,
		},
	)
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed)

	// Automatically create a feed follow for the current user
	feedFollow, err := createFeedFollow(s, user.ID, feed.ID)
	if err != nil {
		fmt.Printf("Warning: Could not automatically follow feed: %v\n", err)
	} else {
		fmt.Printf("User %s is now following feed: %s\n", feedFollow.UserName, feedFollow.FeedName)
	}

	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	ctx := context.Background()

	feeds, err := s.db.GetFeedsWithUsers(ctx)
	if err != nil {
		return fmt.Errorf("error getting feeds: %w", err)
	}

	fmt.Println("Feeds:")
	fmt.Println("------------------------------")
	for _, feed := range feeds {
		fmt.Printf("Name: %s\n", feed.Name)
		fmt.Printf("URL: %s\n", feed.Url)
		fmt.Printf("Created by: %s\n", feed.UserName)
		fmt.Println("------------------------------")
	}

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}
