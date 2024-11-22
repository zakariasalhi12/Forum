package utils

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
	User_ID    int
	Title      string
	Content    string
	Categories []string
}

type Comment struct {
	PostId  int
	UserID  int
	Content string
}

type LikesDislikes struct {
	PostId    int
	IsComment bool
	IsLike    bool
}
