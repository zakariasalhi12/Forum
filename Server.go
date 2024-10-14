package main

import (
	"fmt"
	"log"
	"net/http"

	forum "forum/BackEnd/Handlers"
	"forum/BackEnd/api"
	"forum/BackEnd/db"
)

const (
	Port  = ":8080"
	Dns   = "localhost"
	Red   = "\033[1;31m"
	Green = "\033[1;32m"
	Rest  = "\033[0,0m"
)

func main() {
	if err := db.ConnectTodb(); err != nil {
		log.Fatal(Red, err.Error(), Rest)
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("FrontEnd/static"))))
	// handlers

	http.HandleFunc("/", forum.HandleMain)
	http.HandleFunc("/login", forum.HandleLogin)
	http.HandleFunc("/register", forum.HandleRegister)

	// apis

	http.HandleFunc("/api/login", api.LoginApi)
	http.HandleFunc("/api/register", api.RegisterAPI)
	http.HandleFunc("/api/create", api.PostsAPI)
	// http.HandleFunc("/api/like", nil)
	// http.HandleFunc("/api/delete", nil)
	// http.HandleFunc("/api/update", nil)
	// http.HandleFunc("/api/comment", nil)

	fmt.Println(Green + "Server Started at : http://" + Dns + Port + Rest)
	if err := http.ListenAndServe(Port, nil); err != nil {
		log.Fatal(Red + err.Error() + Rest)
	}
}
