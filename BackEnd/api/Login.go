package api

import (
	"encoding/json"
	"net/http"

	"forum/BackEnd/utils"
)

func LoginApi(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ErrorWriter(w, "Methode not allowed", http.StatusMethodNotAllowed)
		return
	}
	var user utils.Login

	user.Email = r.FormValue("email")
	user.Password = r.FormValue("password")

	Response, err := json.Marshal(user)
	if err != nil {
		utils.ErrorWriter(w, "Cannot marchal data", 500)
		return
	}
	w.WriteHeader(200)
	w.Write(Response)
}
