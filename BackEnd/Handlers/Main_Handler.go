package handlers

import (
	"html/template"
	"log"
	"net/http"

	utils "forum/BackEnd/helpers"
)

func HandleMain(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		utils.ErrorWriter(w, "Error 404", 404)
		return
	}
	tem, err := template.ParseFiles("FrontEnd/templates/index.html")
	if err != nil {
		log.Fatal(err)
	}

	tem.Execute(w, nil)
}
