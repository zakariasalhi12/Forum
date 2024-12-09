package api

import (
	"net/http"
	"strconv"

	models "forum/BackEnd/Models"
	"forum/BackEnd/helpers"
)

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helpers.Writer(w, map[string]string{"Error": "Methode not allowed"}, http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	UserId := r.FormValue("id")
	if UserId == "" {
		helpers.Writer(w, map[string]string{"Error": "Id is importent"}, http.StatusBadRequest)
		return
	}

	Id, err := strconv.Atoi(UserId)
	if err != nil {
		helpers.Writer(w, map[string]string{"Error": "ID must be int"}, http.StatusBadRequest)
		return
	}

	UserInfo := &models.User{Id: Id}

	if err := UserInfo.GetDate(); err != nil {
		helpers.Writer(w, map[string]string{"Error": err.Error()}, http.StatusBadRequest)
		return
	}

	if err := UserInfo.GetRole(); err != nil {
		helpers.Writer(w, map[string]string{"Error": err.Error()}, http.StatusBadRequest)
		return
	}

	if err := UserInfo.GetTotalPosts(); err != nil {
		helpers.Writer(w, map[string]string{"Error": err.Error()}, http.StatusBadRequest)
		return
	}

	helpers.Writer(w, &UserInfo, 200)
}
