package db

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func ConnectTodb(name string) error {
	// open and create the db if the db is not exist
	var err error
	Db, err = sql.Open("sqlite3", name)
	if err != nil {
		return err
	}
	// ping the db to see if it connected
	if err = Db.Ping(); err != nil {
		return err
	}
	Schema, err := os.ReadFile("BackEnd/db/schema/setup.sql")
	if err != nil {
		return err
	}
	_, err = Db.Exec(string(Schema))
	if err != nil {
		return err
	}
	return nil
}
