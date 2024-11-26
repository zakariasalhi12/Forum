package api

import (
	"database/sql"
	"net/http"

	"forum/BackEnd/db"
	"forum/BackEnd/utils"
)

func AllPostsApi(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.Writer(w, map[string]string{"Error": "Methode not allowed"}, http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var NewPosts []utils.AllPosts
	NewPosts, err := GetPosts(r)
	if err != nil {
		utils.Writer(w, map[string]string{"Error": err.Error()}, http.StatusBadRequest)
		return
	}
	utils.Writer(w, NewPosts, 200)
}

func GetPosts(r *http.Request) ([]utils.AllPosts, error) {
	var posts []utils.AllPosts

	rows, err := db.Db.Query("SELECT id, user_id, title, content FROM posts ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post utils.AllPosts
		if err := rows.Scan(&post.Id, &post.User_id, &post.Title, &post.Content); err != nil {
			return nil, err
		}
		PostCategories, err := GetCategories(post.Id)
		if err != nil {
			return nil, err
		}
		PostComments, err := GetComments(r, post.Id)
		if err != nil {
			return nil, err
		}
		Likes, err := GetLikes(r, post.Id, false)
		if err != nil {
			return nil, err
		}
		Dislikes, err := GetDislikes(r, post.Id, false)
		if err != nil {
			return nil, err
		}
		post.Categories = PostCategories
		post.Comments = PostComments
		post.Likes = Likes
		post.Dislikes = Dislikes
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func GetCategories(PostId int) ([]string, error) {
	Categories := []string(nil)

	rows, err := db.Db.Query("SELECT categorie FROM categories WHERE post_id = ?", PostId)
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

func GetComments(r *http.Request, postId int) ([]utils.Comment2, error) {
	var Comments []utils.Comment2
	rows, err := db.Db.Query("SELECT id , user_id , content FROM comments WHERE post_id = ?", postId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		comment := utils.Comment2{}
		if err := rows.Scan(&comment.Id, &comment.UserID, &comment.Content); err != nil {
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
		Comments = append(Comments, comment)
	}

	return Comments, nil
}

func GetLikes(r *http.Request, Id int, isComment bool) (utils.Likes, error) {
	var Likes utils.Likes

	UserID, err := GetUserID(r)
	if err == nil {
		var exists int
		db.Db.QueryRow("SELECT COUNT(*) FROM likes_dislikes WHERE post_or_comment_id = ? AND is_comment = ? AND user_id = ? AND is_like = TRUE", Id, isComment, UserID).Scan(&exists)
		if exists == 1 {
			Likes.IsLiked = true
		}
	}

	err = db.Db.QueryRow("SELECT COUNT(*) FROM likes_dislikes WHERE post_or_comment_id = ? AND is_comment = ? AND is_like = TRUE ", Id, isComment).Scan(&Likes.Counter)
	if err == sql.ErrNoRows {
		return Likes, nil
	}
	if err != nil {
		return utils.Likes{}, err
	}
	return Likes, nil
}

func GetDislikes(r *http.Request, Id int, isComment bool) (utils.Dislikes, error) {
	var Dislikes utils.Dislikes

	UserID, err := GetUserID(r)
	if err == nil {
		var exists int
		db.Db.QueryRow("SELECT COUNT(*) FROM likes_dislikes WHERE post_or_comment_id = ? AND is_comment = ? AND user_id = ? AND is_like = FALSE", Id, isComment, UserID).Scan(&exists)
		if exists == 1 {
			Dislikes.IsDislike = true
		}
	}
	err = db.Db.QueryRow("SELECT COUNT(*) FROM likes_dislikes WHERE post_or_comment_id = ? AND is_comment = ? AND is_like = FALSE ", Id, isComment).Scan(&Dislikes.Counter)
	if err == sql.ErrNoRows {
		return Dislikes, nil
	}
	if err != nil {
		return utils.Dislikes{}, err
	}
	return Dislikes, nil
}
