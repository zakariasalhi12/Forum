package handlers

import (
	"log"
	"net/http"
	"html/template"
)

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	tem, err := template.ParseFiles("FrontEnd/templates/register.html")
	if err != nil {
		log.Fatal(err)
	}

	tem.Execute(w, nil)
}
