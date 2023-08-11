package views

import (
	"html/template"
)

func NotFoundPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/notFoundPage.html"))
	return temp
}