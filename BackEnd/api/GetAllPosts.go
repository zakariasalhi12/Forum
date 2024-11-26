package api

import (
	"database/sql"
	"net/http"

	"forum/BackEnd/db"
	helpers "forum/BackEnd/helpers/Api"
)

func AllPostsApi(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helpers.Writer(w, map[string]string{"Error": "Methode not allowed"}, http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var NewPosts []helpers.AllPosts
	NewPosts, err := GetPosts(r)
	if err != nil {
		helpers.Writer(w, map[string]string{"Error": err.Error()}, http.StatusInternalServerError)
		return
	}
	helpers.Writer(w, NewPosts, 200)
}

func GetPosts(r *http.Request) ([]helpers.AllPosts, error) {
	var posts []helpers.AllPosts

	rows, err := db.Db.Query("SELECT id, user_id, title, content FROM posts ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post helpers.AllPosts
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

func GetComments(r *http.Request, postId int) ([]helpers.Comments, error) {
	var Comments []helpers.Comments
	rows, err := db.Db.Query("SELECT id , user_id , content FROM comments WHERE post_id = ? ORDER BY created_at DESC", postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		comment := helpers.Comments{}
		if err := rows.Scan(&comment.Id, &comment.UserID, &comment.Content); err != nil {
			return nil, err
		}
		UserName, err := helpers.GetUserName(comment.UserID)
		if err != nil {
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
		comment.UserName = UserName
		Comments = append(Comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return Comments, nil
}

func GetLikes(r *http.Request, Id int, isComment bool) (helpers.Likes, error) {
	var Likes helpers.Likes

	UserID, err := helpers.GetUserID(r)
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
		return helpers.Likes{}, err
	}
	return Likes, nil
}

func GetDislikes(r *http.Request, Id int, isComment bool) (helpers.Dislikes, error) {
	var Dislikes helpers.Dislikes

	UserID, err := helpers.GetUserID(r)
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
		return helpers.Dislikes{}, err
	}
	return Dislikes, nil
}
