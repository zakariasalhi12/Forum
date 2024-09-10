package forum

import (
	"log"
	"net/http"
	"html/template"
)

func HandleMain(w http.ResponseWriter, r *http.Request) {
	tem, err := template.ParseFiles("FrontEnd/templates/index.html")
	if err != nil {
		log.Fatal(err)
	}

	tem.Execute(w, nil)
}
