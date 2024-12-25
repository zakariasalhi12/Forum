package api

import (
	"net/http"

	models "forum/BackEnd/Models"
	"forum/BackEnd/config"
)

func GetMyPosts(r *http.Request, filter string, offset, limit int) ([]models.AllPosts, error) {
	Session := &models.Session{}
	err := Session.GetUserID(r)
	if err != nil {
		return nil, err
	}
	UserID := Session.UserID

	var posts []models.AllPosts
	rows, err := config.Config.Database.Query("SELECT id, user_id, title, content, created_at FROM posts WHERE user_id = ? ORDER BY created_at DESC LIMIT ? OFFSET ?", UserID, limit, offset)
	if err != nil {
		return nil, err
	}
	if err := GetAllPosts(r, rows, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}

func GetLikePosts(r *http.Request, filter string, offset, limit int) ([]models.AllPosts, error) {
	Session := &models.Session{}
	err := Session.GetUserID(r)
	if err != nil {
		return nil, err
	}
	UserID := Session.UserID

	var posts []models.AllPosts
	query := `
		SELECT DISTINCT p.id, p.user_id, p.title, p.content, p.created_at 
		FROM posts p
		JOIN likes_dislikes ld ON p.id = ld.post_or_comment_id
		WHERE ld.user_id = ? AND ld.is_like = TRUE AND ld.is_comment = FALSE 
		ORDER BY p.created_at DESC LIMIT ? OFFSET ?
	`
	rows, err := config.Config.Database.Query(query, UserID, limit, offset)
	if err != nil {
		return nil, err
	}

	if err := GetAllPosts(r, rows, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}

func GetTagPosts(r *http.Request, filter, tagfilter string, offset, limit int) ([]models.AllPosts, error) {
	var posts []models.AllPosts
	query := `
		SELECT DISTINCT p.id, p.user_id, p.title, p.content, p.created_at 
		FROM posts p
		JOIN categories c ON p.id = c.post_id
		WHERE c.categorie = ? 
		ORDER BY p.created_at DESC LIMIT ? OFFSET ?
	`
	rows, err := config.Config.Database.Query(query, tagfilter, limit, offset)
	if err != nil {
		return nil, err
	}

	if err := GetAllPosts(r, rows, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}
