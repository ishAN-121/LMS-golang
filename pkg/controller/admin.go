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
	booksList := models.Books()
	var data types.Data
	data.Books = booksList.Books
	data.Error = ""
	t.Execute(w,data)

}


func AddNewBook(w http.ResponseWriter, r *http.Request){
	
	Title := r.FormValue("title")
	Author := r.FormValue("author")
	Copies_str := r.FormValue("copies")
	var error types.Error

	if (Title == "" || Author == "" || Copies_str == ""){
		error.Msg  = "Invalid Inputs"
	}
	Copies,err := strconv.Atoi(Copies_str)
	if err != nil{
		log.Println(err)
	}
	
	if (Copies < 0 ){
		error.Msg  = "Can't have negative copies"
	}
	if (error.Msg == "") {
	error = models.AddNewbook(Title,Author,Copies)
	}
	t := views.AddNewbookpage()
	
	t.Execute(w,error)


}
func Addbook(w http.ResponseWriter, r *http.Request){
	
	Title := r.FormValue("title")
	Author := r.FormValue("author")
	Copies_str := r.FormValue("copies")
	var error types.Error

	if (Title == "" || Author == "" || Copies_str == ""){
		error.Msg  = "Invalid Inputs"
	}
	Copies,err := strconv.Atoi(Copies_str)
	if err != nil{
		log.Println(err)
	}
	
	if (Copies < 0 ){
		error.Msg  = "Can't add negative copies"
	}
	if (error.Msg == "") {
	error = models.Addbook(Title,Author,Copies)
	}
	booksList := models.Books()
	var data types.Data
	data.Books = booksList.Books
	data.Error = error.Msg
	t := views.AddDeleteBookPage()
	t.Execute(w,data)


}



func Deletebook(w http.ResponseWriter, r *http.Request){
	Title := r.FormValue("title")
	Author := r.FormValue("author")
	Copies_str := r.FormValue("copies")
	var error types.Error
	if (Title == "" || Author == "" || Copies_str == ""){
		error.Msg  = "Invalid Inputs"
	}
	Copies,err := strconv.Atoi(Copies_str)
	if err != nil{
		log.Println(err)
	}
	
	if (Copies < 0 ){
		error.Msg  = "Can't Delete negative copies"
	}
	if (error.Msg == "") {
		error = models.Deletebook(Title,Author,Copies)
		}
	booksList := models.Books()
	var data types.Data
	data.Books = booksList.Books
	data.Error = error.Msg
	t := views.AddDeleteBookPage()
	t.Execute(w,data)
	
}

func Admincheckout(w http.ResponseWriter, r *http.Request){
	requestedbooks := models.Requestedbooks()
	var data types.Data
	data.Error = ""
	data.Requests = requestedbooks.Requests
	t := views.Admincheckout()
	t.Execute(w,data)
}

func Approvecheckout(w http.ResponseWriter, r *http.Request){
	Id := r.FormValue("requestids")
	var error types.Error
	var data types.Data
	error = models.Approvecheckout(Id)
	data.Error = error.Msg
	requestedbooks := models.Requestedbooks()
	data.Requests = requestedbooks.Requests
	t := views.Admincheckout()
	t.Execute(w,data)
}

func Denycheckout(w http.ResponseWriter, r *http.Request){
	Id := r.FormValue("requestids")
	var error types.Error
	var data types.Data
	error = models.Denycheckout(Id)
	data.Error = error.Msg
	requestedbooks := models.Requestedbooks()
	data.Requests = requestedbooks.Requests
	t := views.Admincheckout()
	t.Execute(w,data)
}

func Admincheckin(w http.ResponseWriter, r *http.Request){
	checkedinbooks := models.Checkedinbooks()
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
	userIds := models.AdminRequestUserIds()
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