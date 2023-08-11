package controller

import (
	"net/http"
	"LMS/pkg/types"
	"LMS/pkg/views"
	"LMS/pkg/models"
	"strconv"
	"log"

)

func Adminpage(w http.ResponseWriter, r *http.Request){
	t := views.Adminpage()
	var user types.User
	user.Username = r.Header.Get("username")
	t.Execute(w,user)
}

func AddNewBookPage(w http.ResponseWriter, r *http.Request){
	t := views.AddNewbookpage()
	var err types.Error
	err.Msg = ""
	t.Execute(w,err)
}

func AddDeleteBookPage(w http.ResponseWriter, r *http.Request){
	t := views.AddDeleteBookPage()
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
	
	Title := r.FormValue("title")
	Author := r.FormValue("author")
	Copies_str := r.FormValue("copies")
	var msg types.Error
	var err error

	if (Title == "" || Author == "" || Copies_str == ""){
		msg.Msg  = "Invalid Inputs"
	}
	Copies,_:= strconv.Atoi(Copies_str)
	if (Copies < 0 ){
		msg.Msg  = "Can't have negative copies"
	}
	if (msg.Msg == "") {
	msg, err = models.AddNewbook(Title,Author,Copies)
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
	}
	}
	t := views.AddNewbookpage()
	
	t.Execute(w,msg)
}
func Addbook(w http.ResponseWriter, r *http.Request){
	
	Title := r.FormValue("title")
	Author := r.FormValue("author")
	Copies_str := r.FormValue("copies")
	var msg types.Error
	var err error

	db, err := models.Connection()
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
		log.Printf("error %s connecting to the database", err)
	}

	if (Title == "" || Author == "" || Copies_str == ""){
		msg.Msg  = "Invalid Inputs"
	}
	Copies,_ := strconv.Atoi(Copies_str)
	
	if (Copies < 0 ){
		msg.Msg  = "Can't add negative copies"
	}
	if (msg.Msg == "") {
	msg,err = models.Addbook(Title,Author,Copies)
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
	t := views.AddDeleteBookPage()
	t.Execute(w,data)
}



func Deletebook(w http.ResponseWriter, r *http.Request){
	Title := r.FormValue("title")
	Author := r.FormValue("author")
	Copies_str := r.FormValue("copies")
	var msg types.Error

	db, err := models.Connection()
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
		log.Printf("error %s connecting to the database", err)
	}

	if (Title == "" || Author == "" || Copies_str == ""){
		msg.Msg  = "Invalid Inputs"
	}
	Copies,_ := strconv.Atoi(Copies_str)

	if (Copies < 0 ){
		msg.Msg  = "Can't Delete negative copies"
	}
	if (msg.Msg == "") {
		msg,err = models.Deletebook(Title,Author,Copies)
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
	t := views.AddDeleteBookPage()
	t.Execute(w,data)
	
}

func Admincheckout(w http.ResponseWriter, r *http.Request){
	requestedbooks,err := models.Requestedbooks()
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
	}
	var data types.Data
	data.Error = ""
	data.Requests = requestedbooks.Requests
	t := views.Admincheckout()
	t.Execute(w,data)
}

func Approvecheckout(w http.ResponseWriter, r *http.Request){
	Id := r.FormValue("requestids")
	
	var data types.Data
	msg,err := models.Approvecheckout(Id)
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
	}
	data.Error = msg.Msg
	requestedbooks,err := models.Requestedbooks()
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
	}
	data.Requests = requestedbooks.Requests
	t := views.Admincheckout()
	t.Execute(w,data)
}

func Denycheckout(w http.ResponseWriter, r *http.Request){
	Id := r.FormValue("requestids")
	var data types.Data
	msg,err := models.Denycheckout(Id)
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
	}
	data.Error = msg.Msg
	requestedbooks,err := models.Requestedbooks()
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
	}
	data.Requests = requestedbooks.Requests
	t := views.Admincheckout()
	t.Execute(w,data)
}

func Admincheckin(w http.ResponseWriter, r *http.Request){
	checkedinbooks,err := models.Checkedinbooks()
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
	}
	t := views.Admincheckin()
	t.Execute(w,checkedinbooks)

}

func Approvecheckin(w http.ResponseWriter, r *http.Request){
	Id := r.FormValue("requestids")
	models.Approvecheckin(Id)
	Admincheckin(w,r)
}

func Denycheckin(w http.ResponseWriter, r *http.Request){
	Id := r.FormValue("requestids")
	models.Denycheckin(Id)
	Admincheckin(w,r)
}

func Adminrequest(w http.ResponseWriter, r *http.Request){
	userIds,err := models.AdminRequestUserIds()
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
	}
	t := views.Adminrequest()
	t.Execute(w,userIds)
}

func Approveadminrequest(w http.ResponseWriter, r *http.Request){
	Id := r.FormValue("userids")
	models.Approveadminrequest(Id)
	Adminrequest(w,r)
}

func Denyadminrequest(w http.ResponseWriter, r *http.Request){
	Id := r.FormValue("userids")
	models.Denyadminrequest(Id)
	Adminrequest(w,r)
}