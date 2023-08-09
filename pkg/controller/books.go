package controller

import (
	"net/http"

	"LMS/pkg/models"
	"LMS/pkg/views"
)

func Books(w http.ResponseWriter, r *http.Request) {
	t := views.Books()
	booksList := models.Books()
	t.Execute(w, booksList)
}
