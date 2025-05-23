# Gator - RSS Blog Aggregator

A command-line RSS feed aggregator built with Go and PostgreSQL.

## Prerequisites

- **Go** (1.19 or later)
- **PostgreSQL** database

## Installation

Install the gator CLI:

```bash
go install
```

## Setup

1. Create a `.gatorconfig.json` file in your home directory:

```json
{
  "db_url": "postgres://username:password@localhost/gator?sslmode=disable",
  "current_user_name": ""
}
```

2. Set up your PostgreSQL database and update the `db_url` with your credentials.

3. Run the database migrations (the program will handle this automatically).

## Usage

### User Management
```bash
gator register <username>    # Create a new user
gator login <username>       # Switch to a user
gator users                  # List all users
```

### Feed Management
```bash
gator addfeed <name> <url>   # Add an RSS feed
gator feeds                  # List all feeds
gator follow <feed_name>     # Follow a feed
gator following              # List feeds you're following
gator unfollow <feed_name>   # Unfollow a feed
```

### Reading Posts
```bash
gator agg                    # Fetch latest posts from feeds
gator browse [limit]         # Browse recent posts
```

### Other
```bash
gator reset                  # Reset all data
```

## Example Workflow

1. Register and login: `gator register john && gator login john`
2. Add a feed: `gator addfeed "Go Blog" https://go.dev/blog/feed.atom`
3. Follow the feed: `gator follow "Go Blog"`
4. Fetch posts: `gator agg`
5. Browse posts: `gator browse 10`
