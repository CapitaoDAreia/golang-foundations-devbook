package main

import (
	"log"
	"net/http"
)

/*
* // Communication protocol // client-server model // Request - Response //
* // Routes // URI - Unique Resource Identifier // Methods
 */

func main() {

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Users registration."))
	})

	log.Fatal(http.ListenAndServe(":80", nil)) //Create server

}
