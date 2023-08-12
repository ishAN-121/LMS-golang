package controller

import (
	"net/http"
	"LMS/pkg/types"
	"LMS/pkg/views"
	"LMS/pkg/models"
	"strconv"
	"log"

)

func AdminPage(response http.ResponseWriter, request *http.Request){

	var user types.User
	user.Username = request.Header.Get("username")
	tempelateFunc := views.GetTemplate("adminPage")
	template := tempelateFunc()
	template.Execute(response,user)
}

func AddNewBookPage(response http.ResponseWriter, request *http.Request){
	
	var error types.Error
	error.Message = ""
	tempelateFunc := views.GetTemplate("addNewBookPage")
	template := tempelateFunc()
	template.Execute(response,error)
}

func UpdateBookPage(response http.ResponseWriter, request *http.Request){
	
	db, error := models.Connection()
	if error != nil {
		http.Redirect(response, request, "/serverError", http.StatusFound)
		log.Printf("error %s connecting to the database", error)
	}
	booksList,error := models.GetBooks(db)
	if error != nil {
		http.Redirect(response, request, "/serverError", http.StatusFound)
	}
	var data types.Data
	data.Books = booksList.Books
	data.Error = ""
	tempelateFunc := views.GetTemplate("updateBookPage")
	template := tempelateFunc()
	template.Execute(response,data)
}


func AddNewBook(response http.ResponseWriter, request *http.Request){
	
	var book types.Book
	var message types.Error
	var error error

	book.Title = request.FormValue("title")
	book.Author = request.FormValue("author")
	Copies_str := request.FormValue("copies")
	

	if (book.Title == "" || book.Author == "" || Copies_str == ""){
		message.Message  = "Invalid Inputs"
	}

	book.Copies,_= strconv.Atoi(Copies_str)
	if (book.Copies < 0 ){
		message.Message  = "Can'template have negative copies"
	}
	if (message.Message == "") {
	message, error = models.AddNewBook(book.Title,book.Author,book.Copies)
	if error != nil {
		http.Redirect(response, request, "/serverError", http.StatusFound)
	}
	}
	tempelateFunc := views.GetTemplate("addNewBookPage")
	template := tempelateFunc()
	template.Execute(response,message)
}


func AddBook(response http.ResponseWriter, request *http.Request){
	
	var book types.Book
	var message types.Error
	var error error

	book.Title = request.FormValue("title")
	book.Author = request.FormValue("author")
	copiesStr := request.FormValue("copies")
	updateBookType := request.FormValue("update")
	

	db, error := models.Connection()
	if error != nil {
		http.Redirect(response, request, "/serverError", http.StatusFound)
		log.Printf("error %s connecting to the database", error)
	}

	if (book.Title == "" || book.Author == "" || copiesStr == ""){
		message.Message  = "Invalid Inputs"
	}
	book.Copies,_ = strconv.Atoi(copiesStr)
	
	if (book.Copies < 0 ){
		message.Message  = "Copies can not be negative"
	}
	if (message.Message == "") {
		if (updateBookType == "add"){
	message,error = models.AddBook(book.Title,book.Author,book.Copies)
	if error != nil {
		http.Redirect(response, request, "/serverError", http.StatusFound)
		}
	}else{
		message,error = models.DeleteBook(book.Title,book.Author,book.Copies)
		if error != nil {
			http.Redirect(response, request, "/serverError", http.StatusFound)
			}
	}
	if error != nil {
		http.Redirect(response, request, "/serverError", http.StatusFound)
	}
	}
	booksList,error := models.GetBooks(db)
	if error != nil {
		http.Redirect(response, request, "/serverError", http.StatusFound)
	}

	var data types.Data
	data.Books = booksList.Books
	data.Error = message.Message
	tempelateFunc := views.GetTemplate("updateBookPage")
	template := tempelateFunc()
	template.Execute(response,data)
}

func AdminCheckout(response http.ResponseWriter, request *http.Request){
	requestedBooks,error := models.RequestedBooks()
	if error != nil {
		http.Redirect(response, request, "/serverError", http.StatusFound)
	}
	var data types.Data
	data.Error = ""
	data.Requests = requestedBooks.Requests
	tempelateFunc := views.GetTemplate("adminCheckout")
	template := tempelateFunc()
	template.Execute(response,data)
}

func ApproveCheckout(response http.ResponseWriter, request *http.Request){
	var bookRequest types.Request
	bookRequest.Id = request.FormValue("request_ids")
	approveType := request.FormValue("approve")
	var data types.Data
	var message types.Error
	var error error
	
	if (approveType == "true"){
	message,error = models.ApproveCheckout(bookRequest.Id)
	}else{
	message,error = models.DenyCheckout(bookRequest.Id)
	}

	if error != nil {
		http.Redirect(response, request, "/serverError", http.StatusFound)
	}
	data.Error = message.Message
	requestedBooks,error := models.RequestedBooks()
	if error != nil {
		http.Redirect(response, request, "/serverError", http.StatusFound)
	}
	data.Requests = requestedBooks.Requests
	tempelateFunc := views.GetTemplate("adminCheckout")
	template := tempelateFunc()
	template.Execute(response,data)
}


func AdminCheckin(response http.ResponseWriter, request *http.Request){
	checkedinBooks,error := models.CheckedinBooks()
	if error != nil {
		http.Redirect(response, request, "/serverError", http.StatusFound)
	}
	tempelateFunc := views.GetTemplate("adminCheckin")
	template := tempelateFunc()
	template.Execute(response,checkedinBooks)

}

func ApproveCheckin(response http.ResponseWriter, request *http.Request){
	var bookRequest types.Request
	bookRequest.Id = request.FormValue("request_ids")
	approveType := request.FormValue("approve")
	if (approveType == "true"){
		models.ApproveCheckin(bookRequest.Id)
		}else{
		models.DenyCheckin(bookRequest.Id)
		}
	AdminCheckin(response,request)
}


func AdminRequest(response http.ResponseWriter, request *http.Request){
	userIds,error := models.AdminRequestUserIds()
	if error != nil {
		http.Redirect(response, request, "/serverError", http.StatusFound)
	}
	tempelateFunc := views.GetTemplate("adminRequest")
	template := tempelateFunc()
	template.Execute(response,userIds)
}

func ApproveAdminRequest(response http.ResponseWriter, request *http.Request){
	var adminRequest types.Request
	adminRequest.Id = request.FormValue("user_ids")
	approveType := request.FormValue("approve")
	if (approveType == "true"){
		models.ApproveAdminRequest(adminRequest.Id)
		}else{
		models.DenyAdminRequest(adminRequest.Id)
		}
	AdminRequest(response,request)
}
