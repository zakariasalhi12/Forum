package api

import (
	"net/http"

	"forum/BackEnd/db"
	"forum/BackEnd/utils"

	"github.com/gofrs/uuid"
)

func RegisterAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ErrorWriter(w, "Methode not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	NewUser := utils.Register{
		UserName: r.FormValue("username"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
		Role:     "user",
	}
	if !utils.EmailChecker(NewUser.Email) {
		utils.Writer(w, map[string]string{"Error": "Bad Email Format"}, 400)
		return
	}
	if !utils.CheckEmpty(NewUser.Email, NewUser.Password, NewUser.Password) {
		utils.Writer(w, map[string]string{"Error": "Request Cant Be Empty"}, 400)
		return
	}
	_, err := db.Db.Exec("INSERT INTO users (username, email, password, role) VALUES (?, ?, ?, ?)", NewUser.UserName, NewUser.Email, NewUser.Password, NewUser.Role)
	if err != nil {
		utils.Writer(w, map[string]string{"Error": "Email Already Used"}, 400)
		return
	}
	newuuid, err := uuid.NewV4()
	if err != nil {
		utils.Writer(w, map[string]string{"Error": "An unexpected error occurred. Please try again later."}, 500)
		return
	}
	utils.Writer(w, map[string]string{"token": newuuid.String()}, 200)
}
