package controller

import (
	"net/http"

	"LMS/pkg/views"
)

func Welcome(response http.ResponseWriter, request *http.Request){
	tempelateFunc := views.GetTemplate("startPage")
	template := tempelateFunc()
	template.Execute(response,nil)
}