package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func ConnectTodb(name string) (error) {
	// open and create the db if the db is not exist
	Db, err := sql.Open("sqlite3", name)
	if err != nil {
		return err
	}
	defer Db.Close()
	// ping the db to see if it connected
	if err = Db.Ping(); err != nil {
		return err
	}

	createTables()
	return nil
}
