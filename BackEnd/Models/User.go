package models

import (
	"database/sql"
	"errors"

	"forum/BackEnd/config"
)

var ErrInvalidUserID = errors.New("invalid userid")

type User struct {
	Id         int
	UserName   string
	CreatedAt  string
	Role       string
	TotalPosts int
	Email      string
}

func (User *User) CheckerUserExists() (bool, error) {
	var Exists bool

	if err := config.Config.Database.QueryRow("SELECT COUNT(1) FROM users WHERE id = ?", User.Id).Scan(&Exists); err != nil {
		return false, err
	}

	return Exists, nil
}

func (User *User) GetUserName() error {
	err := config.Config.Database.QueryRow("SELECT username FROM users WHERE id = ?", User.Id).Scan(&User.UserName)
	if err == sql.ErrNoRows {
		return ErrInvalidUserID
	}
	if err != nil {
		return err
	}
	return nil
}

func (User *User) GetUserEmail() error {
	err := config.Config.Database.QueryRow("SELECT email FROM users WHERE id = ?", User.Id).Scan(&User.Email)
	if err == sql.ErrNoRows {
		return ErrInvalidUserID
	}
	if err != nil {
		return err
	}
	return nil
}

func (User *User) GetRole() error {
	err := config.Config.Database.QueryRow("SELECT role FROM users WHERE id = ?", User.Id).Scan(&User.Role)
	if err == sql.ErrNoRows {
		return ErrInvalidUserID
	}
	if err != nil {
		return err
	}
	return nil
}

func (User *User) GetDate() error {
	err := config.Config.Database.QueryRow("SELECT created_at FROM users WHERE id = ?", User.Id).Scan(&User.CreatedAt)
	if err == sql.ErrNoRows {
		return ErrInvalidUserID
	}
	if err != nil {
		return err
	}
	return nil
}

func (User *User) GetTotalPosts() error {
	err := config.Config.Database.QueryRow("SELECT COUNT(*) FROM posts WHERE user_id = ?", User.Id).Scan(&User.TotalPosts)
	if err == sql.ErrNoRows {
		return ErrInvalidUserID
	}
	if err != nil {
		return err
	}
	return nil
}
