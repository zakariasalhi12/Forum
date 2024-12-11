package models

type Comment struct {
	UserID  int
	PostId  int    `json:"postid"`
	Content string `json:"content"`
}
