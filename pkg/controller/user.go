package controller

import (
	"net/http"
	"strconv"
	"log"

	"LMS/pkg/types"
	"LMS/pkg/views"
	"LMS/pkg/models"
)

func Userpage(w http.ResponseWriter, r *http.Request){
	t := views.Userpage()
	var user types.User
	user.Username = r.Header.Get("username")
	t.Execute(w,user)
}

func Checkoutpage(w http.ResponseWriter, r *http.Request){
	booksList:= models.Books()
	t := views.Checkoutpage()
	var err types.Error
	var data types.Data
	data.Books = booksList.Books
	data.Error = err.Msg
	t.Execute(w,data)
}

func Checkout(w http.ResponseWriter, r *http.Request){
	bookId_str := r.FormValue("bookIds")
	username := r.Header.Get("username")

	var error types.Error
	bookId, err := strconv.Atoi(bookId_str)
	if err != nil{
		log.Println(err)
	}
	error = models.Checkout(username,bookId)
	booksList:= models.Books()
	
	var data types.Data
	data.Books = booksList.Books
	data.Error = error.Msg
	t := views.Checkoutpage()
	t.Execute(w,data)
	
}

func Checkinpage(w http.ResponseWriter, r *http.Request){
	username := r.Header.Get("username")
	t := views.Checkinpage()
	booksList:= models.Issuedbooks(username)
	var err types.Error
	var data types.Data
	data.Books = booksList.Books
	data.Error = err.Msg
	t.Execute(w,data)
}

func Checkin (w http.ResponseWriter, r *http.Request){
	bookId_str := r.FormValue("bookIds")
	username := r.Header.Get("username")
	
	bookId, err := strconv.Atoi(bookId_str)
	if err != nil{
		log.Println(err)
	}
	models.Checkin(username,bookId)
	Checkinpage(w,r)
}

func Issuedbooks(w http.ResponseWriter, r *http.Request){
	username := r.Header.Get("username")
	t := views.Issuedbooks()
	booksList := models.Issuedbooks(username)
	t.Execute(w,booksList)
}

func Makeadminrequest(w http.ResponseWriter, r *http.Request){
	var user types.User
	user.Username = r.Header.Get("username")
	models.Adminrequest(user.Username)
	t := views.Userpage()
	t.Execute(w,user)
}