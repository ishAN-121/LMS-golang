package controller

import (
	"net/http"
	"log"

	"LMS/pkg/models"
	"LMS/pkg/views"
)

func Books(response http.ResponseWriter, request *http.Request) {

	db, error := models.Connection()
	if error != nil {
		http.Redirect(response, request, "/serverError", http.StatusFound)
		log.Printf("error %s connecting to the database", error)
	}
	booksList,error := models.GetBooks(db)
	if error != nil {
		http.Redirect(response, request, "/serverError", http.StatusFound)
	}
	tempelateFunc := views.GetTemplate("booksPage")
	template := tempelateFunc()
	template.Execute(response,booksList)
}
