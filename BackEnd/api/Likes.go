package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"forum/BackEnd/db"
	"forum/BackEnd/utils"
)

func AddLikeAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.Writer(w, map[string]string{"Error": "Methode not allowed"}, http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	counter := 0
	Err := db.Db.QueryRow("select is_like from likes_dislikes where is_like = true")

	Err.Scan(counter)

	Res, err := json.Marshal(map[string]string{"Counter:": strconv.Itoa(counter)})
	if err != nil {
		utils.Writer(w, map[string]string{"Error": err.Error()}, http.StatusBadRequest)
		return
	}
	w.Write(Res)
}
