package views

import (
	"html/template"
)

func Userpage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/user.html"))
	return temp
}