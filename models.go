package main

import (
	"database/sql"
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
	LastFetchedAt *time.Time `json:"last_fetched_at"`
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
		LastFetchedAt: nullTimeToTimePtr(feed.LastFetchedAt),
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

type Post struct {
	ID          uuid.UUID  `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Title       string     `json:"title"`
	Url         string     `json:"url"`
	Description *string    `json:"description"`
	PublishedAt *time.Time `json:"published_at"`
	FeedID      uuid.UUID  `json:"feed_id"`
}

func databasePostToPost(post database.Post) Post {
	return Post{
		ID:          post.ID,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		Title:       post.Title,
		Url:         post.Url,
		Description: nullStringToStringPtr(post.Description),
		PublishedAt: nullTimeToTimePtr(post.PublishedAt),
		FeedID:      post.FeedID,
	}
}

func databasePostsToPosts(posts []database.Post) []Post {
	result := make([]Post, len(posts))
	for i, post := range posts {
		result[i] = databasePostToPost(post)
	}
	return result
}

func nullTimeToTimePtr(t sql.NullTime) *time.Time {
	if t.Valid {
		return &t.Time
	}
	return nil
}

func nullStringToStringPtr(s sql.NullString) *string {
	if s.Valid {
		return &s.String
	}
	return nil
}