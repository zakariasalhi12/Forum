package models

import (
	"forum/BackEnd/config"
)

type LikesDislikes struct {
	PostOrCommentId int  `json:"postorcommentid"`
	IsComment       bool `json:"iscomment"`
	IsLike          bool `json:"islike"`
}

type TotalLikesAndDislikes struct {
	PostsLikes       int
	PostsDislikes    int
	CommentsLikes    int
	CommentsDislikes int
	AlreadyLiked     bool
}

func (l *LikesDislikes) AlreadyLiked(UserId int) bool {
	var exists bool
	err := config.Config.Database.QueryRow("SELECT COUNT(1) FROM likes_dislikes WHERE post_or_comment_id = ? AND user_id = ? AND is_like = ? AND is_comment = ?", l.PostOrCommentId, UserId, l.IsLike, l.IsComment).Scan(&exists)
	if err != nil {
		return exists
	}
	return exists
}

func (l *LikesDislikes) IsExistComment() bool {
	var exists bool
	if err := config.Config.Database.QueryRow("SELECT COUNT(1) FROM comments WHERE id = ?", l.PostOrCommentId).Scan(&exists); err != nil {
		config.Config.ServerLogGenerator(err.Error())
		return exists
	}
	return exists
}

func (l *LikesDislikes) IsExistPost() bool {
	var exists bool
	if err := config.Config.Database.QueryRow("SELECT COUNT(1) FROM posts WHERE id = ?", l.PostOrCommentId).Scan(&exists); err != nil {
		config.Config.ServerLogGenerator(err.Error())
		return false
	}
	return exists
}

func (l *LikesDislikes) DeleteLikeOrDislike(UserId int) error {
	_, err := config.Config.Database.Exec("DELETE FROM likes_dislikes WHERE post_or_comment_id = ? AND user_id = ? AND is_like = ? AND is_comment = ?", l.PostOrCommentId, UserId, l.IsLike, l.IsComment)
	if err != nil {
		config.Config.ServerLogGenerator(err.Error())
		return err
	}
	return nil
}

func (l *LikesDislikes) InsertLikeOrDislike(UserId int) error {
	_, err := config.Config.Database.Exec("INSERT INTO likes_dislikes (post_or_comment_id, user_id, is_like, is_comment) VALUES (?, ?, ?, ?)", l.PostOrCommentId, UserId, l.IsLike, l.IsComment)
	if err != nil {
		config.Config.ServerLogGenerator(err.Error())
		return err
	}
	return nil
}

func (t *TotalLikesAndDislikes) CountTotal(Id int) error {
	err := config.Config.Database.QueryRow("SELECT COUNT(*) FROM likes_dislikes WHERE is_like = TRUE AND is_comment = FALSE AND post_or_comment_id = ?", Id).Scan(&t.PostsLikes)
	if err != nil {
		config.Config.ServerLogGenerator(err.Error())
		return err
	}
	err = config.Config.Database.QueryRow("SELECT COUNT(*) FROM likes_dislikes WHERE is_like = FALSE AND is_comment = FALSE AND post_or_comment_id = ?", Id).Scan(&t.PostsDislikes)
	if err != nil {
		config.Config.ServerLogGenerator(err.Error())
		return err
	}
	err = config.Config.Database.QueryRow("SELECT COUNT(*) FROM likes_dislikes WHERE is_like = TRUE AND is_comment = TRUE AND post_or_comment_id = ?", Id).Scan(&t.CommentsLikes)
	if err != nil {
		config.Config.ServerLogGenerator(err.Error())
		return err
	}

	err = config.Config.Database.QueryRow("SELECT COUNT(*) FROM likes_dislikes WHERE is_like = FALSE AND is_comment = TRUE AND post_or_comment_id = ?", Id).Scan(&t.CommentsDislikes)
	if err != nil {
		config.Config.ServerLogGenerator(err.Error())
		return err
	}
	return nil
}
