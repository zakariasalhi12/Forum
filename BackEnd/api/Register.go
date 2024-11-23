package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"forum/BackEnd/db"
	"forum/BackEnd/utils"

	"github.com/gofrs/uuid"
)

func RegisterAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.Writer(w, map[string]string{"Error": "Methode not allowed"}, http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var NewUser utils.Register
	if err := json.NewDecoder(r.Body).Decode(&NewUser); err != nil {
		utils.Writer(w, map[string]string{"Error": "An unexpected error occurred. Please try again later."}, 500)
		return
	}
	NewUser.Role = "user"
	if !utils.EmailChecker(NewUser.Email) {
		utils.Writer(w, map[string]string{"Error": "Bad Email Format"}, 400)
		return
	}
	if utils.CheckEmpty(NewUser.UserName, NewUser.Email, NewUser.Password) {
		utils.Writer(w, map[string]string{"Error": "Request Cant Be Empty"}, 400)
		return
	}
	Res, err := db.Db.Exec("INSERT INTO users (username, email, password, role) VALUES (?, ?, ?, ?)", NewUser.UserName, NewUser.Email, NewUser.Password, NewUser.Role)
	if err != nil {
		utils.Writer(w, map[string]string{"Error": "Email Already Used"}, 400)
		return
	}
	newuuid, err := uuid.NewV4()
	if err != nil {
		utils.Writer(w, map[string]string{"Error": "An unexpected error occurred. Please try again later."}, 500)
		return
	}
	LastId, err := Res.LastInsertId()
	if err != nil {
		utils.Writer(w, map[string]string{"Error": "An unexpected error occurred. Please try again later."}, 500)
		return
	}
	if err := SessionCreate(w, LastId, newuuid.String()); err != nil {
		utils.Writer(w, map[string]string{"Error": "An unexpected error occurred. Please try again later."}, 500)
		return
	}
	utils.Writer(w, map[string]string{"token": newuuid.String(), "UserId": strconv.Itoa(int(LastId))}, 200)
}

func SessionCreate(w http.ResponseWriter, userID int64, token string) error {
	if _, err := db.Db.Exec("INSERT INTO sessions (user_id, token) VALUES (?, ?)", userID, token); err != nil {
		return err
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Path:    "/",
		Expires: time.Now().Add(24 * time.Hour),
	})
	return nil
}
