package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

/*
* // Communication protocol // client-server model // Request - Response //
* // Routes // URI - Unique Resource Identifier // Methods
 */

var templates *template.Template

type user struct {
	Name  string
	Email string
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	usr := user{
		"Igor",
		"igor.email@email.com",
	}

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "home.html", usr)
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "users.html", usr)
	})

	fmt.Println("Server listen on PORT 80...")
	log.Fatal(http.ListenAndServe(":80", nil)) //Create server

}
