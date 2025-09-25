package middlewares

import (
	"net/http"

	"github.com/souther1407/blog/interfaces"
)

func InjectDB(cb func(http.ResponseWriter, *http.Request, interfaces.ApiConfig), apiConfig interfaces.ApiConfig) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cb(w, r, apiConfig)
	}
}
