package views

import (
	"html/template"
)

func ServerErrorPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/serverErrorPage.html"))
	return temp
}