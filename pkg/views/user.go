package views

import (
	"html/template"
)

func Userpage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/user.html"))
	return temp
}

func Checkoutpage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/checkout.html"))
	return temp
}

func Checkinpage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/checkin.html"))
	return temp
}


func Issuedbooks() *template.Template {
	temp := template.Must(template.ParseFiles("templates/issuedbooks.html"))
	return temp
}