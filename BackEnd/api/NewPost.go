package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"forum/BackEnd/db"
	"forum/BackEnd/utils"
)

func PostsAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.Writer(w, map[string]string{"Error": "Methode not allowed"}, http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var NewPost utils.Posts
	if err := json.NewDecoder(r.Body).Decode(&NewPost); err != nil {
		utils.Writer(w, map[string]string{"Error": "An unexpected error occurred. Please try again later."}, 500)
		return
	}
	UserID, err := GetUserID(r)
	if err != nil {
		utils.Writer(w, map[string]string{"Error": err.Error()}, http.StatusBadRequest)
		return
	}
	Res, err := db.Db.Exec("INSERT INTO posts (user_id ,title ,content) VALUES (? ,? ,?)", UserID, NewPost.Title, NewPost.Content)
	if err != nil {
		utils.Writer(w, map[string]string{"Error": err.Error()}, http.StatusInternalServerError)
		return
	}
	LastID, err := Res.LastInsertId()
	if err != nil {
		utils.Writer(w, map[string]string{"Error": err.Error()}, http.StatusInternalServerError)
		return
	}
	err = InsertToCategory(NewPost.Categories, LastID)
	if err != nil {
		utils.Writer(w, map[string]string{"Error": err.Error()}, http.StatusInternalServerError)
		return
	}
	utils.Writer(w, map[string]string{"Message": "Post Created successfuly"}, 200)
}

func GetUserID(r *http.Request) (int, error) {
	var userID int
	tokenCookie, err := r.Cookie("token")
	if err != nil {
		return -1, errors.New("you must be logged in to create a post")
	}
	err = db.Db.QueryRow("SELECT user_id FROM sessions WHERE token = ?", tokenCookie.Value).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return -1, errors.New("session not found or expired")
		}
		return -1, err
	}

	return userID, nil
}

func InsertToCategory(categories []string, postid int64) error {
	for _, categorie := range categories {
		_, err := db.Db.Exec("INSERT INTO categories (post_id , categorie) VALUES (?, ?)", postid, categorie)
		if err != nil {
			return err
		}
	}
	return nil
}
