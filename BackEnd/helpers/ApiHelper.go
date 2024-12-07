package helpers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"time"

	"forum/BackEnd/db"
)

const (
	MethodError = "Method Not Allowed."
	ServerError = "An unexpected error occurred. Please try again later."
)

func Mapper(str1, str2 string) map[string]string {
	return map[string]string{str1: str2}
}

// write a json response from the given data
func Writer(w http.ResponseWriter, response any, status int) {
	res, err := json.Marshal(response)
	if err != nil {
		Writer(w, map[string]string{"Error": "An unexpected error occurred. Please try again later."}, 500)
		return
	}
	w.WriteHeader(status)
	w.Write(res)
}

// Create a session using uuid
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

// this function get an id and return the username of the user
func GetUserName(id int) (string, error) {
	var UserName string
	err := db.Db.QueryRow("SELECT username FROM users WHERE id = ?", id).Scan(&UserName)
	if err == sql.ErrNoRows {
		return "", errors.New("invalid userid")
	}
	if err != nil {
		return "", err
	}
	return UserName, nil
}

// get the user id from the token in cookies

func GetUserID(r *http.Request) (int, error) {
	var userID int
	tokenCookie, err := r.Cookie("token")
	if err != nil {
		return -1, errors.New("you are not logged")
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

// Email validation
func EmailChecker(email string) bool {
	EmailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return EmailRegex.MatchString(email)
}

func PasswordChecker(password string) bool {
	return len(password) > 5
}

// Machi wkita daba
func TagsChecker(arr []string) bool {
	return false
}

// Machi wkita daba
func UserNameChecker(s string) bool {
	return false
}

func CheckEmpty(args ...string) bool {
	for _, arg := range args {
		if arg == "" {
			return true
		}
	}
	return false
}
