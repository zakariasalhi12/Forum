package config

import "database/sql"

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
