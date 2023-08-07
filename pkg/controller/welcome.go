package controller

import (
	"net/http"

	"LMS/pkg/views"
	
)

func Welcome(w http.ResponseWriter, r *http.Request){
	t := views.StartPage()
	t.Execute(w,nil)
	

}