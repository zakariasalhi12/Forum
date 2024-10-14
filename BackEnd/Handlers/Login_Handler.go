package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	tem, err := template.ParseFiles("FrontEnd/templates/login.html")
	if err != nil {
		log.Fatal(err)
	}

	tem.Execute(w, nil)
}
