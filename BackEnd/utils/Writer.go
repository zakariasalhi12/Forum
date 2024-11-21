package utils

import (
	"encoding/json"
	"net/http"
)

func Writer(w http.ResponseWriter, response any, status int) {
	res, err := json.Marshal(response)
	if err != nil {
		ErrorWriter(w, err.Error(), 500)
		return
	}
	w.WriteHeader(status)
	w.Write(res)
}
