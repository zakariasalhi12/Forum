package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"forum/BackEnd/db"
	"forum/BackEnd/helpers"

	"github.com/gofrs/uuid"
)

func RegisterAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.Writer(w, map[string]string{"Error": "Methode not allowed"}, http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var NewUser helpers.Register
	if err := json.NewDecoder(r.Body).Decode(&NewUser); err != nil {
		helpers.Writer(w, map[string]string{"Error": "An unexpected error occurred. Please try again later."}, 500)
		return
	}
	NewUser.Role = "user"
	if !helpers.EmailChecker(NewUser.Email) {
		helpers.Writer(w, map[string]string{"Error": "Bad Email Format"}, 400)
		return
	}
	if helpers.CheckEmpty(NewUser.UserName, NewUser.Email, NewUser.Password) {
		helpers.Writer(w, map[string]string{"Error": "Request Cant Be Empty"}, 400)
		return
	}
	Res, err := db.Db.Exec("INSERT INTO users (username, email, password, role) VALUES (?, ?, ?, ?)", NewUser.UserName, NewUser.Email, NewUser.Password, NewUser.Role)
	if err != nil {
		helpers.Writer(w, map[string]string{"Error": "Email Already Used"}, 400)
		return
	}
	newuuid, err := uuid.NewV4()
	if err != nil {
		helpers.Writer(w, map[string]string{"Error": "An unexpected error occurred. Please try again later."}, 500)
		return
	}
	LastId, err := Res.LastInsertId()
	if err != nil {
		helpers.Writer(w, map[string]string{"Error": "An unexpected error occurred. Please try again later."}, 500)
		return
	}
	if err := helpers.SessionCreate(w, LastId, newuuid.String()); err != nil {
		helpers.Writer(w, map[string]string{"Error": "An unexpected error occurred. Please try again later."}, 500)
		return
	}
	helpers.Writer(w, map[string]string{"token": newuuid.String(), "UserId": strconv.Itoa(int(LastId))}, 200)
}