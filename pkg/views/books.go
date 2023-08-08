package views

import (
	"html/template"
)

func Books() *template.Template {
	temp := template.Must(template.ParseFiles("templates/books.html"))
	return temp
}