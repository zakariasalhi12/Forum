package models

import (
	"database/sql"
	"errors"

	"forum/BackEnd/db"
)

var ErrInvalidUserID = errors.New("invalid userid")

type User struct {
	Id         int
	UserName   string
	CreatedAt  string
	Role       string
	TotalPosts int
}

func (User *User) GetUserName() error {
	err := db.Db.QueryRow("SELECT username FROM users WHERE id = ?", User.Id).Scan(&User.UserName)
	if err == sql.ErrNoRows {
		return ErrInvalidUserID
	}
	if err != nil {
		return err
	}
	return nil
}

func (User *User) GetRole() error {
	err := db.Db.QueryRow("SELECT role FROM users WHERE id = ?", User.Id).Scan(&User.Role)
	if err == sql.ErrNoRows {
		return ErrInvalidUserID
	}
	if err != nil {
		return err
	}
	return nil
}

func (User *User) GetDate() error {
	err := db.Db.QueryRow("SELECT created_at FROM users WHERE id = ?", User.Id).Scan(&User.CreatedAt)
	if err == sql.ErrNoRows {
		return ErrInvalidUserID
	}
	if err != nil {
		return err
	}
	return nil
}

func (User *User) GetTotalPosts() error {
	err := db.Db.QueryRow("SELECT COUNT(*) FROM posts WHERE user_id = ?", User.Id).Scan(&User.TotalPosts)
	if err == sql.ErrNoRows {
		return ErrInvalidUserID
	}
	if err != nil {
		return err
	}
	return nil
}
