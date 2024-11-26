package helpers

import (
	"net/http"
	"time"

	"forum/BackEnd/db"
)

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
