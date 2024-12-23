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
		if Status == 500 {
			config.Config.ServerLogGenerator(err.Error())
		}
		helpers.Writer(w, map[string]string{"Error": err.Error()}, Status)
		return
	}

	session := &models.Session{}
	if err := session.GetUserID(r); err != nil {
		config.Config.ServerLogGenerator(err.Error())
		helpers.Writer(w, map[string]string{"Error": err.Error()}, 500)
		return
	}
	if NewLikeOrDislike.IsComment && !NewLikeOrDislike.IsExistComment() {
		helpers.Writer(w, map[string]string{"Error": "CommentId is not exist"}, 400)
		return
	}
	if !NewLikeOrDislike.IsComment && !NewLikeOrDislike.IsExistPost() {
		helpers.Writer(w, map[string]string{"Error": "PostID is not exist"}, 400)
		return
	}

	IsLiked := NewLikeOrDislike.AlreadyLiked(int(session.UserID))
	CloneLike := models.LikesDislikes{
		IsComment:       NewLikeOrDislike.IsComment,
		IsLike:          !NewLikeOrDislike.IsLike,
		PostOrCommentId: NewLikeOrDislike.PostOrCommentId,
	}
	ReverseLike := CloneLike.AlreadyLiked(int(session.UserID))
	if ReverseLike {
		if err := CloneLike.DeleteLikeOrDislike(int(session.UserID)); err != nil {
			helpers.Writer(w, map[string]string{"Error": err.Error()}, 500)
			return
		}
	}
	if IsLiked {
		if err := NewLikeOrDislike.DeleteLikeOrDislike(int(session.UserID)); err != nil {
			helpers.Writer(w, map[string]string{"Error": err.Error()}, 500)
			return
		}
	} else {
		if err := NewLikeOrDislike.InsertLikeOrDislike(int(session.UserID)); err != nil {
			helpers.Writer(w, map[string]string{"Error": err.Error()}, 500)
			return
		}
	}

	Total := &models.TotalLikesAndDislikes{AlreadyLiked: IsLiked}
	if err := Total.CountTotal(NewLikeOrDislike.PostOrCommentId); err != nil {
		helpers.Writer(w, map[string]string{"Error": err.Error()}, 500)
		return
	}
	helpers.Writer(w, Total, 200)
}
