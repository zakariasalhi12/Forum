package api

import (
	"net/http"

	"forum/BackEnd/db"
	"forum/BackEnd/utils"
)

func LogoutAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		utils.Writer(w, map[string]string{"Error": "Methode not allowed"}, http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	cookie, err := r.Cookie("token")
	if err != nil {
		utils.Writer(w, "Unauthorized: Token missing or invalid", http.StatusUnauthorized)
		return
	}

	_, err = db.Db.Exec("DELETE FROM sessions WHERE token = ?", cookie.Value)
	if err != nil {
		utils.Writer(w, map[string]string{"Error": "Failed to log out user."}, http.StatusInternalServerError)
		return
	}

	utils.Writer(w, map[string]string{"Message": "Logout successful"}, 200)
}
