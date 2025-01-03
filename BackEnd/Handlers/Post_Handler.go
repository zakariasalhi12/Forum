package Handlers

import (
	"io"
	"net/http"
	"os"

	"forum/BackEnd/config"
	utils "forum/BackEnd/helpers"
)

func HandlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ErrorWriter(w, "Methode not allowed", http.StatusMethodNotAllowed)
		return
	}
	PostID := r.FormValue("id")
	Res, err := http.Get("http://localhost" + config.Config.Port + "/api/posts?id=" + PostID)
	if err != nil {
		config.Config.ServerLogGenerator(err.Error())
		utils.ErrorWriter(w, "An unexpected error occurred. Please try again later.", http.StatusInternalServerError)
		return
	}
	defer Res.Body.Close()
	Body, err := io.ReadAll(Res.Body)
	if err != nil {
		config.Config.ServerLogGenerator(err.Error())
		utils.ErrorWriter(w, "An unexpected error occurred. Please try again later.", http.StatusInternalServerError)
		return
	}
	if string(Body) == "null" {
		utils.ErrorWriter(w, "Invalid Post Id", http.StatusBadRequest)
		return
	}
	Data, err := os.ReadFile("FrontEnd/Templates/post.html")
	if err != nil {
		config.Config.ServerLogGenerator(err.Error())
		utils.ErrorWriter(w, "Error 500", 500)
		return
	}
	w.Write(Data)
}
