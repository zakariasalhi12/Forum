package api

import (
	"net/http"

	"forum/BackEnd/utils"
)

func AddLikeAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.Writer(w, map[string]string{"Error": "Methode not allowed"}, http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	utils.Writer(w, map[string]string{"Message": "Request successfuly"}, 200)
}
