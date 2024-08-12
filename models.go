package main

import (
	"time"

	"github.com/eduvedras/Blog-Aggregator/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

type Feed struct {
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name string `json:"name"`
	Url string `json:"url"`
	User_id uuid.UUID `json:"user_id"`
}

type FeedFollow struct{
	ID uuid.UUID `json:"id"`
	FeedID uuid.UUID `json:"feed_id"`
	UserID uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func databaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.Apikey,
	}
}

func databaseFeedToFeed(feed database.Feed) Feed {
	return Feed{
		ID: feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name: feed.Name,
		Url: feed.Url,
		User_id: feed.UserID,
	}
}

func databaseFeedsToFeeds(databaseFeeds []database.Feed) []Feed{
	feeds := []Feed{}
	for _, databaseFeed := range databaseFeeds{
		feeds = append(feeds, databaseFeedToFeed(databaseFeed))
	}
	return feeds
}

func databaseFeedFollowToFeedFollow(databaseFeedFollow database.FeedFollow) FeedFollow{
	return FeedFollow{
		ID: databaseFeedFollow.ID,
		FeedID: databaseFeedFollow.FeedID,
		UserID: databaseFeedFollow.UserID,
		CreatedAt: databaseFeedFollow.CreatedAt,
		UpdatedAt: databaseFeedFollow.UpdatedAt,
	}
}

func databaseFeedFollowsToFeedFollows(databaseFeedFollows []database.FeedFollow) []FeedFollow{
	feedFollows := []FeedFollow{}
	for _, databaseFeedFollow := range databaseFeedFollows{
		feedFollows = append(feedFollows, databaseFeedFollowToFeedFollow(databaseFeedFollow))
	}
	return feedFollows
}