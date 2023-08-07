package views

import (
	"html/template"
)

func RegisterPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/register.html"))
	return temp
}