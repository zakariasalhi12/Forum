package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectTodb() error {
	dbName := "forum.db"
	// open and create the db if the db is not exist
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return err
	}
	defer db.Close()
	// ping the db to see if it connected
	err = db.Ping()
	if err != nil {
		return err
	}

	fmt.Printf("Database '%s' connected successfully!\n", dbName)
	return nil
}
