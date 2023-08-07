package controller

import (
	
	"LMS/pkg/models"
	
	"fmt"
	"net/http"
)
var admin bool

func Middleware (next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie :=  r.Header.Get("Cookie")
		if(len(cookie) < 10){
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}else{
			cookieid := r.Header.Get("Cookie")[10:]
			fmt.Println(cookieid)
			var user_exists bool
			var username string
			username,user_exists,admin = models.Middleware(cookieid)
			if user_exists{
				r.Header.Add("username", username)
				next(w,r)
		
			}
		
		}
   	}
}

func Middleware_direct(next http.HandlerFunc)http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie :=  r.Header.Get("Cookie")
		if(len(cookie) < 10){
			next(w,r)
		}else{
			cookieid := r.Header.Get("Cookie")[10:]
			var user_exists bool
			var username string
			username,user_exists,admin = models.Middleware(cookieid)
			if user_exists{
				r.Header.Add("username", username)
				if admin{
					http.Redirect(w, r, "/admin", http.StatusSeeOther)
				}else{
					http.Redirect(w, r, "/user", http.StatusSeeOther)
				}
			}
		}
	}
}

func Is_admin(next http.HandlerFunc)http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if admin{
			next(w,r)
		}else{
			http.Redirect(w, r, "/user", http.StatusSeeOther)
		}
	}
}