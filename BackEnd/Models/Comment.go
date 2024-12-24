package models

import (
	"errors"
	"html"

	"forum/BackEnd/config"
)

type Comment struct {
	UserID  int
	PostId  int    `json:"postid"`
	Content string `json:"content"`
}

func CheckCommentExist(CommentId int) (bool, error) {
	var Exists bool
	if err := config.Config.Database.QueryRow("SELECT COUNT(1) FROM comments WHERE id = ? ", CommentId).Scan(&Exists); err != nil {
		return false, err
	}
	return Exists, nil
}

func (c *Comment) CheckCommentValidation() error {
	if len(c.Content) >= 250 {
		return errors.New("the maximum comment content length is 250 characters")
	}
	return nil
}

func (c *Comment) AddComment() error {
	html.EscapeString(c.Content)
	_, err := config.Config.Database.Exec("INSERT INTO comments (post_id, user_id, content) VALUES (?, ?, ?)", c.PostId, c.UserID, c.Content)
	if err != nil {
		return err
	}
	return nil
}
