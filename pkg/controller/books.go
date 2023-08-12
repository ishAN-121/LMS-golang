package controller

import (
	"net/http"
	"log"

	"LMS/pkg/models"
	"LMS/pkg/views"
)

func Books(w http.ResponseWriter, r *http.Request) {

	db, err := models.Connection()
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
		log.Printf("error %s connecting to the database", err)
	}
	booksList,err := models.GetBooks(db)
	if err != nil {
		http.Redirect(w, r, "/serverError", http.StatusFound)
	}
	tempelateFunc := views.GetTemplate("booksPage")
	t := tempelateFunc()
	t.Execute(w,booksList)
}
