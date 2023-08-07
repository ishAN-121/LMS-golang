package controller

import (
	"net/http"
	"fmt"

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
	var err types.Error
	var admin bool

	user.Username = username

	
	if (username == "" || password == "" ){
		err.Msg = "Enter all the details"
	}
	admin, err = models.Authenticate(w,r,username,password)
	if ((err.Msg != "")&& (err.Msg != "Login successful")){
	t := views.LoginPage()
	t.Execute(w,err)
	}else{
		fmt.Println(admin)
		if !admin{
			http.Redirect(w, r, "/user", http.StatusSeeOther)
		}else{
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
		}
	}
	
}