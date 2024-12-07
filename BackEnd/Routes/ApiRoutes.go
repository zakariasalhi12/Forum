package routes

import (
	"net/http"

	"forum/BackEnd/api"
	Auth "forum/BackEnd/api/Auth"
)

func ApiRoutes() {
	http.HandleFunc("/api/login", Auth.LoginApi)
	http.HandleFunc("/api/logout", Auth.LogoutAPI)
	http.HandleFunc("/api/register", Auth.RegisterAPI)
	http.HandleFunc("/api/post", api.PostsAPI)
	http.HandleFunc("/api/like", api.AddLikeAPI)
	http.HandleFunc("/api/comment", api.NewCommentAPI)
	http.HandleFunc("/api/posts", api.AllPostsApi)
	http.HandleFunc("/api/isloged", Auth.Islogged)
	http.HandleFunc("/api/userinfo", api.GetUserInfo)
}
