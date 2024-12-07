package api

import (
	"net/http"

	"forum/BackEnd/db"
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
	var UserInfo helpers.UserInfo

	if err := db.Db.QueryRow("SELECT created_at, role FROM users WHERE id = ? ", UserId).Scan(&UserInfo.CreateDate, &UserInfo.Role); err != nil {
		helpers.Writer(w, map[string]string{"Error": "User not exist"}, http.StatusBadRequest)
		return
	}

	if err := db.Db.QueryRow("SELECT COUNT(*) FROM posts WHERE user_id = ?", UserId).Scan(&UserInfo.TotalPosts); err != nil {
		helpers.Writer(w, map[string]string{"Error": err.Error()}, http.StatusBadRequest)
		return
	}

	helpers.Writer(w, &UserInfo, 200)
}
