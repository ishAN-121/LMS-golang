package views

import (
	"html/template"
)

func AdminPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/admin.html"))
	return temp
}

func AddNewBookPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/addNewBook.html"))
	return temp
}


func AdminCheckout() *template.Template {
	temp := template.Must(template.ParseFiles("templates/adminCheckout.html"))
	return temp
}

func AdminCheckin() *template.Template {
	temp := template.Must(template.ParseFiles("templates/adminCheckin.html"))
	return temp
}

func AdminRequest () *template.Template {
	temp := template.Must(template.ParseFiles("templates/adminRequest.html"))
	return temp
}

func UpdateBookPage () *template.Template {
	temp := template.Must(template.ParseFiles("templates/updateBook.html"))
	return temp
}