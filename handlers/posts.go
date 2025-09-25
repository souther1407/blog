package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/souther1407/blog/helpers"
	"github.com/souther1407/blog/interfaces"
	"github.com/souther1407/blog/internal/database"
	"github.com/souther1407/blog/models"
)

type CreatePostBody struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request, apiConfig interfaces.ApiConfig, userId uuid.UUID) {
	body, err := helpers.GetBody[CreatePostBody](r)
	if err != nil {
		helpers.ResponseWithError(w, 400, "Error, invalid body format")
		return
	}

	newPost, err := apiConfig.DB.CreatePost(r.Context(), database.CreatePostParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Title:     body.Title,
		Content:   body.Content,
		AuthorID:  userId,
	})

	if err != nil {
		log.Println("In \"CreatePostHandler\", error at creating new post", err)
		helpers.ResponseWithError(w, 500, "Error at creating new post")
		return
	}

	helpers.ResponseWithJSON(w, 201, models.ParseDBPostToPost(newPost))

}
