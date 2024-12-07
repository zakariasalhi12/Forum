package helpers

import (
	"html/template"
	"net/http"
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
		http.Error(w, "Error Parsing file", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(NewErr.StatusCode)
	t.Execute(w, NewErr)
}
