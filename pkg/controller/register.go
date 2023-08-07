package controller

import(
	"net/http"
	"fmt"

	"LMS/pkg/views"
	"LMS/pkg/types"
	"LMS/pkg/models"
)

func Register(w http.ResponseWriter, r *http.Request){
	t := views.RegisterPage()
	var err types.Error
	err.Msg = ""
	t.Execute(w,err)
}

func Adduser(w http.ResponseWriter, r *http.Request){
	username := r.FormValue("username")
	password := r.FormValue("password")
	confirmpassword := r.FormValue("confirmpassword")
	fmt.Println(username)
	var err types.Error

	if (username == "" || password == "" || confirmpassword == ""){
		err.Msg = "Enter all the details"
	}else if (password != confirmpassword) {
		err.Msg = "Passwords do not match"
	}else {
		err = models.Adduser(username, password, confirmpassword)
	}

	
	t := views.RegisterPage()
	
	t.Execute(w,err)

}