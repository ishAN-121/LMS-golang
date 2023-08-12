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

	var user types.User
	user.Username = r.Header.Get("username")
	tempelateFunc := views.GetTemplate("adminPage")
	t := tempelateFunc()
	t.Execute(w,user)
}

func AddNewBookPage(w http.ResponseWriter, r *http.Request){
	
	var err types.Error
	err.Msg = ""
	tempelateFunc := views.GetTemplate("addNewBookPage")
	t := tempelateFunc()
	t.Execute(w,err)
}

func UpdateBookPage(w http.ResponseWriter, r *http.Request){
	
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
	tempelateFunc := views.GetTemplate("updateBookPage")
	t := tempelateFunc()
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
	tempelateFunc := views.GetTemplate("addNewBookPage")
	t := tempelateFunc()
	t.Execute(w,msg)
}


func AddBook(w http.ResponseWriter, r *http.Request){
	
	var book types.Book
	var msg types.Error
	var err error

	book.Title = r.FormValue("title")
	book.Author = r.FormValue("author")
	Copies_str := r.FormValue("copies")
	updateBookType := r.FormValue("update")
	

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
		msg.Msg  = "Copies can not be negative"
	}
	if (msg.Msg == "") {
		if (updateBookType == "add"){
	msg,err = models.AddBook(book.Title,book.Author,book.Copies)
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
		}
	}else{
		msg,err = models.DeleteBook(book.Title,book.Author,book.Copies)
		if err != nil {
			http.Redirect(w, r, "/serverError", http.StatusFound)
			}
	}
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
	tempelateFunc := views.GetTemplate("updateBookPage")
	t := tempelateFunc()
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
	tempelateFunc := views.GetTemplate("adminCheckout")
	t := tempelateFunc()
	t.Execute(w,data)
}

func ApproveCheckout(w http.ResponseWriter, r *http.Request){
	var request types.Request
	request.Id = r.FormValue("requestIds")
	approveType := r.FormValue("approve")
	var data types.Data
	var msg types.Error
	var err error
	
	if (approveType == "true"){
	msg,err = models.ApproveCheckout(request.Id)
	}else{
	msg,err = models.DenyCheckout(request.Id)
	}

	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
	}
	data.Error = msg.Msg
	requestedBooks,err := models.RequestedBooks()
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
	}
	data.Requests = requestedBooks.Requests
	tempelateFunc := views.GetTemplate("adminCheckout")
	t := tempelateFunc()
	t.Execute(w,data)
}


func AdminCheckin(w http.ResponseWriter, r *http.Request){
	checkedinBooks,err := models.CheckedinBooks()
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
	}
	tempelateFunc := views.GetTemplate("adminCheckin")
	t := tempelateFunc()
	t.Execute(w,checkedinBooks)

}

func ApproveCheckin(w http.ResponseWriter, r *http.Request){
	var request types.Request
	request.Id = r.FormValue("requestIds")
	approveType := r.FormValue("approve")
	if (approveType == "true"){
		models.ApproveCheckin(request.Id)
		}else{
		models.DenyCheckin(request.Id)
		}
	AdminCheckin(w,r)
}


func AdminRequest(w http.ResponseWriter, r *http.Request){
	userIds,err := models.AdminRequestUserIds()
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
	}
	tempelateFunc := views.GetTemplate("adminRequest")
	t := tempelateFunc()
	t.Execute(w,userIds)
}

func ApproveAdminRequest(w http.ResponseWriter, r *http.Request){
	var request types.Request
	request.Id = r.FormValue("userIds")
	approveType := r.FormValue("approve")
	if (approveType == "true"){
		models.ApproveAdminRequest(request.Id)
		}else{
		models.DenyAdminRequest(request.Id)
		}
	AdminRequest(w,r)
}
