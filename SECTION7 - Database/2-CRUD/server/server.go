package server

import (
	"crud/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type user struct {
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Insert an user in database.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body) //Read the body request
	if err != nil {
		w.WriteHeader(405)
		w.Write([]byte("FAIL!"))
		return
	}

	var user user

	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error to convert user -> struct."))
		return
	}

	DB, err := database.Connect()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error on connect with database..."))
		return
	}

	defer DB.Close()

	statement, err := DB.Prepare("insert into users (name, email) values (?, ?)")
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error on create statement."))
		return
	}

	defer statement.Close()

	insertion, err := statement.Exec(user.Name, user.Email)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error on execute insertion."))
		return
	}

	insertedId, err := insertion.LastInsertId()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error on captura insertedID."))
		return
	}

	w.WriteHeader(201)
	w.Write([]byte(fmt.Sprintf("Success on insert, id %v", insertedId)))
}
