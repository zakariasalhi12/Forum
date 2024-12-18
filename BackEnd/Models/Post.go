package models

import (
	"forum/BackEnd/config"
	"html"
)

type Posts struct {
	User_ID    int
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	Categories []string `json:"categories"`
}

func CheckPostExist(PostId int) (bool, error) {
	var Exists bool
	if err := config.Config.Database.QueryRow("SELECT COUNT(1) FROM POSTS WHERE id = ? ", PostId).Scan(&Exists); err != nil {
		return false, err
	}
	return Exists, nil
}

func (p *Posts) AddPost() (int, error) {
	p.Content = html.EscapeString(p.Content)
	p.Title = html.EscapeString(p.Title)
	Res, err := config.Config.Database.Exec("INSERT INTO posts (user_id ,title ,content) VALUES (? ,? ,?)", p.User_ID, p.Title, p.Content)
	if err != nil {
		return 0, err
	}
	LastID, err := Res.LastInsertId()
	if err != nil {
		return 0, err
	}
	if err := p.InserCategorys(int(LastID)); err != nil {
		return 0, err
	}

	return int(LastID), nil
}

func (p *Posts) InserCategorys(PostId int) error {
	for _, categorie := range p.Categories {
		_, err := config.Config.Database.Exec("INSERT INTO categories (post_id , categorie) VALUES (?, ?)", PostId, categorie)
		if err != nil {
			return err
		}
	}
	return nil
}
