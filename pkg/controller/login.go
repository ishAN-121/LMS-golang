package controller

import (
	"net/http"


	"LMS/pkg/views"
	"LMS/pkg/types"
	"LMS/pkg/models"

)

func Loginpage(w http.ResponseWriter, r *http.Request){
	t := views.LoginPage()
	var err types.Error
	err.Msg = ""
	t.Execute(w,err)
}

func Login(w http.ResponseWriter, r *http.Request){
	username := r.FormValue("username")
	password := r.FormValue("password")
	
	var user types.User
	var msg types.Error
	var err error
	var admin bool

	user.Username = username

	
	if (username == "" || password == "" ){
		msg.Msg = "Enter all the details"
	}
	admin, msg,err = models.Authenticate(w,r,username,password)
	if err !=nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
	}
	if ((msg.Msg != "")&& (msg.Msg != "Login successful")){
	t := views.LoginPage()
	t.Execute(w,err)
	}else{

		if !admin{
			http.Redirect(w, r, "/user", http.StatusSeeOther)
		}else{
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
		}
	}
	
}