package controller

import (
	"net/http"
	"strconv"
	"log"

	"LMS/pkg/types"
	"LMS/pkg/views"
	"LMS/pkg/models"
)

func UserPage(w http.ResponseWriter, r *http.Request){
	t := views.UserPage()
	var user types.User
	user.Username = r.Header.Get("username")
	t.Execute(w,user)
}

func CheckoutPage(w http.ResponseWriter, r *http.Request){
	db, err := models.Connection()
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
		log.Printf("error %s connecting to the database", err)
	}

	booksList,err := models.GetBooks(db)
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
	}
	t := views.CheckoutPage()
	var error types.Error
	var data types.Data
	data.Books = booksList.Books
	data.Error = error.Msg
	t.Execute(w,data)
}

func Checkout(w http.ResponseWriter, r *http.Request){
	var request types.Request
	bookId_str := r.FormValue("bookIds")
	request.Username = r.Header.Get("username")

	db, err := models.Connection()
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
		log.Printf("error %s connecting to the database", err)
	}

	var msg types.Error
	request.BookId, _ = strconv.Atoi(bookId_str)
	msg,err = models.Checkout(request.Username,request.BookId)
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
	t := views.CheckoutPage()
	t.Execute(w,data)
	
}

func CheckinPage(w http.ResponseWriter, r *http.Request){
	var user types.User
	user.Username = r.Header.Get("username")
	
	booksList,err:= models.IssuedBooks(user.Username)
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
	}
	var msg types.Error
	var data types.Data
	data.Books = booksList.Books
	data.Error = msg.Msg
	t := views.CheckinPage()
	t.Execute(w,data)
}

func Checkin (w http.ResponseWriter, r *http.Request){
	var request types.Request
	bookId_str := r.FormValue("bookIds")
	request.Username = r.Header.Get("username")

	
	request.BookId, _ = strconv.Atoi(bookId_str)
	models.Checkin(request.Username,request.BookId)
	CheckinPage(w,r)
}

func IssuedBooks(w http.ResponseWriter, r *http.Request){
	var user types.User
	user.Username = r.Header.Get("username")
	
	booksList,err := models.IssuedBooks(user.Username)
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
	}
	t := views.IssuedBooks()
	t.Execute(w,booksList)
}

func MakeAdminRequest(w http.ResponseWriter, r *http.Request){
	var user types.User
	user.Username = r.Header.Get("username")
	models.AdminRequest(user.Username)
	t := views.UserPage()
	t.Execute(w,user)
}