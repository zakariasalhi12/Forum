package models

import (
	"errors"
	"net/http"
	"regexp"

	"forum/BackEnd/db"

	"github.com/gofrs/uuid"
)

type Register struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string
}

var (
	InvalidPassword  = errors.New("password must be at least 6 characters long")
	InvalidEmail     = errors.New("invalid Email")
	InvalidUserName  = errors.New("invalid username")
	EmailAlreadyUsed = errors.New("email already used")
)

func NewUser() *Register {
	return &Register{
		Role: "User",
	}
}

func (R *Register) AddUserTodb(w http.ResponseWriter) error {
	if err := R.RegisterValidation(); err != nil {
		return err
	}
	Res, err := db.Db.Exec("INSERT INTO users (username, email, password, role) VALUES (?, ?, ?, ?)", R.UserName, R.Email, R.Password, R.Role)
	if err != nil {
		return EmailAlreadyUsed
	}
	newuuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	LastId, err := Res.LastInsertId()
	if err != nil {
		return err
	}
	NewSession := NewSession(w, newuuid, LastId)
	if err := NewSession.CreateSession(); err != nil {
		return err
	}
	return nil
}

func (R *Register) RegisterValidation() error {
	// The username must be between 3 and 20 characters and can contain letters, numbers, underscores, and hyphens
	UserNameValidation := regexp.MustCompile(`^[a-zA-Z0-9_-]{3,20}$`)
	// Validates email format: the username and domain can contain letters, numbers, and certain special characters, with a 2+ character top-level domain.
	EmailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !UserNameValidation.MatchString(R.UserName) {
		return InvalidUserName
	}
	if !EmailRegex.MatchString(R.Email) {
		return InvalidEmail
	}
	if len(R.Password) < 6 {
		return InvalidPassword
	}

	return nil
}
