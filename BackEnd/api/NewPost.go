package api

import (
	"net/http"

	models "forum/BackEnd/Models"
	"forum/BackEnd/config"
	"forum/BackEnd/helpers"
)

func PostsAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.Writer(w, map[string]string{"Error": "Methode not allowed"}, http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var NewPost models.Posts

	Status, err := helpers.ParseRequestBody(r, &NewPost)
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
	NewPost.User_ID = int(Session.UserID)

	if err := NewPost.CheckPost(); err != nil {
		helpers.Writer(w, map[string]string{"Error": err.Error()}, http.StatusBadRequest)
		return
	}

	if _, err := NewPost.AddPost(); err != nil {
		config.Config.ServerLogGenerator(err.Error())
		helpers.Writer(w, map[string]string{"Error": helpers.ErrServer.Error()}, 500)
		return
	}

	helpers.Writer(w, map[string]string{"Message": "Post Created successfuly"}, 200)
}
