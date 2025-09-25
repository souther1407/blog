package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/souther1407/blog/helpers"
	"github.com/souther1407/blog/interfaces"
	"github.com/souther1407/blog/internal/database"
)

func CreateUser(w http.ResponseWriter, r *http.Request, apiConfig interfaces.ApiConfig) {
	type Body struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	body, err := helpers.GetBody[Body](r)

	if err != nil {
		log.Println("In \"CreateUser\", error parsing body json", err)
		helpers.ResponseWithError(w, 400, "Error at parsing json")
		return
	}
	passwordHash, err := helpers.HashPassword(body.Password)
	if err != nil {
		log.Println("In \"CreateUser\", error hashing password", err)
		helpers.ResponseWithError(w, 400, "Error at password hashing")
		return
	}
	newUser, err := apiConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      body.Name,
		Email:     body.Email,
		Password:  passwordHash,
	})

	if err != nil {
		log.Println("In \"CreateUser\", error creating user", err)
		helpers.ResponseWithError(w, 400, "Error at creating user")
		return
	}

	helpers.ResponseWithJSON(w, 201, interfaces.DBUserToUser(newUser))

}
