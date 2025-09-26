package middlewares

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/souther1407/blog/helpers"
	"github.com/souther1407/blog/interfaces"
)

func GetAuth(next func(http.ResponseWriter, *http.Request, interfaces.ApiConfig, uuid.UUID), apiconfig interfaces.ApiConfig) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionStorage := apiconfig.CookieStore
		session, err := sessionStorage.Get(r, "sessionid")
		if err != nil {
			fmt.Println(err)
			helpers.ResponseWithError(w, 401, "Unauthorized")
			return
		}
		fmt.Println(session.Values)
		if auth, ok := session.Values["authenticated"]; auth == false || !ok {
			helpers.ResponseWithError(w, 401, "Unauthorized")
			return
		}

		userId := session.Values["userid"].(string)
		userUUID, err := uuid.Parse(userId)
		if err != nil {
			log.Println("Error parsing user id UUID", err)
			helpers.ResponseWithError(w, 401, "Unauthorized")
			return
		}
		next(w, r, apiconfig, userUUID)
	}
}
