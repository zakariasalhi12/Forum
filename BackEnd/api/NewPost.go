package api

import (
	"encoding/json"
	"html"
	"io"
	"net/http"

	"forum/BackEnd/db"
	helpers "forum/BackEnd/helpers/Api_Helper"
)

func PostsAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.Writer(w, map[string]string{"Error": "Methode not allowed"}, http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var NewPost helpers.Posts
	Response, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.Writer(w, map[string]string{"Error": "An unexpected error occurred. Please try again later."}, 500)
		return
	}
	if err := json.Unmarshal(Response, &NewPost); err != nil {
		helpers.Writer(w, map[string]string{"Error": "Invalid Request"}, 400)
		return
	}
	if helpers.CheckEmpty(NewPost.Title, NewPost.Content) {
		helpers.Writer(w, map[string]string{"Error": "Request Cant be empty"}, 400)
		return
	}
	NewPost.Content = html.EscapeString(NewPost.Content)
	NewPost.Title = html.EscapeString(NewPost.Title)
	UserID, err := helpers.GetUserID(r)
	if err != nil {
		helpers.Writer(w, map[string]string{"Error": err.Error()}, http.StatusBadRequest)
		return
	}
	Res, err := db.Db.Exec("INSERT INTO posts (user_id ,title ,content) VALUES (? ,? ,?)", UserID, NewPost.Title, NewPost.Content)
	if err != nil {
		helpers.Writer(w, map[string]string{"Error": err.Error()}, http.StatusInternalServerError)
		return
	}
	LastID, err := Res.LastInsertId()
	if err != nil {
		helpers.Writer(w, map[string]string{"Error": err.Error()}, http.StatusInternalServerError)
		return
	}
	err = InsertToCategory(NewPost.Categories, LastID)
	if err != nil {
		helpers.Writer(w, map[string]string{"Error": err.Error()}, http.StatusInternalServerError)
		return
	}
	helpers.Writer(w, map[string]string{"Message": "Post Created successfuly"}, 200)
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
