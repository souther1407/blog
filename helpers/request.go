package helpers

import (
	"encoding/json"
	"net/http"
)

func GetBody[T any](r *http.Request) (body T, err error) {
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&body)
	return body, err
}
