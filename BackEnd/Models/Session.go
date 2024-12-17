package models

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"forum/BackEnd/config"
)

var (
	ErrLogout         = errors.New("failed to log out user")
	ErrInvalidToken   = errors.New("invalid token")
	ErrNotLogged      = errors.New("you are not logged")
	ErrSessionExpired = errors.New("session expired")
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
	Row, err := config.Config.Database.Exec("DELETE FROM sessions WHERE token = ?", s.Token)
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
	if _, err := config.Config.Database.Exec("INSERT INTO sessions (user_id, token) VALUES (?, ?)", s.UserID, s.Token); err != nil {
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
	result, err := config.Config.Database.Exec("UPDATE sessions SET token = ? WHERE user_id = ?", s.Token, s.UserID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		_, err := config.Config.Database.Exec("INSERT INTO sessions (user_id, token) VALUES (?, ?)", s.UserID, s.Token)
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

func (s *Session) GetUserID(r *http.Request) error {
	Token, err := r.Cookie("token")
	if err != nil {
		return ErrNotLogged
	}
	s.Token = Token.Value
	err = config.Config.Database.QueryRow("SELECT user_id FROM sessions WHERE token = ?", s.Token).Scan(&s.UserID)
	if err == sql.ErrNoRows {
		return ErrInvalidToken
	}
	if err != nil {
		return err
	}
	return nil
}

func (s *Session) DeleteExpiredSession() error {
	Row, err := config.Config.Database.Exec("DELETE FROM sessions WHERE user_id = ? AND DATETIME(created_at, '+24 hours') <= CURRENT_TIMESTAMP", s.UserID)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		config.Config.ServerLogGenerator(err.Error())
		return err
	}
	counter, err := Row.RowsAffected()
	if err != nil {
		config.Config.ServerLogGenerator(err.Error())
		return err
	}
	if counter == 1 {
		http.SetCookie(s.Response, &http.Cookie{
			Name:    "token",
			Value:   "",
			Path:    "/",
			Expires: time.Now(),
		})
		return ErrSessionExpired
	}

	return nil
}
