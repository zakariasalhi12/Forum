package main

import (
	"log"
	"net/http"

	routes "forum/BackEnd/Routes"
	"forum/BackEnd/db"
)

const (
	Port = ":8080"
)

func main() {
	err := db.ConnectTodb("BackEnd/db/forum.db")
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	log.Println("Database connected successfully!")
	defer db.Db.Close()

	routes.HandlersRoute()
	routes.ApiRoutes()

	log.Println("Server Started at : http://localhost" + Port)
	log.Fatal(http.ListenAndServe(Port, nil))
}
