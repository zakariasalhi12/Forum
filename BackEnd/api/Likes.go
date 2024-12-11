package api

import (
	"net/http"

	models "forum/BackEnd/Models"
	"forum/BackEnd/config"
	"forum/BackEnd/helpers"
)

func AddLikeAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.Writer(w, map[string]string{"Error": "Methode not allowed"}, http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var NewLikeOrDislike models.LikesDislikes

	// Get the Body Request And Parse It into my newuser Model
	Status, err := helpers.ParseRequestBody(r, &NewLikeOrDislike)
	if err != nil {
		helpers.Writer(w, map[string]string{"Error": err.Error()}, Status)
		return
	}

	session := &models.Session{}
	if err := session.GetUserID(r); err != nil {
		helpers.Writer(w, map[string]string{"Error": err.Error()}, 500)
		return
	}

	var exists int
	if NewLikeOrDislike.IsComment {
		if err := config.Config.Database.QueryRow("SELECT COUNT(*) FROM comments WHERE id = ?", NewLikeOrDislike.PostOrCommentId).Scan(&exists); err != nil {
			helpers.Writer(w, map[string]string{"Error": err.Error()}, 500)
			return
		}
		if exists == 0 {
			helpers.Writer(w, map[string]string{"Error": "CommentId is not exist"}, 400)
			return
		}

	}
	if !NewLikeOrDislike.IsComment {
		if err := config.Config.Database.QueryRow("SELECT COUNT(*) FROM posts WHERE id = ?", NewLikeOrDislike.PostOrCommentId).Scan(&exists); err != nil {
			helpers.Writer(w, map[string]string{"Error": err.Error()}, 500)
			return
		}
		if exists == 0 {
			helpers.Writer(w, map[string]string{"Error": "PostID is not exist"}, 400)
			return
		}
	}

	IsLiked := NewLikeOrDislike.AlreadyLiked(int(session.UserID))
	CloneLike := models.LikesDislikes{
		IsComment:       NewLikeOrDislike.IsComment,
		IsLike:          !NewLikeOrDislike.IsLike,
		PostOrCommentId: NewLikeOrDislike.PostOrCommentId,
	}
	ReverseLike := CloneLike.AlreadyLiked(int(session.UserID))
	if ReverseLike {
		_, err = config.Config.Database.Exec("DELETE FROM likes_dislikes WHERE post_or_comment_id = ? AND user_id = ? AND is_like = ? AND is_comment = ?", CloneLike.PostOrCommentId, session.UserID, CloneLike.IsLike, CloneLike.IsComment)
		if err != nil {
			helpers.Writer(w, map[string]string{"Error": err.Error()}, 500)
			return
		}
	}
	if IsLiked {
		_, err = config.Config.Database.Exec("DELETE FROM likes_dislikes WHERE post_or_comment_id = ? AND user_id = ? AND is_like = ? AND is_comment = ?", NewLikeOrDislike.PostOrCommentId, session.UserID, NewLikeOrDislike.IsLike, NewLikeOrDislike.IsComment)
		if err != nil {
			helpers.Writer(w, map[string]string{"Error": err.Error()}, 500)
			return
		}
	} else {
		_, err = config.Config.Database.Exec("INSERT INTO likes_dislikes (post_or_comment_id, user_id, is_like, is_comment) VALUES (?, ?, ?, ?)", NewLikeOrDislike.PostOrCommentId, session.UserID, NewLikeOrDislike.IsLike, NewLikeOrDislike.IsComment)
		if err != nil {
			helpers.Writer(w, map[string]string{"Error": err.Error()}, 500)
			return
		}
	}

	var PostLikesCounter, PostDislikesCounter, CommentsLikeCounter, CommentsDislikesCounter int

	err = config.Config.Database.QueryRow("SELECT COUNT(*) FROM likes_dislikes WHERE is_like = TRUE AND is_comment = FALSE AND post_or_comment_id = ?", NewLikeOrDislike.PostOrCommentId).Scan(&PostLikesCounter)
	if err != nil {
		helpers.Writer(w, map[string]string{"Error": err.Error()}, 500)
		return
	}
	err = config.Config.Database.QueryRow("SELECT COUNT(*) FROM likes_dislikes WHERE is_like = FALSE AND is_comment = FALSE AND post_or_comment_id = ?", NewLikeOrDislike.PostOrCommentId).Scan(&PostDislikesCounter)
	if err != nil {
		helpers.Writer(w, map[string]string{"Error": err.Error()}, 500)
		return
	}
	err = config.Config.Database.QueryRow("SELECT COUNT(*) FROM likes_dislikes WHERE is_like = TRUE AND is_comment = TRUE AND post_or_comment_id = ?", NewLikeOrDislike.PostOrCommentId).Scan(&CommentsLikeCounter)
	if err != nil {
		helpers.Writer(w, map[string]string{"Error": err.Error()}, 500)
		return
	}
	err = config.Config.Database.QueryRow("SELECT COUNT(*) FROM likes_dislikes WHERE is_like = FALSE AND is_comment = TRUE AND post_or_comment_id = ?", NewLikeOrDislike.PostOrCommentId).Scan(&CommentsDislikesCounter)
	if err != nil {
		helpers.Writer(w, map[string]string{"Error": err.Error()}, 500)
		return
	}
	helpers.Writer(w, models.TotalLikesAndDislikes{
		PostsLikes:       PostLikesCounter,
		PostsDislikes:    PostDislikesCounter,
		CommentsLikes:    CommentsLikeCounter,
		CommentsDislikes: CommentsDislikesCounter,
		AlreadyLiked:     IsLiked,
	}, 200)
}
