package controller

import(
	"net/http"

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
	var msg types.Error
	var err error

	if (username == "" || password == "" || confirmpassword == ""){
		msg.Msg = "Enter all the details"
	}else if (password != confirmpassword) {
		msg.Msg = "Passwords do not match"
	}else {
		msg,err = models.Adduser(username, password, confirmpassword)
		if err != nil{
			http.Redirect(w, r, "/serverError", http.StatusFound)
		}
	}
	t := views.RegisterPage()
	t.Execute(w,err)
}