package views

import (
	"html/template"
)

var templateMap = map[string]string{
	"adminPage":       "admin.html",
	"addNewBookPage":  "addNewBook.html",
	"adminCheckout":   "adminCheckout.html",
	"adminCheckin":    "adminCheckin.html",
	"adminRequest":    "adminRequest.html",
	"updateBookPage":  "updateBook.html",
	"booksPage":       "books.html",
	"loginPage":       "login.html",
	"notFoundPage":    "notFoundPage.html",
	"registerPage":    "register.html",
	"serverErrorPage": "serverErrorPage.html",
	"startPage":       "welcome.html",
	"userPage":		   "user.html", 
	"checkoutPage":    "checkout.html",
	"checkinPage":     "checkin.html",
	"issuedBooks":     "issuedBooks.html",
}

func LoadTemplate(templateName string) *template.Template {
	temp := template.Must(template.ParseFiles("templates/" + templateName))
	return temp
}

func GetTemplate(templateName string) func() *template.Template {
	return func() *template.Template {
		return LoadTemplate(templateMap[templateName])
	}
}
