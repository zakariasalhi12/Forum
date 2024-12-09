package auth

import (
	"net/http"
	"strconv"

	"forum/BackEnd/helpers"
)

func Islogged(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helpers.Writer(w, map[string]string{"Error": helpers.ErrMethod.Error()}, http.StatusMethodNotAllowed)
		return
	}
	id, err := helpers.GetUserID(r)
	if err != nil {
		helpers.Writer(w, map[string]string{"Error": err.Error()}, 400)
		return
	}
	UserName, err := helpers.GetUserName(id)
	if err != nil {
		helpers.Writer(w, map[string]string{"Error": err.Error()}, 400)
		return
	}

	helpers.Writer(w, map[string]string{"username": UserName, "user_id": strconv.Itoa(id)}, 200)
}
