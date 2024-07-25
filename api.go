package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	store Store
}


// constructor for API service 
func NewAPIServer(addr string, store Store) (*APIServer, error ){
	return &APIServer{addr: addr, store: store}, nil
}

// method to initalize the router
func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	//registering our services 

	log.Print("Starting the server at ", s.addr)

	// log any errors comming from http 
	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}
