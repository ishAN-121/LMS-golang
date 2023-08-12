package controller

import (
	"net/http"
	"LMS/pkg/models"
	"LMS/pkg/types"
)

func Logout(response http.ResponseWriter, request *http.Request){
	var user types.User
	user.Username = request.Header.Get("username")

	cookie := http.Cookie{
		Name:     "sessionId",
		Value:    "",
		MaxAge:    -1,    
		HttpOnly: true, 
	}
	http.SetCookie(response,&cookie)
	models.Logout(user.Username)
	http.Redirect(response, request, "/", http.StatusFound)
}