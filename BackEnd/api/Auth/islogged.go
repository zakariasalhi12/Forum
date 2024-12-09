package auth

import (
	"net/http"
	"strconv"

	models "forum/BackEnd/Models"
	"forum/BackEnd/helpers"
)

func Islogged(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helpers.Writer(w, map[string]string{"Error": helpers.ErrMethod.Error()}, http.StatusMethodNotAllowed)
		return
	}
	Session := &models.Session{}
	if err := Session.GetUserID(r); err != nil {
		helpers.Writer(w, map[string]string{"Error": err.Error()}, 400)
		return
	}
	User := &models.User{Id: int(Session.UserID)}
	if err := User.GetUserName(); err != nil {
		helpers.Writer(w, map[string]string{"Error": err.Error()}, 400)
		return
	}

	helpers.Writer(w, map[string]string{"username": User.UserName, "user_id": strconv.Itoa(int(Session.UserID))}, 200)
}
