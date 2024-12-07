package Handlers

import (
	"net/http"
	"os"

	utils "forum/BackEnd/helpers"
)

func HandleMain(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		utils.ErrorWriter(w, "Error 404", 404)
		return
	}
	Data, err := os.ReadFile("FrontEnd/Templates/index.html")
	if err != nil {
		utils.ErrorWriter(w, "Error 500", 500)
		return
	}
	w.Write(Data)
}
