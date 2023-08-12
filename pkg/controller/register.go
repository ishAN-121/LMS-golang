package controller

import(
	"net/http"

	"LMS/pkg/views"
	"LMS/pkg/types"
	"LMS/pkg/models"
)

func Register(response http.ResponseWriter, request *http.Request){

	var error types.Error
	error.Message = ""
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
	confirmPassword := request.FormValue("confirm_password")
	

	if (user.Username == "" || user.Password == "" || confirmPassword == ""){
		message.Message = "Enter all the details"
	}else if (user.Password != confirmPassword) {
		message.Message = "Passwords do not match"
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