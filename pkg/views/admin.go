package views

import (
	"html/template"
)

func Adminpage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/admin.html"))
	return temp
}