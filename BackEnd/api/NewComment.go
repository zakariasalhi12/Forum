package api

import (
	"net/http"

	"forum/BackEnd/db"
	"forum/BackEnd/utils"
)

func NewCommentAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.Writer(w, map[string]string{"Error": "Methode not allowed"}, http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	PostId := r.FormValue("postid")
	PostContent := r.FormValue("content")
	UserId, err := GetUserID(r)
	if err != nil {
		utils.Writer(w, map[string]string{"Error": err.Error()}, http.StatusBadRequest)
		return
	}
	_, err = db.Db.Exec("INSERT INTO comments (post_id, user_id, content) VALUES (?, ?, ?)", PostId, UserId, PostContent)
	if err != nil {
		utils.Writer(w, map[string]string{"Error": err.Error()}, http.StatusInternalServerError)
		return
	}

	utils.Writer(w, map[string]string{"Message": "Comment Created successfuly"}, 200)
}
