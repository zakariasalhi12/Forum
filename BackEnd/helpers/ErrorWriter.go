package helpers

import (
	"fmt"
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
	t, err := template.ParseFiles("FrontEnd/templates/Error.html")
	if err != nil {
		fmt.Println("Error :" + err.Error())
		http.Error(w, "Error Parsing file", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(NewErr.StatusCode)
	t.Execute(w, NewErr)
}
