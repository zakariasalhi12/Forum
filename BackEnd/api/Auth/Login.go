package auth

import (
	"fmt"
	"net/http"

	models "forum/BackEnd/Models"
	"forum/BackEnd/config"
	"forum/BackEnd/helpers"

	"github.com/gofrs/uuid"
)

func LoginApi(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.Writer(w, map[string]string{"Error": helpers.ErrMethod.Error()}, http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var User models.Login

	Status, err := helpers.ParseRequestBody(r, &User)
	if err != nil {
		helpers.Writer(w, map[string]string{"Error": err.Error()}, Status)
		return
	}

	if helpers.CheckEmpty(User.Email, User.Password) {
		helpers.Writer(w, map[string]string{"Error": "Request Cant be empty"}, 400)
		return
	}
	err = User.LoginValidation()
	if err == models.ErrEmptyRequest || err == models.ErrInvalidPasswordOrEmail {
		helpers.Writer(w, map[string]string{"Error": err.Error()}, 400)
		return
	}
	if err != nil {
		helpers.Writer(w, map[string]string{"Error": helpers.ErrServer.Error()}, 500)
		return
	}
	uuid, err := uuid.NewV4()
	if err != nil {
		helpers.Writer(w, map[string]string{"Error": helpers.ErrServer.Error()}, 500)
		return
	}
	NewSession := models.NewSession(w, uuid.String(), int64(User.ID))
	if err := NewSession.UpdateSessionForUser(); err != nil {
		helpers.Writer(w, map[string]string{"Error": "An unexpected error occurred. Please try again later."}, 500)
		return
	}

	helpers.Writer(w, map[string]string{"Message": "Logged successful!"}, 200)

	// Logs Part
	LoggedUsser := models.User{Id: int(User.ID)}
	LoggedUsser.GetUserName()
	LoggedUsser.GetUserEmail()
	config.Config.ApiLogGenerator(fmt.Sprintf(`New Login | UserName : "%s" , Email : "%s"`, LoggedUsser.UserName, LoggedUsser.Email))
}
