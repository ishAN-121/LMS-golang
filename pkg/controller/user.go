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
	t := views.Checkoutpage()
	var err types.Error
	t.Execute(w,err)
}

func Checkout(w http.ResponseWriter, r *http.Request){
	bookId_str := r.FormValue("bookId")
	title := r.FormValue("title")
	author := r.FormValue("author")
	username := r.Header.Get("username")

	log.Println(title)

	var error types.Error

	if (title == "" || author == "" || bookId_str == ""){
		error.Msg  = "Invalid Inputs"
	}
	bookId, err := strconv.Atoi(bookId_str)
	if err != nil{
		log.Println(err)
	}
	if (error.Msg == "") {
		error = models.Checkout(title,author,username,bookId)
	}
		t := views.Checkoutpage()
		t.Execute(w,error)
	
}

func Checkinpage(w http.ResponseWriter, r *http.Request){
	t := views.Checkinpage()
	var err types.Error
	t.Execute(w,err)
}

func Checkin(w http.ResponseWriter, r *http.Request){
	bookId_str := r.FormValue("bookId")
	title := r.FormValue("title")
	author := r.FormValue("author")
	username := r.Header.Get("username")

	var error types.Error

	if (title == "" || author == "" || bookId_str == ""){
		error.Msg  = "Invalid Inputs"
	}
	bookId, err := strconv.Atoi(bookId_str)
	if err != nil{
		log.Println(err)
	}
	if (error.Msg == "") {
		error = models.Checkin(title,author,username,bookId)
	}
		t := views.Checkinpage()
		t.Execute(w,error)
}

func Issuedbooks(w http.ResponseWriter, r *http.Request){
	username := r.Header.Get("username")
	t := views.Issuedbooks()
	booksList := models.Issuedbooks(username)
	t.Execute(w,booksList)
}

func Adminrequest(w http.ResponseWriter, r *http.Request){
	var user types.User
	user.Username = r.Header.Get("username")
	models.Adminrequest(user.Username)
	t := views.Userpage()
	t.Execute(w,user)
}