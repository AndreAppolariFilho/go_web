package main

import (
	"time"

	"github.com/AndreAppolariFilho/go_rssaggregator/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string `json:"name"`
	ApiKey    string `json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description *string `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url         string `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func databaseUserToUser(dbUser database.User) User{
	return User{
		ID: dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name: dbUser.Name,
		ApiKey: dbUser.ApiKey,
	}
}

func databaseFeedToFeed(dbFeed database.Feed) Feed{
	return Feed{
		ID: dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name: dbFeed.Name,
		Url: dbFeed.Url,
		UserID: dbFeed.UserID,
	}
}


func databaseFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow{
	return FeedFollow{
		ID: dbFeedFollow.ID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		UserID: dbFeedFollow.UserID,
		FeedID: dbFeedFollow.FeedID,
	}
}


func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed{
	feeds := []Feed{}
	for _, dbFeed := range dbFeeds{
		feeds = append(feeds, databaseFeedToFeed(dbFeed))
	}
	return feeds
}


func databaseFeedFollowsToFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow{
	feedFollows := []FeedFollow{}
	for _, dbFeedFollow := range dbFeedFollows{
		feedFollows = append(feedFollows, databaseFeedFollowToFeedFollow(dbFeedFollow))
	}
	return feedFollows
}

func databasePostToPost(dbPost database.Post) Post{
	var description *string
	if dbPost.Description.Valid{
		description = &dbPost.Description.String
	}
	return Post{
		ID: dbPost.ID,
		CreatedAt: dbPost.CreatedAt,
		PublishedAt: dbPost.PublishedAt,
		Title: dbPost.Title,
		Description: description,
		UpdatedAt: dbPost.UpdatedAt,
		Url: dbPost.Url,
		FeedID: dbPost.FeedID,
	}
}

func databasePostsToPosts(dbPosts []database.Post) []Post{
	posts := []Post{}
	for _, dbPost := range dbPosts{
		posts = append(posts, databasePostToPost(dbPost))
	}
	return posts
}