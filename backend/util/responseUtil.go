package util

import (
	"encoding/json"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, responseCode int, errMessage string) {
	RespondWithJson(w, responseCode, map[string]string{"error": errMessage})

}

func RespondWithJson(w http.ResponseWriter, responseCode int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseCode)
	w.Write(response)
}

func RespondWithCode(w http.ResponseWriter, responseCode int) {
	w.WriteHeader(responseCode)
}
