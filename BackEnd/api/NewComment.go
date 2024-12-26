package api

import (
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
		if Status == 500 {
			config.Config.ServerLogGenerator(err.Error())
		}
		helpers.Writer(w, map[string]string{"Error": err.Error()}, Status)
		return
	}

	Session := &models.Session{}
	if err := Session.GetUserID(r); err != nil {
		helpers.Writer(w, map[string]string{"Error": err.Error()}, http.StatusBadRequest)
		return
	}

	Comment.UserID = int(Session.UserID)
	if err := Comment.CheckCommentValidation(); err != nil {
		helpers.Writer(w, map[string]string{"Error": err.Error()}, http.StatusBadRequest)
		return
	}
	if err := Comment.AddComment(); err != nil {
		config.Config.ServerLogGenerator(err.Error())
		helpers.Writer(w, map[string]string{"Error": helpers.ErrServer.Error()}, 500)
		return
	}

	helpers.Writer(w, map[string]string{"Message": "Comment Created successfuly"}, 200)
}
