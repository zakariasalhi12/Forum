package utils

import (
	"fmt"
	"html/template"
	"net/http"
)

func ErrorWriter(w http.ResponseWriter, Err string, stcode int) {
	NewErr := Error{
		Error:      Err,
		StatusCode: stcode,
	}
	t, err := template.ParseFiles("templates/Error.html")
	if err != nil {
		fmt.Println("Error :" + err.Error())
		http.Error(w, "Error Parsing file", http.StatusInternalServerError)
		return
	}
	fmt.Println(Err)
	w.WriteHeader(NewErr.StatusCode)
	t.Execute(w, NewErr)
}
