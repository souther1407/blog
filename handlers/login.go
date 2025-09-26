package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/souther1407/blog/helpers"
	"github.com/souther1407/blog/interfaces"
)

type LoginBody struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request, apiConfig interfaces.ApiConfig) {
	body, err := helpers.GetBody[LoginBody](r)

	if err != nil {
		log.Println("Error parsing body", err)
		helpers.ResponseWithError(w, 400, fmt.Sprintf("Error parsing body %v", err))
		return
	}

	cookieStore := apiConfig.CookieStore
	db := apiConfig.DB
	session, err := cookieStore.Get(r, "sessionid")
	if err != nil {
		log.Println("Error getting sessionid", err)
		helpers.ResponseWithError(w, 400, fmt.Sprintf("Error getting session token %v", err))
		return
	}

	if session.Values["authenticated"] == true {
		helpers.ResponseWithError(w, 400, "User already authenticated")
		return
	}

	user, err := db.GetUser(r.Context(), body.Name)

	if err != nil {
		log.Println("Error getting user by name", err)
		helpers.ResponseWithError(w, 400, "Error by authenticate :c")
		return
	}

	if !helpers.CheckPasswordHash(body.Password, user.Password) {
		helpers.ResponseWithError(w, 401, "User and/or password not exists")
		return
	}

	session.Values["authenticated"] = true
	session.Values["userid"] = user.ID.String()
	err = session.Save(r, w)
	if err != nil {
		log.Println("Error setting sessionid", err)
		helpers.ResponseWithError(w, 400, fmt.Sprintf("Error setting sessionid %v", err))
		return
	}

	helpers.ResponseWithJSON(w, 200, struct{ Status string }{Status: "ok"})
}

func Logout(w http.ResponseWriter, r *http.Request, apiConfig interfaces.ApiConfig) {
	cookieStore := apiConfig.CookieStore
	session, err := cookieStore.Get(r, "sessionid")
	if err != nil {
		log.Println("Error getting sessionid", err)
		helpers.ResponseWithError(w, 400, fmt.Sprintf("Error getting session token %v", err))
		return
	}
	session.Values["authenticated"] = false
	session.Values["userid"] = ""
	session.Save(r, w)
	helpers.ResponseWithJSON(w, 200, struct{ Status string }{Status: "ok"})
}
