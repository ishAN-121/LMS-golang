package controller

import (
	
	"LMS/pkg/models"
	"LMS/pkg/types"
	
	"net/http"
)
var admin bool

func Middleware (next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie :=  r.Header.Get("Cookie")
		if(len(cookie) < 10){
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}else{
			cookieId := r.Header.Get("Cookie")[10:]
			var userExists bool
			var user types.User

			user.Username,userExists,user.Admin = models.Middleware(cookieId)
			admin = user.Admin
			if userExists{
				r.Header.Add("username", user.Username)
				next(w,r)
		
			}
		
		}
   	}
}

func MiddlewareDirect(next http.HandlerFunc)http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie :=  r.Header.Get("Cookie")
		
		if(len(cookie) < 10){
			next(w,r)
		}else{
			cookieid := r.Header.Get("Cookie")[10:]
			var user_exists bool
			var user types.User
			user.Username,user_exists,user.Admin = models.Middleware(cookieid)
			admin = user.Admin
			if user_exists{
				r.Header.Add("username", user.Username)
				if admin{
					http.Redirect(w, r, "/admin", http.StatusSeeOther)
				}else{
					http.Redirect(w, r, "/user", http.StatusSeeOther)
				}
			}
		}
	}
}

func IsAdmin(next http.HandlerFunc)http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if admin{
			next(w,r)
		}else{
			http.Redirect(w, r, "/user", http.StatusSeeOther)
		}
	}
}