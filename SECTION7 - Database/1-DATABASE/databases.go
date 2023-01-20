package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //the _ indicates to Go that we need this package, so it can't discarded.
)

func main() {
	stringKeyConnection := "golang:golang@/devbook_golang_course?charset=utf8&parseTime=True&loc=Local" //string with parameters to connect with database

	database, err := sql.Open("mysql", stringKeyConnection) //open connection with database

	if err != nil {
		log.Fatal(err)
	}

	defer database.Close() //defer to close connection with database

	if err := database.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection open.")

	lines, err := database.Query("select * from users") //consult data in a table of the database

	if err != nil {
		log.Fatal(err)
	}
	defer lines.Close() //defer to close lines, wich are a kind of cursor

	fmt.Println(lines)
}
