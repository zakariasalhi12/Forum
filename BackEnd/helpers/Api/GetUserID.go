package helpers

import (
	"database/sql"
	"errors"
	"net/http"

	"forum/BackEnd/db"
)

func GetUserID(r *http.Request) (int, error) {
	var userID int
	tokenCookie, err := r.Cookie("token")
	if err != nil {
		return -1, errors.New("you must be logged in to create a post")
	}
	err = db.Db.QueryRow("SELECT user_id FROM sessions WHERE token = ?", tokenCookie.Value).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return -1, errors.New("session not found or expired")
		}
		return -1, err
	}

	return userID, nil
}
