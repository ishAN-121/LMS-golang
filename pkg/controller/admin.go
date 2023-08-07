package controller

import (
	"net/http"
	"LMS/pkg/types"
	"LMS/pkg/views"
)

func Adminpage(w http.ResponseWriter, r *http.Request){
	t := views.Adminpage()
	var user types.User
	user.Username = r.Header.Get("username")
	t.Execute(w,user)
}