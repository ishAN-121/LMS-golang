package controller

import (
	"net/http"
	"LMS/pkg/models"
	"LMS/pkg/types"
)

func Logout(w http.ResponseWriter, r *http.Request){
	var user types.User
	user.Username = r.Header.Get("username")

	cookie := http.Cookie{
		Name:     "sessionId",
		Value:    "",
		MaxAge:    -1,    
		HttpOnly: true, 
	}
	http.SetCookie(w,&cookie)
	models.Logout(user.Username)
	http.Redirect(w, r, "/", http.StatusFound)
}