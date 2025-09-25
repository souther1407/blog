package models

import (
	"github.com/google/uuid"
	"github.com/souther1407/blog/internal/database"
)

type Post struct {
	Id    uuid.UUID `json:"id"`
	Title string    `json:"title"`
}

func ParseDBPostToPost(post database.CreatePostRow) Post {
	return Post{
		Id:    post.ID,
		Title: post.Title,
	}
}
