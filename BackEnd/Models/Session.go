package models

import (
	"errors"
	"net/http"
	"time"

	"forum/BackEnd/db"
)

var (
	ErrLogout       = errors.New("failed to log out user")
	ErrInvalidToken = errors.New("invalid token")
)

type Session struct {
	Token    string
	UserID   int64
	Expires  time.Time
	Response http.ResponseWriter
	Path     string
}

/*
The user ID is not needed when deleting a session.
The user ID can be retrieved directly from the token.
You can use the getId() method to retrieve the user ID.
*/

func NewSession(w http.ResponseWriter, token string, UserId int64) *Session {
	return &Session{
		Token:    token,
		UserID:   UserId,
		Expires:  time.Now().Add(24 * time.Hour),
		Response: w,
		Path:     "/",
	}
}

func (s *Session) DeleteSession() error {
	Row, err := db.Db.Exec("DELETE FROM sessions WHERE token = ?", s.Token)
	if err != nil {
		return ErrLogout
	}
	Counter, err := Row.RowsAffected()
	if err != nil {
		return ErrLogout
	}

	if Counter == 0 {
		return ErrLogout
	}
	http.SetCookie(s.Response, &http.Cookie{
		Name:    "token",
		Value:   "",
		Path:    s.Path,
		Expires: time.Now(),
	})

	return nil
}

func (s *Session) CreateSession() error {
	if _, err := db.Db.Exec("INSERT INTO sessions (user_id, token) VALUES (?, ?)", s.UserID, s.Token); err != nil {
		return err
	}
	http.SetCookie(s.Response, &http.Cookie{
		Name:    "token",
		Value:   s.Token,
		Path:    "/",
		Expires: s.Expires,
	})
	return nil
}

func (s *Session) UpdateSessionForUser() error {
	result, err := db.Db.Exec("UPDATE sessions SET token = ? WHERE user_id = ?", s.Token, s.UserID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		_, err := db.Db.Exec("INSERT INTO sessions (user_id, token) VALUES (?, ?)", s.UserID, s.Token)
		if err != nil {
			return err
		}
	}
	http.SetCookie(s.Response, &http.Cookie{
		Name:    "token",
		Value:   s.Token,
		Path:    s.Path,
		Expires: s.Expires,
	})

	return nil
}
