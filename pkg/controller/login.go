package controller

import (
	"net/http"


	"LMS/pkg/views"
	"LMS/pkg/types"
	"LMS/pkg/models"

)

func LoginPage(w http.ResponseWriter, r *http.Request){
	t := views.LoginPage()
	var err types.Error
	err.Msg = ""
	t.Execute(w,err)
}

func Login(w http.ResponseWriter, r *http.Request){
	var user types.User
	var msg types.Error
	var err error

	user.Username = r.FormValue("username")
	user.Password = r.FormValue("password")

	if (user.Username == "" || user.Password == "" ){
		msg.Msg = "Enter all the details"
	}

	user.Admin, msg,err = models.Authenticate(w,r,user.Username,user.Password)
	if err !=nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
	}

	if ((msg.Msg != "")&& (msg.Msg != "Login successful")){
	t := views.LoginPage()
	t.Execute(w,err)
	}else{

		if !user.Admin{
			http.Redirect(w, r, "/user", http.StatusSeeOther)
		}else{
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
		}
	}
	
}