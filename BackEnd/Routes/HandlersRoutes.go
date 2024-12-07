package routes

import (
	"net/http"

	"forum/BackEnd/Handlers"
)

func HandlersRoute() {
	http.Handle("/Static/", http.StripPrefix("/Static/", http.FileServer(http.Dir("FrontEnd/Static"))))
	http.HandleFunc("/", Handlers.HandleMain)
	http.HandleFunc("/register", Handlers.HandleRegister)
	http.HandleFunc("/post", Handlers.HandlePost)
}
