package helpers

import "time"

type Error struct {
	Error      string
	StatusCode int
}

type Register struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string
}

type Login struct {
	ID       int
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Posts struct {
	User_ID    int
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	Categories []string `json:"categories"`
}

type Comment struct {
	PostId  int
	UserID  int
	Content string
}

type LikesDislikes struct {
	PostOrCommentId int
	IsComment       bool
	IsLike          bool
}

type Likes struct {
	Counter int
	IsLiked bool
}

type Dislikes struct {
	Counter   int
	IsDislike bool
}

type AllPosts struct {
	Id         int
	User_id    int
	Title      string
	Content    string
	Categories []string
	Comments   []Comments
	Likes      Likes
	Dislikes   Dislikes
	CreatedAt  time.Time
}

type Comments struct {
	Id        int
	UserID    int
	UserName  string
	Content   string
	Likes     Likes
	Dislikes  Dislikes
	CreatedAt time.Time
}
