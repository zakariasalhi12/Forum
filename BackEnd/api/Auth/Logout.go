package auth

import (
	"net/http"

	models "forum/BackEnd/Models"
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
	Session := models.NewSession(w, cookie.Value, 0)

	err = Session.DeleteSession()
	if err == models.ErrInvalidToken {
		helpers.Writer(w, map[string]string{"Unauthorized": err.Error()}, http.StatusUnauthorized)
		return
	}
	if err != nil {
		helpers.Writer(w, map[string]string{"Error": err.Error()}, 500)
		return
	}

	helpers.Writer(w, map[string]string{"Message": "Logout successfuly"}, 200)
}
