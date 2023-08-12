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

	router.HandleFunc("/",MiddlewareDirect(controller.Welcome)).Methods("GET")
	router.HandleFunc("/register",controller.Register).Methods("GET")
	router.HandleFunc("/login",MiddlewareDirect(controller.LoginPage)).Methods("GET")

	router.HandleFunc("/register",controller.AddUser).Methods("POST")
	router.HandleFunc("/login",MiddlewareDirect(controller.Login)).Methods("POST") 

	//Middleware methods

	router.HandleFunc("/user",Middleware(IsUser(controller.UserPage))).Methods("GET")
	router.HandleFunc("/books",Middleware((controller.Books))).Methods("GET")
	router.HandleFunc("/checkout", Middleware(IsUser(controller.CheckoutPage))).Methods("GET")
	router.HandleFunc("/checkin", Middleware(IsUser(controller.CheckinPage))).Methods("GET")
	router.HandleFunc("/issuedBooks", Middleware(IsUser(controller.IssuedBooks))).Methods("GET")
	router.HandleFunc("/makeAdminRequest", Middleware(IsUser(controller.MakeAdminRequest))).Methods("GET")
	router.HandleFunc("/logout", Middleware(controller.Logout)).Methods("GET")

	router.HandleFunc("/checkout", Middleware(IsUser(controller.Checkout))).Methods("POST")
	router.HandleFunc("/checkin", Middleware(IsUser(controller.Checkin))).Methods("POST")

	//MIddleWare and IsAdmin methods

	router.HandleFunc("/admin",Middleware(IsAdmin((controller.AdminPage)))).Methods("GET")
	router.HandleFunc("/addNewBook",Middleware(IsAdmin((controller.AddNewBookPage)))).Methods("GET")
	router.HandleFunc("/adminCheckout",Middleware(IsAdmin((controller.AdminCheckout)))).Methods("GET")
	router.HandleFunc("/adminCheckin",Middleware(IsAdmin((controller.AdminCheckin)))).Methods("GET")
	router.HandleFunc("/adminRequest",Middleware(IsAdmin((controller.AdminRequest)))).Methods("GET")
	router.HandleFunc("/updateBookPage",Middleware(IsAdmin((controller.UpdateBookPage)))).Methods("GET")
	
	router.HandleFunc("/addNewBook" ,Middleware(IsAdmin((controller.AddNewBook)))).Methods("POST")
	router.HandleFunc("/updateBook",Middleware(IsAdmin((controller.AddBook)))).Methods("POST")
	router.HandleFunc("/approveCheckout",Middleware(IsAdmin((controller.ApproveCheckout)))).Methods("POST")
	router.HandleFunc("/approveCheckin",Middleware(IsAdmin((controller.ApproveCheckin)))).Methods("POST")
	router.HandleFunc("/approveAdminRequest",Middleware(IsAdmin((controller.ApproveAdminRequest)))).Methods("POST")

	router.HandleFunc("/serverError",controller.ServerError).Methods("GET")
	router.NotFoundHandler = http.HandlerFunc(controller.NotFound)

	http.ListenAndServe(":8000", router)

}