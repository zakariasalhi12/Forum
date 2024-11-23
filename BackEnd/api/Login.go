package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"forum/BackEnd/db"
	"forum/BackEnd/utils"

	"github.com/gofrs/uuid"
)

func LoginApi(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.Writer(w, map[string]string{"Error": "Methode not allowed"}, http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var NewUser utils.Login
	if err := json.NewDecoder(r.Body).Decode(&NewUser); err != nil {
		utils.Writer(w, map[string]string{"Error": "Invalid Request"}, 400)
		return
	}
	if utils.CheckEmpty(NewUser.Email, NewUser.Password) {
		utils.Writer(w, map[string]string{"Error": "Request Cant be empty"}, 400)
		return
	}
	var UserId int
	err := db.Db.QueryRow("SELECT id FROM users WHERE email = ? AND password = ?", NewUser.Email, NewUser.Password).Scan(&UserId)
	if err == sql.ErrNoRows {
		utils.Writer(w, map[string]string{"Error": "Email or Password Incorrect"}, 400)
		return
	}
	if err != nil {
		utils.Writer(w, map[string]string{"Error": err.Error()}, 500)
		return
	}
	uuid, err := uuid.NewV4()
	if err != nil {
		utils.Writer(w, map[string]string{"Error": "An unexpected error occurred. Please try again later."}, 500)
	}
	if err = UpdateSessionForUser(w, UserId, uuid.String()); err != nil {
		utils.Writer(w, map[string]string{"Error": "An unexpected error occurred. Please try again later."}, 500)
	}
	utils.Writer(w, map[string]string{"token": uuid.String(), "userid": strconv.Itoa(UserId)}, 200)
}

func UpdateSessionForUser(w http.ResponseWriter, UserId int, token string) error {
	if _, err := db.Db.Exec("UPDATE sessions SET token = ? WHERE user_id = ?", token, UserId); err != nil {
		if err == sql.ErrNoRows {
			if _, err := db.Db.Exec("INSERT INTO sessions (user_id , token) VALUES (? , ?)", UserId, token); err != nil {
				return err
			}
			return err
		}
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Path:    "/",
		Expires: time.Now().Add(24 * time.Hour),
	})
	return nil
}
