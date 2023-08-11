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
	db, err := models.Connection()
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
		log.Printf("error %s connecting to the database", err)
	}

	booksList,err := models.GetBooks(db)
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
	}
	t := views.Checkoutpage()
	var error types.Error
	var data types.Data
	data.Books = booksList.Books
	data.Error = error.Msg
	t.Execute(w,data)
}

func Checkout(w http.ResponseWriter, r *http.Request){
	bookId_str := r.FormValue("bookIds")
	username := r.Header.Get("username")

	db, err := models.Connection()
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
		log.Printf("error %s connecting to the database", err)
	}

	var msg types.Error
	bookId, _ := strconv.Atoi(bookId_str)
	msg,err = models.Checkout(username,bookId)
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)

	}
	booksList,err := models.GetBooks(db)
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
	}
	
	var data types.Data
	data.Books = booksList.Books
	data.Error = msg.Msg
	t := views.Checkoutpage()
	t.Execute(w,data)
	
}

func Checkinpage(w http.ResponseWriter, r *http.Request){
	username := r.Header.Get("username")
	t := views.Checkinpage()
	booksList,err:= models.Issuedbooks(username)
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
	}
	var msg types.Error
	var data types.Data
	data.Books = booksList.Books
	data.Error = msg.Msg
	t.Execute(w,data)
}

func Checkin (w http.ResponseWriter, r *http.Request){
	bookId_str := r.FormValue("bookIds")
	username := r.Header.Get("username")
	
	bookId, _ := strconv.Atoi(bookId_str)
	models.Checkin(username,bookId)
	Checkinpage(w,r)
}

func Issuedbooks(w http.ResponseWriter, r *http.Request){
	username := r.Header.Get("username")
	t := views.Issuedbooks()
	booksList,err := models.Issuedbooks(username)
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
	}
	t.Execute(w,booksList)
}

func Makeadminrequest(w http.ResponseWriter, r *http.Request){
	var user types.User
	user.Username = r.Header.Get("username")
	models.Adminrequest(user.Username)
	t := views.Userpage()
	t.Execute(w,user)
}