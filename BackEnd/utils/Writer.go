package utils

import (
	"encoding/json"
	"net/http"
)

func Writer(w http.ResponseWriter, response any, status int) {
	res, err := json.Marshal(response)
	if err != nil {
		Writer(w, map[string]string{"Error": "An unexpected error occurred. Please try again later."}, 500)
		return
	}
	w.WriteHeader(status)
	w.Write(res)
}
