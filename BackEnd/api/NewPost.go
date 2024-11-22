package api

import (
	"database/sql"
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
	UserID, err := GetUserID(r)
	if err != nil {
		utils.Writer(w, map[string]string{"Error": err.Error()}, http.StatusBadRequest)
		return
	}
	NewPost := utils.Posts{
		User_ID:    UserID,
		Title:      r.FormValue("title"),
		Content:    r.FormValue("description"),
		Categories: r.Form["categories"],
	}
	Res, err := db.Db.Exec("INSERT INTO posts (user_id ,title ,content) VALUES (? ,? ,?)", NewPost.User_ID, NewPost.Title, NewPost.Content)
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
	var UserIdint int
	UUID, err := r.Cookie("token")
	if err != nil {
		return -1, errors.New("You Most Be Loged To create a post")
	}
	err = db.Db.QueryRow("SELECT user_id FROM session WHERE token = ?", UUID.Value).Scan(&UserIdint)
	if err == sql.ErrNoRows {
		return -1, err
	}
	if err != nil {
		return -1, err
	}
	return UserIdint, nil
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
