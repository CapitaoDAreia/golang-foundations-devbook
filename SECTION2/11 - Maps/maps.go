package main

import "fmt"

func main() {
	//Declaration adn attribution
	//map[typeof keys]typof values
	user := map[string]string{
		"name":    "Igor",
		"surname": "Silva",
	}

	//Nesting
	user2 := map[string]map[string]string{
		"name": {
			"first": "Igor",
			"last":  "Silva",
		},
		"course": {
			"name":     "Analisys and Dev of Sys",
			"duration": "2,5 years",
		},
	}

	fmt.Println("----- User1 -----")
	fmt.Println(user)
	fmt.Println("----- User2 -----")
	fmt.Println(user2)

	//To delete one key
	fmt.Println("----- User2 modified by delete function -----")
	delete(user2, "name")
	fmt.Println(user2)
}
