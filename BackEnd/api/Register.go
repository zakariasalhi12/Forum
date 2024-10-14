package api

import (
	"encoding/json"
	"net/http"

	"forum/BackEnd/utils"
)

func RegisterAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ErrorWriter(w, "Methode not allowed", http.StatusMethodNotAllowed)
		return
	}
	var NewUser utils.Register

	// Need to check the data after send it to response and hash the password

	NewUser.UserName = r.FormValue("username")
	NewUser.Email = r.FormValue("email")
	NewUser.Password = r.FormValue("password")

	Response, err := json.Marshal(NewUser)
	if err != nil {
		utils.ErrorWriter(w, "Cannot marchal data", 500)
		return
	}

	w.WriteHeader(200)
	w.Write(Response)
}
