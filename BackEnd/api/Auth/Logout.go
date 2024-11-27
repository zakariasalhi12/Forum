package auth

import (
	"net/http"
	"time"

	"forum/BackEnd/db"
	helpers "forum/BackEnd/helpers/Api_Helper"
)

func LogoutAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helpers.Writer(w, map[string]string{"Error": "Methode not allowed"}, http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	cookie, err := r.Cookie("token")
	if err != nil {
		helpers.Writer(w, "Unauthorized: Token missing or invalid", http.StatusUnauthorized)
		return
	}
	_, err = db.Db.Exec("DELETE FROM sessions WHERE token = ?", cookie.Value)
	if err != nil {
		helpers.Writer(w, map[string]string{"Error": "Failed to log out user."}, http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   "",
		Path:    "/",
		Expires: time.Now(),
	})
	helpers.Writer(w, map[string]string{"Message": "Logout successfuly"}, 200)
}
