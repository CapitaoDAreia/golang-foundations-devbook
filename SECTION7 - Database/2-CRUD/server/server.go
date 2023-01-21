package server

import (
	"crud/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

// Search Users
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	DB, err := database.Connect()
	if err != nil {
		w.Write([]byte("SearchUsers: Error on connect with database"))
		return
	}
	defer DB.Close()

	rows, err := DB.Query("SELECT * FROM users") //Queries are used for consultation only.
	if err != nil {
		w.Write([]byte("SearchUsers: Error on consult for users."))
		return
	}
	defer rows.Close()

	var users []user
	for rows.Next() { //Iterates over query response (rows).
		var user user

		//Copy the values in rows to the var user declared above. The parameters is the address to copy the values.
		if err := rows.Scan(&user.Id, &user.Name, &user.Email); err != nil {
			w.Write([]byte("SearchUsers: Error on scan for user."))
			return
		}

		users = append(users, user) //Append the user of iteration in the users slice.
	}

	w.WriteHeader(200)

	//Encodes users and send as a response
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("SearchUsers: Error on convert users to JSON."))
		return
	}
}

// Search User
func SearchUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		w.Write([]byte("SearchUser: Error on convert param to int."))
		return
	}

	DB, err := database.Connect()
	if err != nil {
		w.Write([]byte("SearchUser: Error on connect with database."))
		return
	}

	row, err := DB.Query("select * from users where id = ?", ID)
	if err != nil {
		w.Write([]byte("SearchUser: Error on search for user."))
		return
	}

	var user user
	if row.Next() {
		if err := row.Scan(&user.Id, &user.Name, &user.Email); err != nil {
			w.Write([]byte("SearchUser: Error on scan for user."))
			return
		}
	}

	w.WriteHeader(200)

	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte("SearchUser: Error on convert user to JSON."))
		return
	}
}
