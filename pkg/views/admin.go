package views

import (
	"html/template"
)

func Adminpage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/admin.html"))
	return temp
}

func AddNewbookpage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/addNewBook.html"))
	return temp
}

func Deletebookpage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/update-delete.html"))
	return temp
}

func Admincheckout() *template.Template {
	temp := template.Must(template.ParseFiles("templates/admin-checkout.html"))
	return temp
}

func Admincheckin() *template.Template {
	temp := template.Must(template.ParseFiles("templates/admin-checkin.html"))
	return temp
}

func Adminrequest () *template.Template {
	temp := template.Must(template.ParseFiles("templates/adminrequest.html"))
	return temp
}

func AddDeleteBookPage () *template.Template {
	temp := template.Must(template.ParseFiles("templates/addDeleteBook.html"))
	return temp
}