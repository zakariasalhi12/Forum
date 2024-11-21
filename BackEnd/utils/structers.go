package utils

import "time"

type Error struct {
	Error      string
	StatusCode int
}

type Register struct {
	UserName string
	Email    string
	Password string
	Role     string
}

type Login struct {
	ID       int
	Email    string
	Password string
}

type Posts struct {
	Title       string
	Description string
	Likes       int
	Dislikes    int
	Comments    []Comment
}

type Comment struct {
	Id       int
	PostId   int
	UserID   int
	Likes    int
	Dislikes int
	Liked    bool
}

type Session struct {
	Id        int
	UserID    int
	Token     string
	CreatedAt time.Time
}
