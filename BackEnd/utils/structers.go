package utils

type Error struct {
	Error      string
	StatusCode int
}

type Register struct {
	UserName string
	Email    string
	Password string
}

type Login struct {
	ID int
	Email    string
	Password string
}

type Posts struct {
	Title       string
	Description string
	Likes       int
	Dislikes    int
	Comments    map[string]string
}
