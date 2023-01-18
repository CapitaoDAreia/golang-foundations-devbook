package main

import "fmt"

type user struct {
	name    string
	surname string
	age     uint8
	address address
}

type address struct {
	address string
}

func main() {
	var user1 user

	user1.name = "Igor"
	user1.surname = "Silva"
	user1.age = 24
	fmt.Println(user1)

	//inference way to do the same above
	user2 := user{"Daniel", "Silva", 24, address{"Rua Altonia"}}
	fmt.Println(user2)

	//that a way to provide just one value
	user3 := user{name: "Lucas"}
	fmt.Println(user3)
}
