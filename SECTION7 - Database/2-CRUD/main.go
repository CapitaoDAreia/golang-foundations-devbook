package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	var PORT string = "80"
	fmt.Printf("Listening on port %v", PORT)
	log.Fatal(http.ListenAndServe(PORT, router))
}
