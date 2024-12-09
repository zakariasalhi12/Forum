package models

import (
	"database/sql"
	"errors"

	"forum/BackEnd/config"
	"forum/BackEnd/helpers"
)

var (
	ErrEmptyRequest           = errors.New("both email and password are required to log in")
	ErrInvalidPasswordOrEmail = errors.New("email or password incorrect")
)

type Login struct {
	ID       int
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *Login) LoginValidation() error {
	if helpers.CheckEmpty(user.Email, user.Password) {
		return ErrEmptyRequest
	}
	err := config.Config.Database.QueryRow("SELECT id FROM users WHERE email = ? AND password = ?", user.Email, user.Password).Scan(&user.ID)
	if err == sql.ErrNoRows {
		return ErrInvalidPasswordOrEmail
	}
	if err != nil {
		return err
	}
	return nil
}
