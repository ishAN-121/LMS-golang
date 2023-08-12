package controller

import (
	"net/http"
	"LMS/pkg/views"
)

func NotFound(response http.ResponseWriter, request *http.Request){
	tempelateFunc := views.GetTemplate("notFoundPage")
	template := tempelateFunc()
	template.Execute(response,nil)
}