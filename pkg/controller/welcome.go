package controller

import (
	"net/http"

	"LMS/pkg/views"
)

func Welcome(w http.ResponseWriter, r *http.Request){
	tempelateFunc := views.GetTemplate("startPage")
	t := tempelateFunc()
	t.Execute(w,nil)
}