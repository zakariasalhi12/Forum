package helpers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

var (
	ErrMethod         = errors.New("method not allowed")
	ErrServer         = errors.New("an unexpected error occurred. please try again later")
	ErrInvalidRequest = errors.New("invalid request")
)

func Mapper(str1, str2 string) map[string]string {
	return map[string]string{str1: str2}
}

// write a json response from the given data
func Writer(w http.ResponseWriter, response any, status int) {
	res, err := json.Marshal(response)
	if err != nil {
		Writer(w, map[string]string{"Error": "An unexpected error occurred. Please try again later."}, 500)
		return
	}
	w.WriteHeader(status)
	w.Write(res)
}

// Machi wkita daba
func TagsChecker(arr []string) bool {
	return false
}

func CheckEmpty(args ...string) bool {
	for _, arg := range args {
		if arg == "" {
			return true
		}
	}
	return false
}

func ParseRequestBody(r *http.Request, Data any) (int, error) {
	Response, err := io.ReadAll(r.Body)
	if err != nil {
		return 500, ErrServer
	}
	if err := json.Unmarshal(Response, Data); err != nil {
		return 400, ErrInvalidRequest
	}
	return -1, nil
}
