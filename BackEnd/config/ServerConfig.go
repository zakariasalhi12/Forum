package config

import (
	"database/sql"
	"log"
	"os"
)

type ServerConfig struct {
	Port       string
	DNS        string
	Database   *sql.DB
	ApiLogs    bool
	ServerLogs bool
	SaveLogs   bool
}

var Config = &ServerConfig{
	Port:       ":8080",
	DNS:        "localhost",
	ApiLogs:    true,
	ServerLogs: true,
	SaveLogs:   true,
}

func (s *ServerConfig) ApiLogGenerator(str string) {
	if !s.ApiLogs {
		return
	}
	log.Println(str)
	if s.SaveLogs {
		file, err := os.OpenFile("Logs/apilogs.log", os.O_CREATE|os.O_WRONLY, 0o644)
		if err != nil {
			log.Println(err)
			return
		}
		defer file.Close()
		log.SetOutput(file)
		log.Println(str)
	}
}

func (s *ServerConfig) ServerLogGenerator(str string) {
	if !s.ServerLogs {
		return
	}
	log.Println(str)
	if s.SaveLogs {
		file, err := os.OpenFile("Logs/serverlogs.log", os.O_CREATE|os.O_WRONLY, 0o644)
		if err != nil {
			log.Println(err)
			return
		}
		defer file.Close()
		log.SetOutput(file)
		log.Println(str)
	}
}
