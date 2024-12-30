package helpers

import (
	"encoding/json"
	"errors"
	"html/template"
	"io"
	"net/http"
	"strings"

	"forum/BackEnd/config"
)

var (
	ErrMethod         = errors.New("method not allowed")
	ErrServer         = errors.New("an unexpected error occurred. please try again later")
	ErrInvalidRequest = errors.New("invalid request")
)

type Error struct {
	Error      string
	StatusCode int
}

func ErrorWriter(w http.ResponseWriter, Err string, stcode int) {
	NewErr := Error{
		Error:      Err,
		StatusCode: stcode,
	}
	t, err := template.ParseFiles("FrontEnd/Templates/Error.html")
	if err != nil {
		config.Config.ServerLogGenerator(err.Error())
		http.Error(w, "Error Parsing file", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(NewErr.StatusCode)
	t.Execute(w, NewErr)
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

func RemoveExtraSpaces(s string) string {
	words := strings.Fields(s)
	return strings.Join(words, " ")
}
