package controller

import(
	"net/http"

	"LMS/pkg/views"
	"LMS/pkg/types"
	"LMS/pkg/models"
)

func Register(response http.ResponseWriter, request *http.Request){

	var error types.Error
	error.Msg = ""
	tempelateFunc := views.GetTemplate("registerPage")
	template := tempelateFunc()
	template.Execute(response,error)
}

func AddUser(response http.ResponseWriter, request *http.Request){

	var message types.Error
	var error error
	var user types.User
	user.Username = request.FormValue("username")
	user.Password = request.FormValue("password")
	confirmPassword := request.FormValue("confirmPassword")
	

	if (user.Username == "" || user.Password == "" || confirmPassword == ""){
		message.Msg = "Enter all the details"
	}else if (user.Password != confirmPassword) {
		message.Msg = "Passwords do not match"
	}else {
		message,error = models.AddUser(user.Username, user.Password)
		if error != nil{
			http.Redirect(response, request, "/serverError", http.StatusFound)
		}
	}
	tempelateFunc := views.GetTemplate("registerPage")
	template := tempelateFunc()
	template.Execute(response,message)
}