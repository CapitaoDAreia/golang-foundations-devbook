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

	//Routes
	router.HandleFunc("/users", server.CreateUser).Methods("POST")

	//General configs
	var PORT string = ":80"
	fmt.Printf("Server is listening on port %v\n", PORT)
	log.Fatal(http.ListenAndServe(PORT, router))
}
