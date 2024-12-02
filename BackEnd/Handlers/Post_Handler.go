package handlers

import (
	"io"
	"net/http"
	"os"

	utils "forum/BackEnd/helpers"
)

func HandlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ErrorWriter(w, "Methode not allowed", http.StatusMethodNotAllowed)
		return
	}
	PostID := r.FormValue("id")
	Res, err := http.Get("http://localhost:8080/api/posts?id=" + PostID)
	if err != nil {
		utils.ErrorWriter(w, "An unexpected error occurred. Please try again later.", http.StatusInternalServerError)
		return
	}
	defer Res.Body.Close()
	Body, err := io.ReadAll(Res.Body)
	if err != nil {
		utils.ErrorWriter(w, "An unexpected error occurred. Please try again later.", http.StatusInternalServerError)
		return
	}
	if string(Body) == "null" {
		utils.ErrorWriter(w, "Invalid Post Id", http.StatusBadRequest)
		return
	}
	Data, err := os.ReadFile("FrontEnd/templates/post.html")
	if err != nil {
		utils.ErrorWriter(w, "Error 500", 500)
		return
	}
	w.Write(Data)
}
