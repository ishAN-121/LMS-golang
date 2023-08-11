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

	//Middleware_direct methods

	r.HandleFunc("/",controller.Middleware_direct(controller.Welcome)).Methods("GET")
	r.HandleFunc("/register",controller.Register).Methods("GET")
	r.HandleFunc("/login",controller.Middleware_direct(controller.Loginpage)).Methods("GET")

	r.HandleFunc("/register",controller.Adduser).Methods("POST")
	r.HandleFunc("/login",controller.Middleware_direct(controller.Login)).Methods("POST") 

	//Middleware methods

	r.HandleFunc("/user",controller.Middleware(controller.Userpage)).Methods("GET")
	r.HandleFunc("/books",controller.Middleware(controller.Books)).Methods("GET")
	r.HandleFunc("/checkout", controller.Middleware(controller.Checkoutpage)).Methods("GET")
	r.HandleFunc("/checkin", controller.Middleware(controller.Checkinpage)).Methods("GET")
	r.HandleFunc("/issuedbooks", controller.Middleware(controller.Issuedbooks)).Methods("GET")
	r.HandleFunc("/makeadminrequest", controller.Middleware(controller.Makeadminrequest)).Methods("GET")
	r.HandleFunc("/logout", controller.Middleware(controller.Logout)).Methods("GET")

	r.HandleFunc("/checkout", controller.Middleware(controller.Checkout)).Methods("POST")
	r.HandleFunc("/checkin", controller.Middleware(controller.Checkin)).Methods("POST")

	//MIddleWare and Is_admin methods

	r.HandleFunc("/admin",controller.Middleware(controller.Is_admin((controller.Adminpage)))).Methods("GET")
	r.HandleFunc("/addnewbook",controller.Middleware(controller.Is_admin((controller.AddNewBookPage)))).Methods("GET")
	r.HandleFunc("/admin-checkout",controller.Middleware(controller.Is_admin((controller.Admincheckout)))).Methods("GET")
	r.HandleFunc("/admin-checkin",controller.Middleware(controller.Is_admin((controller.Admincheckin)))).Methods("GET")
	r.HandleFunc("/adminrequest",controller.Middleware(controller.Is_admin((controller.Adminrequest)))).Methods("GET")
	r.HandleFunc("/adddeletebook",controller.Middleware(controller.Is_admin((controller.AddDeleteBookPage)))).Methods("GET")
	
	r.HandleFunc("/addnewbook",controller.Middleware(controller.Is_admin((controller.AddNewBook)))).Methods("POST")
	r.HandleFunc("/update-add",controller.Middleware(controller.Is_admin((controller.Addbook)))).Methods("POST")
	r.HandleFunc("/update-delete",controller.Middleware(controller.Is_admin((controller.Deletebook)))).Methods("POST")
	r.HandleFunc("/approve-checkout",controller.Middleware(controller.Is_admin((controller.Approvecheckout)))).Methods("POST")
	r.HandleFunc("/deny-checkout",controller.Middleware(controller.Is_admin((controller.Denycheckout)))).Methods("POST")
	r.HandleFunc("/approve-checkin",controller.Middleware(controller.Is_admin((controller.Approvecheckin)))).Methods("POST")
	r.HandleFunc("/deny-checkin",controller.Middleware(controller.Is_admin((controller.Denycheckin)))).Methods("POST")
	r.HandleFunc("/approveadminrequest",controller.Middleware(controller.Is_admin((controller.Approveadminrequest)))).Methods("POST")
	r.HandleFunc("/denyadminrequest",controller.Middleware(controller.Is_admin((controller.Denyadminrequest)))).Methods("POST")

	http.ListenAndServe(":8000", r)

}