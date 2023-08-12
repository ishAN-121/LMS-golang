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

func AddUser(w http.ResponseWriter, r *http.Request){

	var msg types.Error
	var err error
	var user types.User
	user.Username = r.FormValue("username")
	user.Password = r.FormValue("password")
	confirmPassword := r.FormValue("confirmPassword")
	

	if (user.Username == "" || user.Password == "" || confirmPassword == ""){
		msg.Msg = "Enter all the details"
	}else if (user.Password != confirmPassword) {
		msg.Msg = "Passwords do not match"
	}else {
		msg,err = models.AddUser(user.Username, user.Password)
		if err != nil{
			http.Redirect(w, r, "/serverError", http.StatusFound)
		}
	}
	t := views.RegisterPage()
	t.Execute(w,msg)
}