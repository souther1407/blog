package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorJson struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

func ResponseWithError(w http.ResponseWriter, code int, errorMessage string) {
	errorResponse := ErrorJson{
		Code:  code,
		Error: errorMessage,
	}
	data, _ := json.Marshal(errorResponse)
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)

}

func ResponseWithJSON(w http.ResponseWriter, code int, payload any) {
	data, err := json.Marshal(payload)

	if err != nil {
		log.Println("Error al convertir de struct a json: ", err)
		ResponseWithError(w, 400, err.Error())

	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(data)
	}

}
