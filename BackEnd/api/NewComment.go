package api

import (
	"html"
	"net/http"

	models "forum/BackEnd/Models"
	"forum/BackEnd/config"
	"forum/BackEnd/helpers"
)

func NewCommentAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.Writer(w, map[string]string{"Error": "Methode not allowed"}, http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var Comment models.Comment
	Status, err := helpers.ParseRequestBody(r, &Comment)
	if err != nil {
		helpers.Writer(w, map[string]string{"Error": err.Error()}, Status)
		return
	}
	Comment.Content = html.EscapeString(Comment.Content)
	Session := &models.Session{}
	if err := Session.GetUserID(r); err != nil {
		helpers.Writer(w, map[string]string{"Error": err.Error()}, http.StatusBadRequest)
		return
	}
	_, err = config.Config.Database.Exec("INSERT INTO comments (post_id, user_id, content) VALUES (?, ?, ?)", Comment.PostId, Session.UserID, Comment.Content)
	if err != nil {
		helpers.Writer(w, map[string]string{"Error": err.Error()}, http.StatusInternalServerError)
		return
	}

	helpers.Writer(w, map[string]string{"Message": "Comment Created successfuly"}, 200)
}
