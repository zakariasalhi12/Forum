package models

import (
	"errors"
	"html"
	"strings"

	"forum/BackEnd/config"
	"forum/BackEnd/helpers"
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

func (p *Posts) CheckPost() error {
	p.Title, p.Content = helpers.RemoveExtraSpaces(p.Title), helpers.RemoveExtraSpaces(p.Content)
	if helpers.CheckEmpty(p.Title, p.Content) {
		return errors.New("request Cant be empty")
	}

	if len(p.Content) >= 250 {
		return errors.New("the maximum post content length is 250 characters")
	}
	if len(p.Title) >= 50 {
		return errors.New("the maximum post title length is 250 characters")
	}
	if err := p.RemoveDuplicatedInCategorys(); err != nil {
		return err
	}
	return nil
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

func (p *Posts) RemoveDuplicatedInCategorys() error {
	res := []string(nil)
	for _, element := range strings.Fields(strings.Join(p.Categories, " ")) {
		if len(element) >= 15 {
			return errors.New("The maximum topic length is 15 characters. : " + html.EscapeString(element))
		}
		if !Include(res, element) && element != "" {
			res = append(res, html.EscapeString(element))
		}
	}
	if len(res) > 6 {
		return errors.New("the maximum topic lenght is 6")
	}
	p.Categories = res
	return nil
}

func Include(arr []string, sep string) bool {
	for _, element := range arr {
		if element == sep {
			return true
		}
	}
	return false
}
