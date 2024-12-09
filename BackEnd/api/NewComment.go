package api

import (
	"encoding/json"
	"html"
	"io"
	"net/http"

	models "forum/BackEnd/Models"
	"forum/BackEnd/db"
	"forum/BackEnd/helpers"
)

func NewCommentAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.Writer(w, map[string]string{"Error": "Methode not allowed"}, http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var Comment helpers.Comment
	Response, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.Writer(w, map[string]string{"Error": "An unexpected error occurred. Please try again later."}, 500)
		return
	}
	if err := json.Unmarshal(Response, &Comment); err != nil {
		helpers.Writer(w, map[string]string{"Error": "Invalid Request"}, 400)
		return
	}
	Comment.Content = html.EscapeString(Comment.Content)
	Session := &models.Session{}
	if err := Session.GetUserID(r); err != nil {
		helpers.Writer(w, map[string]string{"Error": err.Error()}, http.StatusBadRequest)
		return
	}
	_, err = db.Db.Exec("INSERT INTO comments (post_id, user_id, content) VALUES (?, ?, ?)", Comment.PostId, Session.UserID, Comment.Content)
	if err != nil {
		helpers.Writer(w, map[string]string{"Error": err.Error()}, http.StatusInternalServerError)
		return
	}

	helpers.Writer(w, map[string]string{"Message": "Comment Created successfuly"}, 200)
}
