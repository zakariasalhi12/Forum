package config

import (
	"database/sql"
	"log"
	"os"
)

type ServerConfig struct {
	Port              string
	DNS               string
	Database          *sql.DB
	ApiLogs           bool
	ServerLogs        bool
	SaveLogs          bool
	LogsDirPath       string
	TotalPostsPerPage int
}

var Config = &ServerConfig{
	Port:              ":8082",
	DNS:               "localhost",
	ApiLogs:           true,
	ServerLogs:        true,
	SaveLogs:          true,
	LogsDirPath:       "Logs/",
	TotalPostsPerPage: 20, // NOTE: If you change the value of this variable, make sure to update it in the front-end as well.
}

func (s *ServerConfig) ApiLogGenerator(str string) {
	if !s.ApiLogs {
		return
	}
	log.SetOutput(os.Stdin)
	log.Println(str)
	if s.SaveLogs {
		file, err := os.OpenFile(s.LogsDirPath+"apilogs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
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
	log.SetOutput(os.Stdin)
	log.Println(str)
	if s.SaveLogs {
		file, err := os.OpenFile(s.LogsDirPath+"serverlogs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
		if err != nil {
			log.Println(err)
			return
		}
		defer file.Close()
		log.SetOutput(file)
		log.Println(str)
	}
}
