package controller

import (
	
	"LMS/pkg/models"
	"LMS/pkg/types"
	
	"net/http"
)
var admin bool

func Middleware (next http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		cookie :=  request.Header.Get("Cookie")
		if(len(cookie) < 10){
			http.Redirect(response, request, "/", http.StatusSeeOther)
		}else{
			cookieId := request.Header.Get("Cookie")[10:]
			var userExists bool
			var user types.User

			user.Username,userExists,user.Admin = models.Middleware(cookieId)
			admin = user.Admin
			if userExists{
				request.Header.Add("username", user.Username)
				next(response,request)
		
			}
		
		}
   	}
}

func MiddlewareDirect(next http.HandlerFunc)http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		cookie :=  request.Header.Get("Cookie")
		
		if(len(cookie) < 10){
			next(response,request)
		}else{
			cookieid := request.Header.Get("Cookie")[10:]
			var userExists bool
			var user types.User
			user.Username,userExists,user.Admin = models.Middleware(cookieid)
			admin = user.Admin
			if userExists{
				request.Header.Add("username", user.Username)
				if admin{
					http.Redirect(response, request, "/admin", http.StatusSeeOther)
				}else{
					http.Redirect(response, request, "/user", http.StatusSeeOther)
				}
			}
		}
	}
}

func IsAdmin(next http.HandlerFunc)http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		if admin{
			next(response,request)
		}else{
			http.Redirect(response, request, "/user", http.StatusSeeOther)
		}
	}
}

func IsUser(next http.HandlerFunc)http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		if !admin{
			next(response,request)
		}else{
			http.Redirect(response, request, "/admin", http.StatusSeeOther)
		}
	}
}