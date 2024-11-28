package handlers

import (
	"net/http"
	"os"

	utils "forum/BackEnd/helpers"
	helpers "forum/BackEnd/helpers/Api_Helper"
)

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ErrorWriter(w, "Error 400", http.StatusMethodNotAllowed)
		return
	}

	Data, err := os.ReadFile("FrontEnd/templates/index.html")
	if err != nil {
		utils.ErrorWriter(w, "Error 500", 500)
		return
	}

	if _, err := helpers.GetUserID(r); err == nil {
		http.Redirect(w, r, "/", http.StatusFound)
	}

	w.Write(Data)
}
