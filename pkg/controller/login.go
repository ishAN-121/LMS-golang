package controller

import (
	"net/http"


	"LMS/pkg/views"
	"LMS/pkg/types"
	"LMS/pkg/models"

)

func LoginPage(response http.ResponseWriter, request *http.Request){
	var error types.Error
	error.Message = ""
	tempelateFunc := views.GetTemplate("loginPage")
	template := tempelateFunc()
	template.Execute(response,error)
}

func Login(response http.ResponseWriter, request *http.Request){
	var user types.User
	var message types.Error
	var error error

	user.Username = request.FormValue("username")
	user.Password = request.FormValue("password")

	if (user.Username == "" || user.Password == "" ){
		message.Message = "Enter all the details"
	}

	user.Admin, message,error = models.Authenticate(response,request,user.Username,user.Password)
	if error !=nil {
		http.Redirect(response, request, "/serverError", http.StatusFound)
	}

	if ((message.Message != "")&& (message.Message != "Login successful")){
		tempelateFunc := views.GetTemplate("login")
		template := tempelateFunc()
		template.Execute(response,error)
	}else{

		if !user.Admin{
			http.Redirect(response, request, "/user", http.StatusSeeOther)
		}else{
			http.Redirect(response, request, "/admin", http.StatusSeeOther)
		}
	}
	
}