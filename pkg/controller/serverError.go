package controller

import (
	"net/http"
	"LMS/pkg/views"
)

func ServerError(w http.ResponseWriter, r *http.Request){
	tempelateFunc := views.GetTemplate("serverErrorPage")
	t := tempelateFunc()
	t.Execute(w,nil)
}