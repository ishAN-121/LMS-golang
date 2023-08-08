package controller

import (
	"net/http"
	"LMS/pkg/models"
)

func Logout(w http.ResponseWriter, r *http.Request){
	username := r.Header.Get("username")
	cookie := http.Cookie{
		Name:     "SessionID",
		Value:    "",
		MaxAge:    -1,    
		HttpOnly: true, 
	}
	http.SetCookie(w,&cookie)
	models.Logout(username)
	http.Redirect(w, r, "/", http.StatusFound)
}