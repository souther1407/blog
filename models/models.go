package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/souther1407/blog/internal/database"
)

type Post struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title,omitempty"`
	Author      string    `json:"author,omitempty"`
	Content     string    `json:"content,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

func ParseDBPostToPost(post database.CreatePostRow) Post {
	return Post{
		Id:    post.ID,
		Title: post.Title,
	}
}

func ParseDBPostsToPost(posts []database.GetLastPostsRow) []Post {
	parsedPosts := []Post{}
	for _, p := range posts {
		parsedPosts = append(parsedPosts, Post{Id: p.ID, Title: p.Title, Author: p.Author, CreatedAt: p.CreatedAt, Description: p.Description.String})
	}
	return parsedPosts
}
