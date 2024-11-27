package auth

import (
	"net/http"

	apihelpers "forum/BackEnd/helpers/Api_Helper"
)

func Islogged(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}
	id, err := apihelpers.GetUserID(r)
	if err != nil {
		apihelpers.Writer(w, map[string]string{"Error": err.Error()}, 400)
		return
	}
	UserName, err := apihelpers.GetUserName(id)
	if err != nil {
		apihelpers.Writer(w, map[string]string{"Error": err.Error()}, 400)
		return
	}

	apihelpers.Writer(w, map[string]string{"username": UserName}, 200)
}
