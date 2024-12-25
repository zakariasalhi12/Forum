package api

import (
	"database/sql"
	"net/http"
	"strconv"

	models "forum/BackEnd/Models"
	"forum/BackEnd/config"
	"forum/BackEnd/helpers"
)

func AllPostsApi(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Method != http.MethodGet {
		helpers.Writer(w, map[string]string{"Error": helpers.ErrServer.Error()}, http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	PostID := r.FormValue("id")
	// add pages
	Filter := r.URL.Query().Get("filter")
	Tagfilter := r.URL.Query().Get("tagfilter")
	Offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	limit := 5

	var NewPosts []models.AllPosts
	if Filter == "post" {
		NewPosts, err = GetMyPosts(r, Filter, Offset, limit)
	} else if Filter == "like" {
		NewPosts, err = GetLikePosts(r, Filter, Offset, limit)
	} else if Filter == "Tag" {
		NewPosts, err = GetTagPosts(r, Filter, Tagfilter, Offset, limit)
	} else {
		NewPosts, err = GetPosts(r, PostID, Offset, limit)
	}
	if err != nil {
		helpers.Writer(w, map[string]string{"Error": err.Error()}, http.StatusInternalServerError)
		return
	}
	helpers.Writer(w, NewPosts, 200)
}

func GetPosts(r *http.Request, PostId string, offset, limit int) ([]models.AllPosts, error) {
	var posts []models.AllPosts

	if PostId != "" {
		rows, err := config.Config.Database.Query("SELECT id, user_id, title, content , created_at FROM posts WHERE id = ?", PostId)
		if err != nil {
			return nil, err
		}
		if err := GetAllPosts(r, rows, &posts); err != nil {
			return nil, err
		}
		return posts, nil
	}
	rows, err := config.Config.Database.Query("SELECT id, user_id, title, content , created_at FROM posts ORDER BY created_at DESC LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, err
	}
	if err := GetAllPosts(r, rows, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}

func GetAllPosts(r *http.Request, rows *sql.Rows, posts *[]models.AllPosts) error {
	defer rows.Close()
	for rows.Next() {
		var post models.AllPosts
		if err := rows.Scan(&post.Id, &post.User_id, &post.Title, &post.Content, &post.CreatedAt); err != nil {
			return err
		}
		PostCategories, err := GetCategories(post.Id)
		if err != nil {
			return err
		}
		PostComments, err := GetComments(r, post.Id)
		if err != nil {
			return err
		}
		Likes, err := GetLikes(r, post.Id, false)
		if err != nil {
			return err
		}
		Dislikes, err := GetDislikes(r, post.Id, false)
		if err != nil {
			return err
		}
		User := &models.User{Id: post.User_id}
		if err := User.GetUserName(); err != nil {
			return err
		}
		post.Categories, post.Comments, post.Likes, post.Dislikes, post.UserName = PostCategories, PostComments, Likes, Dislikes, User.UserName
		*posts = append(*posts, post)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

func GetCategories(PostId int) ([]string, error) {
	Categories := []string(nil)

	rows, err := config.Config.Database.Query("SELECT categorie FROM categories WHERE post_id = ?", PostId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var Categorie string
		if err := rows.Scan(&Categorie); err != nil {
			return nil, err
		}
		Categories = append(Categories, Categorie)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return Categories, nil
}

func GetComments(r *http.Request, postId int) ([]models.Comments, error) {
	var Comments []models.Comments
	rows, err := config.Config.Database.Query("SELECT id , user_id , content , created_at FROM comments WHERE post_id = ? ORDER BY created_at DESC", postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		comment := models.Comments{}
		if err := rows.Scan(&comment.Id, &comment.UserID, &comment.Content, &comment.CreatedAt); err != nil {
			return nil, err
		}
		User := &models.User{Id: comment.UserID}
		if err := User.GetUserName(); err != nil {
			return nil, err
		}
		likes, err := GetLikes(r, comment.Id, true)
		if err != nil {
			return nil, err
		}
		Dislikes, err := GetDislikes(r, comment.Id, true)
		if err != nil {
			return nil, err
		}
		comment.Likes = likes
		comment.Dislikes = Dislikes
		comment.UserName = User.UserName
		Comments = append(Comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return Comments, nil
}

func GetLikes(r *http.Request, Id int, isComment bool) (models.Likes, error) {
	var Likes models.Likes
	Session := &models.Session{}
	err := Session.GetUserID(r)
	if err == nil {
		var exists int
		config.Config.Database.QueryRow("SELECT COUNT(*) FROM likes_dislikes WHERE post_or_comment_id = ? AND is_comment = ? AND user_id = ? AND is_like = TRUE", Id, isComment, Session.UserID).Scan(&exists)
		if exists == 1 {
			Likes.IsLiked = true
		}
	}
	err = config.Config.Database.QueryRow("SELECT COUNT(*) FROM likes_dislikes WHERE post_or_comment_id = ? AND is_comment = ? AND is_like = TRUE ", Id, isComment).Scan(&Likes.Counter)
	if err == sql.ErrNoRows {
		return Likes, nil
	}
	if err != nil {
		return models.Likes{}, err
	}
	return Likes, nil
}

func GetDislikes(r *http.Request, Id int, isComment bool) (models.Dislikes, error) {
	var Dislikes models.Dislikes
	Session := &models.Session{}
	err := Session.GetUserID(r)
	if err == nil {
		var exists int
		config.Config.Database.QueryRow("SELECT COUNT(*) FROM likes_dislikes WHERE post_or_comment_id = ? AND is_comment = ? AND user_id = ? AND is_like = FALSE", Id, isComment, Session.UserID).Scan(&exists)
		if exists == 1 {
			Dislikes.IsDislike = true
		}
	}
	err = config.Config.Database.QueryRow("SELECT COUNT(*) FROM likes_dislikes WHERE post_or_comment_id = ? AND is_comment = ? AND is_like = FALSE ", Id, isComment).Scan(&Dislikes.Counter)
	if err == sql.ErrNoRows {
		return Dislikes, nil
	}
	if err != nil {
		return models.Dislikes{}, err
	}
	return Dislikes, nil
}
