package controller

import (
	"net/http"
	"LMS/pkg/types"
	"LMS/pkg/views"
	"LMS/pkg/models"
	"strconv"
	"log"

)

func AdminPage(w http.ResponseWriter, r *http.Request){
	t := views.AdminPage()
	var user types.User
	user.Username = r.Header.Get("username")
	t.Execute(w,user)
}

func AddNewBookPage(w http.ResponseWriter, r *http.Request){
	t := views.AddNewBookPage()
	var err types.Error
	err.Msg = ""
	t.Execute(w,err)
}

func UpdateBookPage(w http.ResponseWriter, r *http.Request){
	t := views.UpdateBookPage()
	db, err := models.Connection()
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
		log.Printf("error %s connecting to the database", err)
	}
	booksList,err := models.GetBooks(db)
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
	}
	var data types.Data
	data.Books = booksList.Books
	data.Error = ""
	t.Execute(w,data)

}


func AddNewBook(w http.ResponseWriter, r *http.Request){
	
	var book types.Book
	var msg types.Error
	var err error

	book.Title = r.FormValue("title")
	book.Author = r.FormValue("author")
	Copies_str := r.FormValue("copies")
	

	if (book.Title == "" || book.Author == "" || Copies_str == ""){
		msg.Msg  = "Invalid Inputs"
	}

	book.Copies,_= strconv.Atoi(Copies_str)
	if (book.Copies < 0 ){
		msg.Msg  = "Can't have negative copies"
	}
	if (msg.Msg == "") {
	msg, err = models.AddNewBook(book.Title,book.Author,book.Copies)
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
	}
	}
	t := views.AddNewBookPage()
	
	t.Execute(w,msg)
}


func AddBook(w http.ResponseWriter, r *http.Request){
	
	var book types.Book
	var msg types.Error
	var err error

	book.Title = r.FormValue("title")
	book.Author = r.FormValue("author")
	Copies_str := r.FormValue("copies")
	

	db, err := models.Connection()
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
		log.Printf("error %s connecting to the database", err)
	}

	if (book.Title == "" || book.Author == "" || Copies_str == ""){
		msg.Msg  = "Invalid Inputs"
	}
	book.Copies,_ = strconv.Atoi(Copies_str)
	
	if (book.Copies < 0 ){
		msg.Msg  = "Can't add negative copies"
	}
	if (msg.Msg == "") {
	msg,err = models.AddBook(book.Title,book.Author,book.Copies)
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
	}
	}
	booksList,err := models.GetBooks(db)
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
	}

	var data types.Data
	data.Books = booksList.Books
	data.Error = msg.Msg
	t := views.UpdateBookPage()
	t.Execute(w,data)
}



func DeleteBook(w http.ResponseWriter, r *http.Request){

	var book types.Book
	var msg types.Error
	var err error

	book.Title = r.FormValue("title")
	book.Author = r.FormValue("author")
	Copies_str := r.FormValue("copies")

	db, err := models.Connection()
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
		log.Printf("error %s connecting to the database", err)
	}

	if (book.Title == "" || book.Author == "" || Copies_str == ""){
		msg.Msg  = "Invalid Inputs"
	}
	book.Copies,_ = strconv.Atoi(Copies_str)

	if (book.Copies < 0 ){
		msg.Msg  = "Can't Delete negative copies"
	}
	if (msg.Msg == "") {
		msg,err = models.DeleteBook(book.Title,book.Author,book.Copies)
		if err != nil {
			http.Redirect(w, r, "/serverError", http.StatusFound)
			}
		}

	booksList,err := models.GetBooks(db)

	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
	}

	var data types.Data
	data.Books = booksList.Books
	data.Error = msg.Msg
	t := views.UpdateBookPage()
	t.Execute(w,data)
	
}

func AdminCheckout(w http.ResponseWriter, r *http.Request){
	requestedBooks,err := models.RequestedBooks()
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
	}
	var data types.Data
	data.Error = ""
	data.Requests = requestedBooks.Requests
	t := views.AdminCheckout()
	t.Execute(w,data)
}

func ApproveCheckout(w http.ResponseWriter, r *http.Request){
	var request types.Request
	request.Id = r.FormValue("requestIds")
	
	var data types.Data
	msg,err := models.ApproveCheckout(request.Id)
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
	}
	data.Error = msg.Msg
	requestedBooks,err := models.RequestedBooks()
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
	}
	data.Requests = requestedBooks.Requests
	t := views.AdminCheckout()
	t.Execute(w,data)
}

func DenyCheckout(w http.ResponseWriter, r *http.Request){
	var request types.Request
	var data types.Data

	request.Id = r.FormValue("requestIds")
	msg,err := models.DenyCheckout(request.Id)

	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
	}
	data.Error = msg.Msg
	requestedBooks,err := models.RequestedBooks()
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
	}
	data.Requests = requestedBooks.Requests
	t := views.AdminCheckout()
	t.Execute(w,data)
}

func AdminCheckin(w http.ResponseWriter, r *http.Request){
	checkedInBooks,err := models.CheckedinBooks()
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
	}
	t := views.AdminCheckin()
	t.Execute(w,checkedInBooks)

}

func ApproveCheckin(w http.ResponseWriter, r *http.Request){
	var request types.Request
	request.Id = r.FormValue("requestIds")
	models.ApproveCheckin(request.Id)
	AdminCheckin(w,r)
}

func DenyCheckin(w http.ResponseWriter, r *http.Request){
	var request types.Request
	request.Id = r.FormValue("requestIds")
	models.DenyCheckin(request.Id)
	AdminCheckin(w,r)
}

func AdminRequest(w http.ResponseWriter, r *http.Request){
	userIds,err := models.AdminRequestUserIds()
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
	}
	t := views.AdminRequest()
	t.Execute(w,userIds)
}

func ApproveAdminRequest(w http.ResponseWriter, r *http.Request){
	var request types.Request
	request.Id = r.FormValue("userIds")
	models.ApproveAdminRequest(request.Id)
	AdminRequest(w,r)
}

func DenyAdminRequest(w http.ResponseWriter, r *http.Request){
	var request types.Request
	request.Id = r.FormValue("userIds")
	models.DenyAdminRequest(request.Id)
	AdminRequest(w,r)
}