package db

import (
	"log"
)

func createTables() {
	userTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);`

	postTable := `
	CREATE TABLE IF NOT EXISTS posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		category TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);`

	commentTable := `
	CREATE TABLE IF NOT EXISTS comments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		post_id INTEGER,
		user_id INTEGER,
		content TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(post_id) REFERENCES posts(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	);`

	likeDislikeTable := `
	CREATE TABLE IF NOT EXISTS likes_dislikes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		post_id INTEGER,
		user_id INTEGER,
		is_like BOOLEAN,
		FOREIGN KEY(post_id) REFERENCES posts(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	);`

	sessionTable := `
	CREATE TABLE IF NOT EXISTS sessions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		token TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);`

	// Execute table creation queries
	if _, err := Db.Exec(userTable); err != nil {
		log.Fatalf("Error creating users table: %v", err)
	}
	if _, err := Db.Exec(postTable); err != nil {
		log.Fatalf("Error creating posts table: %v", err)
	}
	if _, err := Db.Exec(commentTable); err != nil {
		log.Fatalf("Error creating comments table: %v", err)
	}
	if _, err := Db.Exec(likeDislikeTable); err != nil {
		log.Fatalf("Error creating likes_dislikes table: %v", err)
	}
	if _, err := Db.Exec(sessionTable); err != nil {
		log.Fatalf("Error creating sessions table: %v", err)
	}
}
