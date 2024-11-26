package handlers

import (
	"html/template"
	"log"
	"net/http"

	helpers "forum/BackEnd/helpers/Api"
)

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	tem, err := template.ParseFiles("FrontEnd/templates/register.html")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := helpers.GetUserID(r); err == nil {
		http.Redirect(w, r, "/", http.StatusFound)
	}

	tem.Execute(w, nil)
}
