package main

import (
	"crud/server"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//Mux instance
	router := mux.NewRouter()

	//ROUTES
	//POST
	router.HandleFunc("/users", server.CreateUser).Methods("POST")

	//GET
	router.HandleFunc("/users", server.SearchUsers).Methods("GET")
	router.HandleFunc("/user/{id}", server.SearchUser).Methods("GET")

	//PUT
	router.HandleFunc("/user/{id}", server.UpdateUser).Methods("PUT")

	//General configs
	var PORT string = ":80"
	fmt.Printf("Server is listening on port %v\n", PORT)
	log.Fatal(http.ListenAndServe(PORT, router))
}
