package controller

import (
	"net/http"
	"LMS/pkg/views"
)

func NotFound(w http.ResponseWriter, r *http.Request){
	t := views.NotFoundPage()
	t.Execute(w,nil)
}