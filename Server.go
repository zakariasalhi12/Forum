package main

import (
	"log"
	"net/http"

	forum "forum/BackEnd/Handlers"
	"forum/BackEnd/api"
	Auth "forum/BackEnd/api/Auth"
	"forum/BackEnd/db"
)

const (
	Port  = ":8080"
	Red   = "\033[1;31m"
	Green = "\033[1;32m"
	Rest  = "\033[0;0m"
)

func main() {
	if err := db.ConnectTodb("forum.db"); err != nil {
		log.Fatal(Red, err.Error(), Rest)
	}
	defer db.Db.Close()
	log.Println(Green, "Database connected successfully!", Rest)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("FrontEnd/static"))))
	// handlers

	http.HandleFunc("/", forum.HandleMain)
	http.HandleFunc("/login", forum.HandleLogin)
	http.HandleFunc("/register", forum.HandleRegister)

	// api
	http.HandleFunc("/api/login", Auth.LoginApi)
	http.HandleFunc("/api/logout", Auth.LogoutAPI)
	http.HandleFunc("/api/register", Auth.RegisterAPI)
	http.HandleFunc("/api/post", api.PostsAPI)
	http.HandleFunc("/api/like", api.AddLikeAPI)
	http.HandleFunc("/api/comment", api.NewCommentAPI)
	http.HandleFunc("/api/posts", api.AllPostsApi)

	log.Println(Green + "Server Started at : http://localhost" + Port + Rest)
	if err := http.ListenAndServe(Port, nil); err != nil {
		log.Fatal(Red + err.Error() + Rest)
	}
}
