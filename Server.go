package main

import (
	"log"
	"net/http"

	routes "forum/BackEnd/Routes"
	"forum/BackEnd/config"
	"forum/BackEnd/db"
)

func main() {
	err := db.ConnectTodb("BackEnd/db/forum.db")
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	log.Println("Database connected successfully!")
	defer config.Config.Database.Close()

	routes.HandlersRoute()
	routes.ApiRoutes()

	log.Println("Server Started at : http://" + config.Config.DNS + config.Config.Port)
	log.Fatal(http.ListenAndServe(config.Config.Port, nil))
}
