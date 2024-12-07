package models

import (
	"errors"
	"html"
	"net/http"
	"regexp"

	"forum/BackEnd/db"
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
	ErrInvalidPassword  = errors.New("password must be at least 6 characters long")
	ErrInvalidEmail     = errors.New("invalid Email")
	ErrInvalidUserName  = errors.New("invalid username")
	ErrEmailAlreadyUsed = errors.New("email already used")
)

func NewUser() *Register {
	return &Register{
		Role: "User",
	}
}

func (R *Register) AddUserTodb(w http.ResponseWriter) error {
	R.Password = html.EscapeString(R.Password)
	if err := R.RegisterValidation(); err != nil {
		return err
	}
	Res, err := db.Db.Exec("INSERT INTO users (username, email, password, role) VALUES (?, ?, ?, ?)", R.UserName, R.Email, R.Password, R.Role)
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
	// Check if any of the required fields (Email, Password, UserName) are empty
	if helpers.CheckEmpty(R.Email, R.Password, R.UserName) {
		return helpers.ErrInvalidRequest
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
