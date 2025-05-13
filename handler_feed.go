package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"log"
	"time"

	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	if s.cfg.CurrentUserName == "" {
		return fmt.Errorf("no user is logged in")
	}

	ctx := context.Background()
	user, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error getting current user: %w", err)
	}

	now := time.Now()
	feed, err := s.db.CreateFeed(
		ctx,
		database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: now,
			UpdatedAt: now,
			Name:      name,
			Url:       url,
			UserID:    user.ID,
		},
	)

	if err != nil {
		return fmt.Errorf("failed to create feed: %w", err)
	}

	log.Printf("Feed created: %+v\n", feed)
	return nil
}
