package views

import (
	"html/template"
)

func Adminpage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/admin.html"))
	return temp
}

func Addbookpage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/update-add.html"))
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

