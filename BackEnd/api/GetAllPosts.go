package api

import (
	"net/http"

	"forum/BackEnd/utils"
)

func AllPostsApi(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.Writer(w, map[string]string{"Error": "Methode not allowed"}, http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var NewPosts []utils.AllPosts


	


	utils.Writer(w, NewPosts, 200)
}
