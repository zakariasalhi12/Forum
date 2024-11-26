package helpers

import (
	"database/sql"
	"errors"

	"forum/BackEnd/db"
)

func GetUserName(id int) (string, error) {
	var UserName string
	err := db.Db.QueryRow("SELECT username FROM users WHERE id = ?", id).Scan(&UserName)
	if err != nil {
		return "", err
	}
	if err == sql.ErrNoRows {
		return "", errors.New("invalid userid")
	}
	return UserName, nil
}
