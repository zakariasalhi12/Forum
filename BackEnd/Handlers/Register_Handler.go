package handlers

import (
	"net/http"
	"os"

	"forum/BackEnd/helpers"
)

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helpers.ErrorWriter(w, "Error 400", http.StatusMethodNotAllowed)
		return
	}

	Data, err := os.ReadFile("FrontEnd/templates/register.html")
	if err != nil {
		helpers.ErrorWriter(w, "Error 500", 500)
		return
	}

	if _, err := helpers.GetUserID(r); err == nil {
		http.Redirect(w, r, "/", http.StatusFound)
	}

	w.Write(Data)
}
