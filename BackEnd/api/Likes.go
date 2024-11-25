package api

import (
	"encoding/json"
	"net/http"

	"forum/BackEnd/db"
	"forum/BackEnd/utils"
)

func AddLikeAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.Writer(w, map[string]string{"Error": "Methode not allowed"}, http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var NewLikeOrDislike utils.LikesDislikes
	if err := json.NewDecoder(r.Body).Decode(&NewLikeOrDislike); err != nil {
		utils.Writer(w, map[string]string{"Error": "An unexpected error occurred. Please try again later."}, 500)
		return
	}
	UserID, err := GetUserID(r)
	if err != nil {
		utils.Writer(w, map[string]string{"Error": err.Error()}, 500)
		return
	}

	IsLiked, err := AlreadyLiked(UserID, NewLikeOrDislike)
	if err != nil {
		utils.Writer(w, map[string]string{"Error1": err.Error()}, 500)
		return
	}
	if IsLiked {
		_, err = db.Db.Exec("DELETE FROM likes_dislikes WHERE post_id = ? AND user_id = ? AND is_like = ? AND is_comment = ?", NewLikeOrDislike.PostId, UserID, NewLikeOrDislike.IsLike, NewLikeOrDislike.IsComment)
		if err != nil {
			utils.Writer(w, map[string]string{"Error": err.Error()}, 500)
			return
		}
	} else {
		_, err = db.Db.Exec("INSERT INTO likes_dislikes (post_id, user_id, is_like, is_comment) VALUES (?, ?, ?, ?)", NewLikeOrDislike.PostId, UserID, NewLikeOrDislike.IsLike, NewLikeOrDislike.IsComment)
		if err != nil {
			utils.Writer(w, map[string]string{"Error": err.Error()}, 500)
			return
		}
	}

	var PostLikesCounter, PostDislikesCounter, CommentsLikeCounter, CommentsDislikesCounter int

	err = db.Db.QueryRow("SELECT COUNT(*) FROM likes_dislikes WHERE is_like = TRUE AND is_comment = FALSE").Scan(&PostLikesCounter)
	if err != nil {
		utils.Writer(w, map[string]string{"Error": err.Error()}, 500)
		return
	}
	err = db.Db.QueryRow("SELECT COUNT(*) FROM likes_dislikes WHERE is_like = FALSE AND is_comment = FALSE").Scan(&PostDislikesCounter)
	if err != nil {
		utils.Writer(w, map[string]string{"Error": err.Error()}, 500)
		return
	}
	err = db.Db.QueryRow("SELECT COUNT(*) FROM likes_dislikes WHERE is_like = TRUE AND is_comment = TRUE").Scan(&CommentsLikeCounter)
	if err != nil {
		utils.Writer(w, map[string]string{"Error": err.Error()}, 500)
		return
	}
	err = db.Db.QueryRow("SELECT COUNT(*) FROM likes_dislikes WHERE is_like = FALSE AND is_comment = TRUE").Scan(&CommentsDislikesCounter)
	if err != nil {
		utils.Writer(w, map[string]string{"Error": err.Error()}, 500)
		return
	}
	utils.Writer(w, map[string]int{
		"PostsLikes":       PostLikesCounter,
		"PostsDislikes":    PostDislikesCounter,
		"CommentsLikes":    CommentsLikeCounter,
		"CommentsDislikes": CommentsDislikesCounter,
	}, 200)
}

func AlreadyLiked(userid int, LikesAndDislikes utils.LikesDislikes) (bool, error) {
	var exists int
	row := db.Db.QueryRow("SELECT COUNT(*) FROM likes_dislikes WHERE post_id = ? AND user_id = ? AND is_like = ? AND is_comment = ?", LikesAndDislikes.PostId, userid, LikesAndDislikes.IsLike, LikesAndDislikes.IsComment)
	err := row.Scan(&exists)
	if exists == 0 {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}
