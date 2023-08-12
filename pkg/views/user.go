package views

import (
	"html/template"
)

func UserPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/user.html"))
	return temp
}

func CheckoutPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/checkout.html"))
	return temp
}

func CheckinPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/checkin.html"))
	return temp
}


func IssuedBooks() *template.Template {
	temp := template.Must(template.ParseFiles("templates/issuedBooks.html"))
	return temp
}