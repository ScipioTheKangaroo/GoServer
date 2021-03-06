package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// The new router function creates the router and returns it.
// Having it separate makes for easier flexibility and testing
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/assets/", instead of the absolute route itself
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	return r
}

func main() {
	// The router is now formed by calling the newRouter function.
	r := newRouter()
	err := http.ListenAndServe(":8070", r)
	if err != nil {
		log.Fatal("Web server (HTTP): ", err)
	} else {
		fmt.Println("Server started")
	}
}

//Main Handler functions
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
