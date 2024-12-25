package models

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
)

type Likes struct {
	Counter int
	IsLiked bool
}

type Dislikes struct {
	Counter   int
	IsDislike bool
}

type Comments struct {
	Id        int
	UserID    int
	UserName  string
	Content   string
	Likes     Likes
	Dislikes  Dislikes
	CreatedAt string
}

type AllPosts struct {
	Id         int
	User_id    int
	UserName   string
	Title      string
	Content    string
	Categories []string
	Comments   []Comments
	Likes      Likes
	Dislikes   Dislikes
	CreatedAt  string
}

type Filter struct {
	UserPosts bool
	UserLikes bool
	Category  string
}

type Page struct {
	PostID         int
	PageNumber     int
	PageFiltration Filter
	PagePosts      []AllPosts
}

func ParseRequest(r *http.Request) (*Page, error) {
	PostID := r.FormValue("id")
	Id, err := strconv.Atoi(PostID)
	if err != nil && PostID != "" {
		return nil, errors.New("invalid postid")
	}
	PageRequest := r.FormValue("page")
	PageNumber, err := strconv.Atoi(PageRequest)
	if err != nil && PageRequest != "" {
		return nil, errors.New("invalid PageNumber")
	}
	return &Page{
		PostID:     Id,
		PageNumber: PageNumber,
		PageFiltration: Filter{
			UserPosts: strings.TrimSpace(r.FormValue("profil")) == "true",
			UserLikes: strings.TrimSpace(r.FormValue("mylikes")) == "true",
			Category:  strings.TrimSpace(r.FormValue("category")),
		},
	}, nil
}
