package controller

import (
	"net/http"
	"log"

	"LMS/pkg/models"
	"LMS/pkg/views"
)

func Books(w http.ResponseWriter, r *http.Request) {
	t := views.Books()
	db, err := models.Connection()
	if err != nil {
		log.Printf("error %s connecting to the database", err)
	}
	booksList := models.GetBooks(db)
	t.Execute(w, booksList)
}
