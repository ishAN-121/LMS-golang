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

func Addbookpage(w http.ResponseWriter, r *http.Request){
	t := views.Addbookpage()
	var err types.Error
	err.Msg = ""
	t.Execute(w,err)
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
		error.Msg  = "Can't have negative copies"
	}
	if (error.Msg == "") {
	error = models.Addbook(Title,Author,Copies)
	}
	t := views.Addbookpage()
	t.Execute(w,error)


}

func Deletebookpage(w http.ResponseWriter, r *http.Request){
	t := views.Deletebookpage()
	var err types.Error
	err.Msg = ""
	t.Execute(w,err)
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
		t := views.Deletebookpage()
		t.Execute(w,error)
	
}