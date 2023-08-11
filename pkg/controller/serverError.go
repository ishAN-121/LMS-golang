package controller

import (
	"net/http"
	"LMS/pkg/views"
)

func ServerError(w http.ResponseWriter, r *http.Request){
	t := views.ServerErrorPage()
	t.Execute(w,nil)
}