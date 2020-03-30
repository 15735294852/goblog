package main

import (
	"fmt"
	"net/http"
)

func main() {



	http.HandleFunc("/zc",hello)

	http.ListenAndServe(":8080",nil)
	//12


	//p("ChitChat", version(), "started at", config.Address)
	//
	//// handle static assets
	//mux := http.NewServeMux()
	//files := http.FileServer(http.Dir(config.Static))
	//mux.Handle("/static/", http.StripPrefix("/static/", files))
	//
	////
	//// all route patterns matched here
	//// route handler functions defined in other files
	////
	//
	//// index
	//mux.HandleFunc("/", route.index)
	//// error
	//mux.HandleFunc("/err", route.err)
	//
	//// defined in route_auth.go
	//mux.HandleFunc("/login", route.login)
	//mux.HandleFunc("/logout", route.logout)
	//mux.HandleFunc("/signup", route.signup)
	//mux.HandleFunc("/signup_account", route.signupAccount)
	//mux.HandleFunc("/authenticate", route.authenticate)
	//
	//// defined in route_thread.go
	//mux.HandleFunc("/thread/new", route.newThread)
	//mux.HandleFunc("/thread/create", route.createThread)
	//mux.HandleFunc("/thread/post", route.postThread)
	//mux.HandleFunc("/thread/read", route.readThread)
	//
	//// starting up the server
	//server := &http.Server{
	//	Addr:           config.Address,
	//	Handler:        mux,
	//	ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
	//	WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
	//	MaxHeaderBytes: 1 << 20,
	//}
	//server.ListenAndServe()
}

func hello(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w,"Hello Docker Form Golang!")

}
