package api

import (
	"encoding/json"
	"net/http"

	"forum/BackEnd/utils"
)

func PostsAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ErrorWriter(w, "Methode not allowed", http.StatusMethodNotAllowed)
		return
	}
	var NewPost utils.Posts


	// need to check the data after receve it

	NewPost.Title = r.FormValue("title")
	NewPost.Description = r.FormValue("description")

	Response, err := json.Marshal(NewPost)
	if err != nil {
		utils.ErrorWriter(w, "Cannot marchal data", 500)
		return
	}
	w.WriteHeader(200)
	w.Write(Response)
}
