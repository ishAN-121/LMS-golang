package controller

import (
	"net/http"
	"strconv"
	"log"

	"LMS/pkg/types"
	"LMS/pkg/views"
	"LMS/pkg/models"
)

func UserPage(response http.ResponseWriter, r *http.Request){
	
	var user types.User
	user.Username = r.Header.Get("username")
	tempelateFunc := views.GetTemplate("userPage")
	t := tempelateFunc()
	t.Execute(response,user)
}

func CheckoutPage(response http.ResponseWriter, r *http.Request){
	db, error := models.Connection()
	if error != nil {
		http.Redirect(response, r, "/serverError", http.StatusFound)
		log.Printf("error %s connecting to the database", error)
	}

	booksList,error := models.GetBooks(db)
	if error != nil {
		http.Redirect(response, r, "/serverError", http.StatusFound)
	}
	
	var message types.Error
	var data types.Data
	data.Books = booksList.Books
	data.Error = message.Msg
	tempelateFunc := views.GetTemplate("checkoutPage")
	t := tempelateFunc()
	t.Execute(response,data)
}

func Checkout(response http.ResponseWriter, r *http.Request){
	var bookRequest types.Request
	bookId_str := r.FormValue("bookIds")
	bookRequest.Username = r.Header.Get("username")

	db, error := models.Connection()
	if error != nil {
		http.Redirect(response, r, "/serverError", http.StatusFound)
		log.Printf("error %s connecting to the database", error)
	}

	var message types.Error
	bookRequest.BookId, _ = strconv.Atoi(bookId_str)
	message,error = models.Checkout(bookRequest.Username,bookRequest.BookId)
	if error != nil {
		http.Redirect(response, r, "/serverError", http.StatusFound)

	}
	booksList,error := models.GetBooks(db)
	if error != nil {
		http.Redirect(response, r, "/serverError", http.StatusFound)
	}
	
	var data types.Data
	data.Books = booksList.Books
	data.Error = message.Msg
	tempelateFunc := views.GetTemplate("checkoutPage")
	t := tempelateFunc()
	t.Execute(response,data)
	
}

func CheckinPage(response http.ResponseWriter, r *http.Request){
	var user types.User
	user.Username = r.Header.Get("username")
	
	booksList,error:= models.IssuedBooks(user.Username)
	if error != nil {
		http.Redirect(response, r, "/serverError", http.StatusFound)
	}
	var message types.Error
	var data types.Data
	data.Books = booksList.Books
	data.Error = message.Msg
	tempelateFunc := views.GetTemplate("checkinPage")
	t := tempelateFunc()
	t.Execute(response,data)
}

func Checkin (response http.ResponseWriter, r *http.Request){
	var bookRequest types.Request
	bookIdstr := r.FormValue("bookIds")
	bookRequest.Username = r.Header.Get("username")

	
	bookRequest.BookId, _ = strconv.Atoi(bookIdstr)
	models.Checkin(bookRequest.Username,bookRequest.BookId)
	CheckinPage(response,r)
}

func IssuedBooks(response http.ResponseWriter, r *http.Request){
	var user types.User
	user.Username = r.Header.Get("username")
	
	booksList,error := models.IssuedBooks(user.Username)
	if error != nil {
		http.Redirect(response, r, "/serverError", http.StatusFound)
	}
	tempelateFunc := views.GetTemplate("issuedBooks")
	t := tempelateFunc()
	t.Execute(response,booksList)
}

func MakeAdminRequest(response http.ResponseWriter, r *http.Request){
	var user types.User
	user.Username = r.Header.Get("username")
	models.AdminRequest(user.Username)
	tempelateFunc := views.GetTemplate("userPage")
	t := tempelateFunc()
	t.Execute(response,user)
}