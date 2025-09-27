package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
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

type UpdatePostBody struct {
	Title       string `json:"title,omitempty"`
	Content     string `json:"content,omitempty"`
	Description string `json:"description,omitempty"`
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

func UpdatePostHandler(w http.ResponseWriter, r *http.Request, apiConfig interfaces.ApiConfig, userId uuid.UUID) {
	body, err := helpers.GetBody[UpdatePostBody](r)
	if err != nil {
		helpers.ResponseWithError(w, 400, "Error, invalid body format")
		return
	}
	postIdParam := chi.URLParam(r, "post_id")
	postUUID, err := uuid.Parse(postIdParam)
	if err != nil {
		helpers.ResponseWithError(w, 400, "Error, invalid post UUID format")
		return
	}
	updatedPost, err := apiConfig.DB.UpdatePost(r.Context(), database.UpdatePostParams{
		ID:          postUUID,
		UpdatedAt:   time.Now().UTC(),
		Title:       body.Title,
		Content:     body.Content,
		Description: sql.NullString{String: body.Description, Valid: body.Description != ""},
	})

	if err != nil {
		log.Println("In \"UpdatePostHandler\", error at updating  post", err)
		helpers.ResponseWithError(w, 500, "Error at updating post")
		return
	}

	helpers.ResponseWithJSON(w, 200, models.Post{Id: updatedPost.ID, Title: updatedPost.Title, Description: updatedPost.Description.String})

}

func GetLastPostsHandler(w http.ResponseWriter, r *http.Request, apiConfig interfaces.ApiConfig) {
	lastPosts, err := apiConfig.DB.GetLastPosts(r.Context(), 5)
	if err != nil {
		log.Println("Error getting lasts posts", err)
		helpers.ResponseWithError(w, 500, "Error, getting lasts posts")
		return
	}

	helpers.ResponseWithJSON(w, 200, models.ParseDBPostsToPost(lastPosts))

}

func DeletePostHandler(w http.ResponseWriter, r *http.Request, apiConfig interfaces.ApiConfig, userId uuid.UUID) {
	postIdParam := chi.URLParam(r, "post_id")
	postUUID, err := uuid.Parse(postIdParam)

	if err != nil {
		helpers.ResponseWithError(w, 400, "Error, invalid post id format")
		return
	}

	err = apiConfig.DB.DeletePost(r.Context(), database.DeletePostParams{
		ID:       postUUID,
		AuthorID: userId,
	})

	if err != nil {
		log.Println("In \"DeletePostHandler\" Error at deleting a post", err)
		helpers.ResponseWithError(w, 400, "Error at deleting the post")
		return
	}

	helpers.ResponseWithJSON(w, 200, struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	})
}
