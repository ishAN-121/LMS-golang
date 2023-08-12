package controller

import (
	"net/http"
	"LMS/pkg/views"
)

func NotFound(w http.ResponseWriter, r *http.Request){
	tempelateFunc := views.GetTemplate("notFoundPage")
	t := tempelateFunc()
	t.Execute(w,nil)
}