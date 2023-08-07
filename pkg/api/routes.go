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

	r.HandleFunc("/",controller.Middleware_direct(controller.Welcome)).Methods("GET")
	r.HandleFunc("/register",controller.Register).Methods("GET")
	r.HandleFunc("/login",controller.Middleware_direct(controller.Loginpage)).Methods("GET")
	r.HandleFunc("/user",controller.Middleware(controller.Userpage)).Methods("GET")
	r.HandleFunc("/admin",controller.Middleware(controller.Is_admin((controller.Adminpage)))).Methods("GET")

	r.HandleFunc("/register",controller.Adduser).Methods("POST")
	r.HandleFunc("/login",controller.Middleware_direct(controller.Login)).Methods("POST")
	
	

	http.ListenAndServe(":8000", r)

}