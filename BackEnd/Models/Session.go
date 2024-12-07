package models

import (
	"net/http"
	"time"

	"forum/BackEnd/db"

	"github.com/gofrs/uuid"
)

type Session struct {
	Token    uuid.UUID
	UserID   int64
	Expires  time.Time
	Response http.ResponseWriter
	Path     string
}

func NewSession(w http.ResponseWriter, token uuid.UUID, UserId int64) *Session {
	return &Session{
		Token:    token,
		UserID:   UserId,
		Expires:  time.Now().Add(24 * time.Hour),
		Response: w,
		Path:     "/",
	}
}

func (s *Session) CreateSession() error {
	if _, err := db.Db.Exec("INSERT INTO sessions (user_id, token) VALUES (?, ?)", s.UserID, s.Token.String()); err != nil {
		return err
	}
	http.SetCookie(s.Response, &http.Cookie{
		Name:    "token",
		Value:   s.Token.String(),
		Path:    "/",
		Expires: s.Expires,
	})
	return nil
}
