package auth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"forum/BackEnd/db"
	helpers "forum/BackEnd/helpers/Api_Helper"

	"github.com/gofrs/uuid"
)

func LoginApi(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.Writer(w, map[string]string{"Error": "Methode not allowed"}, http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var NewUser helpers.Login

	Response, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.Writer(w, map[string]string{"Error": "An unexpected error occurred. Please try again later."}, 500)
		return
	}
	if err := json.Unmarshal(Response, &NewUser); err != nil {
		helpers.Writer(w, map[string]string{"Error": "Invalid Request"}, 400)
		return
	}

	if helpers.CheckEmpty(NewUser.Email, NewUser.Password) {
		helpers.Writer(w, map[string]string{"Error": "Request Cant be empty"}, 400)
		return
	}
	var UserId int
	err = db.Db.QueryRow("SELECT id FROM users WHERE email = ? AND password = ?", NewUser.Email, NewUser.Password).Scan(&UserId)
	if err == sql.ErrNoRows {
		helpers.Writer(w, map[string]string{"Error": "Email or Password Incorrect"}, 400)
		return
	}
	if err != nil {
		helpers.Writer(w, map[string]string{"Error": err.Error()}, 500)
		return
	}
	uuid, err := uuid.NewV4()
	if err != nil {
		helpers.Writer(w, map[string]string{"Error": "An unexpected error occurred. Please try again later."}, 500)
		return
	}
	if err = UpdateSessionForUser(w, UserId, uuid.String()); err != nil {
		helpers.Writer(w, map[string]string{"Error": "An unexpected error occurred. Please try again later."}, 500)
		return
	}
	helpers.Writer(w, map[string]string{"token": uuid.String(), "userid": strconv.Itoa(UserId)}, 200)
}

func UpdateSessionForUser(w http.ResponseWriter, userID int, token string) error {
	result, err := db.Db.Exec("UPDATE sessions SET token = ? WHERE user_id = ?", token, userID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		_, err := db.Db.Exec("INSERT INTO sessions (user_id, token) VALUES (?, ?)", userID, token)
		if err != nil {
			return fmt.Errorf("error inserting session: %w", err)
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
