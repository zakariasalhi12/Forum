package api

import (
	"net/http"
	"strconv"

	"forum/BackEnd/db"
	"forum/BackEnd/utils"
)

func AddLikeAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.Writer(w, map[string]string{"Error": "Methode not allowed"}, http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	UserID, err := GetUserID(r)
	if err != nil {
		utils.Writer(w, map[string]string{"Error": err.Error()}, 500)
		return
	}
	PostID := r.FormValue("PostID")
	IsComment := r.FormValue("IsComment")
	IsLike := r.FormValue("IsLike")

	_, err = db.Db.Exec("INSERT INTO likes_dislikes (post_id, user_id, is_like, is_comment) VALUES (?, ?, ?, ?)", PostID, UserID, IsLike, IsComment)
	if err != nil {
		utils.Writer(w, map[string]string{"Error": err.Error()}, 500)
		return
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
	utils.Writer(w, map[string]string{
		"PostsLikes":       strconv.Itoa(PostLikesCounter),
		"PostsDislikes":    strconv.Itoa(PostDislikesCounter),
		"CommentsLikes":    strconv.Itoa(CommentsLikeCounter),
		"CommentsDislikes": strconv.Itoa(CommentsDislikesCounter),
	}, 200)
}
