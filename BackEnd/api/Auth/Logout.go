package auth

import (
	"fmt"
	"net/http"

	models "forum/BackEnd/Models"
	"forum/BackEnd/config"
	"forum/BackEnd/helpers"
)

func LogoutAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helpers.Writer(w, map[string]string{"Error": helpers.ErrMethod.Error()}, http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	cookie, err := r.Cookie("token")
	if err != nil {
		helpers.Writer(w, map[string]string{"Unauthorized": "Token missing or invalid"}, http.StatusUnauthorized)
		return
	}
	OldSession := &models.Session{Response: w, Token: cookie.Value}
	OldSession.GetUserID(r)

	if err := OldSession.DeleteSession(); err != nil {
		config.Config.ServerLogGenerator(err.Error())
		helpers.Writer(w, map[string]string{"Error": err.Error()}, 500)
		return
	}
	helpers.Writer(w, map[string]string{"Message": "Logout successfuly"}, 200)
	// Logs Part
	LoggoutUsser := models.User{Id: int(OldSession.UserID)}
	LoggoutUsser.GetUserName()
	LoggoutUsser.GetUserEmail()
	config.Config.ApiLogGenerator(fmt.Sprintf(`New Logout | UserName : "%s" , Email : "%s"`, LoggoutUsser.UserName, LoggoutUsser.Email))
}
