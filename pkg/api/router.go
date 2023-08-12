package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"LMS/pkg/controller"
)

func Start(){
	router := mux.NewRouter()
	static := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets")))
	router.PathPrefix("/assets/").Handler(static)

	//MiddleWareDirect methods

	router.HandleFunc("/",controller.MiddlewareDirect(controller.Welcome)).Methods("GET")
	router.HandleFunc("/register",controller.Register).Methods("GET")
	router.HandleFunc("/login",controller.MiddlewareDirect(controller.LoginPage)).Methods("GET")

	router.HandleFunc("/register",controller.AddUser).Methods("POST")
	router.HandleFunc("/login",controller.MiddlewareDirect(controller.Login)).Methods("POST") 

	//Middleware methods

	router.HandleFunc("/user",controller.Middleware(controller.IsUser(controller.UserPage))).Methods("GET")
	router.HandleFunc("/books",controller.Middleware((controller.Books))).Methods("GET")
	router.HandleFunc("/checkout", controller.Middleware(controller.IsUser(controller.CheckoutPage))).Methods("GET")
	router.HandleFunc("/checkin", controller.Middleware(controller.IsUser(controller.CheckinPage))).Methods("GET")
	router.HandleFunc("/issuedBooks", controller.Middleware(controller.IsUser(controller.IssuedBooks))).Methods("GET")
	router.HandleFunc("/makeAdminRequest", controller.Middleware(controller.IsUser(controller.MakeAdminRequest))).Methods("GET")
	router.HandleFunc("/logout", controller.Middleware(controller.Logout)).Methods("GET")

	router.HandleFunc("/checkout", controller.Middleware(controller.IsUser(controller.Checkout))).Methods("POST")
	router.HandleFunc("/checkin", controller.Middleware(controller.IsUser(controller.Checkin))).Methods("POST")

	//MIddleWare and IsAdmin methods

	router.HandleFunc("/admin",controller.Middleware(controller.IsAdmin((controller.AdminPage)))).Methods("GET")
	router.HandleFunc("/addNewBook",controller.Middleware(controller.IsAdmin((controller.AddNewBookPage)))).Methods("GET")
	router.HandleFunc("/adminCheckout",controller.Middleware(controller.IsAdmin((controller.AdminCheckout)))).Methods("GET")
	router.HandleFunc("/adminCheckin",controller.Middleware(controller.IsAdmin((controller.AdminCheckin)))).Methods("GET")
	router.HandleFunc("/adminRequest",controller.Middleware(controller.IsAdmin((controller.AdminRequest)))).Methods("GET")
	router.HandleFunc("/updateBookPage",controller.Middleware(controller.IsAdmin((controller.UpdateBookPage)))).Methods("GET")
	
	router.HandleFunc("/addNewBook",controller.Middleware(controller.IsAdmin((controller.AddNewBook)))).Methods("POST")
	router.HandleFunc("/updateBook",controller.Middleware(controller.IsAdmin((controller.AddBook)))).Methods("POST")
	router.HandleFunc("/approveCheckout",controller.Middleware(controller.IsAdmin((controller.ApproveCheckout)))).Methods("POST")
	router.HandleFunc("/approveCheckin",controller.Middleware(controller.IsAdmin((controller.ApproveCheckin)))).Methods("POST")
	router.HandleFunc("/approveAdminRequest",controller.Middleware(controller.IsAdmin((controller.ApproveAdminRequest)))).Methods("POST")

	router.HandleFunc("/serverError",controller.ServerError).Methods("GET")
	router.NotFoundHandler = http.HandlerFunc(controller.NotFound)

	http.ListenAndServe(":8000", router)

}