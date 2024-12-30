package models

import (
	"errors"
	"html"
	"net/http"
	"regexp"

	"forum/BackEnd/config"
	"forum/BackEnd/helpers"

	"github.com/gofrs/uuid"
)

type Register struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string
}

var (
	ErrInvalidPassword     = errors.New("password must be at least 6 characters long")
	ErrInvalidEmail        = errors.New("invalid Email")
	ErrInvalidUserName     = errors.New("invalid username")
	ErrEmailAlreadyUsed    = errors.New("email already used")
	ErrUserNameAlreadyUsed = errors.New("username already used")
)

func (R *Register) CheckUsername() (bool, error) {
	var Exists bool
	if err := config.Config.Database.QueryRow("SELECT COUNT(1) FROM users WHERE username = ? ", R.UserName).Scan(&Exists); err != nil {
		return false, err
	}
	return Exists, nil
}

func (R *Register) AddUserTodb(w http.ResponseWriter) error {
	R.Password = html.EscapeString(R.Password)
	if err := R.RegisterValidation(); err != nil {
		return err
	}
	Cheking ,  err := R.CheckUsername()
	if err != nil || !Cheking {
		return ErrUserNameAlreadyUsed
	}

	Res, err := config.Config.Database.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", R.UserName, R.Email, R.Password)
	if err != nil {
		return ErrEmailAlreadyUsed
	}
	newuuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	LastId, err := Res.LastInsertId()
	if err != nil {
		return err
	}
	NewSession := NewSession(w, newuuid.String(), LastId)
	if err := NewSession.CreateSession(); err != nil {
		return err
	}
	return nil
}

func (R *Register) RegisterValidation() error {
	R.Email, R.Password, R.UserName = helpers.RemoveExtraSpaces(R.Email), helpers.RemoveExtraSpaces(R.Password), helpers.RemoveExtraSpaces(R.UserName)
	// Check if any of the required fields (Email, Password, UserName) are empty
	if helpers.CheckEmpty(R.Email, R.Password, R.UserName) {
		return helpers.ErrInvalidRequest
	}
	if len(R.Email) > 40 || len(R.Password) > 30 || len(R.UserName) > 15 {
		return errors.New("content of email or password or username is large")
	}
	// The username must be between 3 and 20 characters and can contain letters, numbers, underscores, and hyphens
	UserNameValidation := regexp.MustCompile(`^[a-zA-Z0-9_-]{3,20}$`)
	// Validates email format: the username and domain can contain letters, numbers, and certain special characters, with a 2+ character top-level domain.
	EmailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !UserNameValidation.MatchString(R.UserName) {
		return ErrInvalidUserName
	}
	if !EmailRegex.MatchString(R.Email) {
		return ErrInvalidEmail
	}
	if len(R.Password) < 6 {
		return ErrInvalidPassword
	}

	return nil
}
