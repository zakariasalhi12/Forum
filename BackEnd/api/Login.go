package api

import (
	"database/sql"
	"net/http"
	"strconv"

	"forum/BackEnd/db"
	"forum/BackEnd/utils"

	"github.com/gofrs/uuid"
)

func LoginApi(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ErrorWriter(w, "Methode not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	user := utils.Login{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	var UserId int
	err := db.Db.QueryRow("SELECT id FROM users WHERE email = ? AND password = ?", user.Email, user.Password).Scan(&UserId)
	if err == sql.ErrNoRows {
		utils.Writer(w, map[string]string{"Error": "Email or Password Incorrect"}, 400)
		return
	}
	if err != nil {
		utils.Writer(w, map[string]string{"Error": err.Error()}, 500)
		return
	}
	uuid, err := uuid.NewV4()
	if err != nil {
		utils.Writer(w, map[string]string{"Error": "An unexpected error occurred. Please try again later."}, 500)
	}
	utils.Writer(w, map[string]string{"token": uuid.String(), "userid": strconv.Itoa(UserId)}, 200)
}
