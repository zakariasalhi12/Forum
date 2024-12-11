package models

import "forum/BackEnd/config"

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

func (l *LikesDislikes) AlreadyLiked(Userid int) bool {
	var exists int
	err := config.Config.Database.QueryRow("SELECT COUNT(*) FROM likes_dislikes WHERE post_or_comment_id = ? AND user_id = ? AND is_like = ? AND is_comment = ?", l.PostOrCommentId, Userid, l.IsLike, l.IsComment).Scan(&exists)
	if exists == 0 || err != nil {
		return false
	}
	return true
}
