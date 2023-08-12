package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"LMS/pkg/controller"
)

func Start(){
	r := mux.NewRouter()
	s := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets")))
	r.PathPrefix("/assets/").Handler(s)

	//MiddleWareDirect methods

	r.HandleFunc("/",controller.MiddlewareDirect(controller.Welcome)).Methods("GET")
	r.HandleFunc("/register",controller.Register).Methods("GET")
	r.HandleFunc("/login",controller.MiddlewareDirect(controller.LoginPage)).Methods("GET")

	r.HandleFunc("/register",controller.AddUser).Methods("POST")
	r.HandleFunc("/login",controller.MiddlewareDirect(controller.Login)).Methods("POST") 

	//Middleware methods

	r.HandleFunc("/user",controller.Middleware(controller.UserPage)).Methods("GET")
	r.HandleFunc("/books",controller.Middleware(controller.Books)).Methods("GET")
	r.HandleFunc("/checkout", controller.Middleware(controller.CheckoutPage)).Methods("GET")
	r.HandleFunc("/checkin", controller.Middleware(controller.CheckinPage)).Methods("GET")
	r.HandleFunc("/issuedBooks", controller.Middleware(controller.IssuedBooks)).Methods("GET")
	r.HandleFunc("/makeAdminRequest", controller.Middleware(controller.MakeAdminRequest)).Methods("GET")
	r.HandleFunc("/logout", controller.Middleware(controller.Logout)).Methods("GET")

	r.HandleFunc("/checkout", controller.Middleware(controller.Checkout)).Methods("POST")
	r.HandleFunc("/checkin", controller.Middleware(controller.Checkin)).Methods("POST")

	//MIddleWare and IsAdmin methods

	r.HandleFunc("/admin",controller.Middleware(controller.IsAdmin((controller.AdminPage)))).Methods("GET")
	r.HandleFunc("/addNewBook",controller.Middleware(controller.IsAdmin((controller.AddNewBookPage)))).Methods("GET")
	r.HandleFunc("/adminCheckout",controller.Middleware(controller.IsAdmin((controller.AdminCheckout)))).Methods("GET")
	r.HandleFunc("/adminCheckin",controller.Middleware(controller.IsAdmin((controller.AdminCheckin)))).Methods("GET")
	r.HandleFunc("/adminRequest",controller.Middleware(controller.IsAdmin((controller.AdminRequest)))).Methods("GET")
	r.HandleFunc("/updateBookPage",controller.Middleware(controller.IsAdmin((controller.UpdateBookPage)))).Methods("GET")
	
	r.HandleFunc("/addNewBook",controller.Middleware(controller.IsAdmin((controller.AddNewBook)))).Methods("POST")
	r.HandleFunc("/addBook",controller.Middleware(controller.IsAdmin((controller.AddBook)))).Methods("POST")
	r.HandleFunc("/deleteBook",controller.Middleware(controller.IsAdmin((controller.DeleteBook)))).Methods("POST")
	r.HandleFunc("/approveCheckout",controller.Middleware(controller.IsAdmin((controller.ApproveCheckout)))).Methods("POST")
	r.HandleFunc("/denyCheckout",controller.Middleware(controller.IsAdmin((controller.DenyCheckout)))).Methods("POST")
	r.HandleFunc("/approveCheckin",controller.Middleware(controller.IsAdmin((controller.ApproveCheckin)))).Methods("POST")
	r.HandleFunc("/denyCheckin",controller.Middleware(controller.IsAdmin((controller.DenyCheckin)))).Methods("POST")
	r.HandleFunc("/approveAdminRequest",controller.Middleware(controller.IsAdmin((controller.ApproveAdminRequest)))).Methods("POST")
	r.HandleFunc("/denyAdminRequest",controller.Middleware(controller.IsAdmin((controller.DenyAdminRequest)))).Methods("POST")


	r.HandleFunc("/serverError",controller.ServerError).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(controller.NotFound)

	http.ListenAndServe(":3000", r)

}