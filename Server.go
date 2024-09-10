package main

import (
	"fmt"
	"net/http"
	"os"

	forum "forum/BackEnd/Handlers"
)

const (
	Port  = ":8080"
	Dns   = "localhost"
	Red   = "\033[1;31m"
	Green = "\033[1;32m"
	Rest  = "\033[0,0m"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("FrontEnd/static"))))

	// handlers

	http.HandleFunc("/", forum.HandleMain)
	http.HandleFunc("/login", forum.HandleLogin)
	http.HandleFunc("/register", forum.HandleRegister)

	// api

	//http.HandleFunc("/api/login", nil)
	//http.HandleFunc("/api/register", nil)
	//http.HandleFunc("/api/create", nil)
	//http.HandleFunc("/api/like", nil)
	//http.HandleFunc("/api/delete", nil)
	//http.HandleFunc("/api/update", nil)
	//http.HandleFunc("/api/comment", nil)


	fmt.Println(Green + "Server Started at : http://" + Dns + Port + Rest)
	err := http.ListenAndServe(Port, nil)
	if err != nil {
		fmt.Println(Red + err.Error() + Rest)
		os.Exit(1)
	}
}
