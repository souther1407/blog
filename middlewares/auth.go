package middlewares

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/souther1407/blog/helpers"
	"github.com/souther1407/blog/interfaces"
)

func GetAuth(next func(http.ResponseWriter, *http.Request, interfaces.ApiConfig, uuid.UUID), apiconfig interfaces.ApiConfig) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionStorage := apiconfig.CookieStore
		session, err := sessionStorage.Get(r, "sessionId")
		if err != nil {
			helpers.ResponseWithError(w, 401, "Unauthorized")
			return
		}

		if session.Values["authenticated"] == false {
			helpers.ResponseWithError(w, 401, "Unauthorized")
			return
		}

		userId := session.Values["userid"].(uuid.UUID)
		next(w, r, apiconfig, userId)
	}
}
