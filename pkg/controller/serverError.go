package controller

import (
	"net/http"
	"LMS/pkg/views"
)

func ServerError(response http.ResponseWriter, request *http.Request){
	tempelateFunc := views.GetTemplate("serverErrorPage")
	template := tempelateFunc()
	template.Execute(response,nil)
}