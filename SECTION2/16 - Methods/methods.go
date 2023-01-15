package main

/*
	A method is a function associated with a struct.
*/

type User struct { //Struct
	name    string
	surname string
	age     uint
}

//   Association Funcname Func return
func (user User) whois() (name string) {
	name = user.name + " " + user.surname
	return name
}

func (user User) showAge() (age uint) {
	return user.age
}

func (user *User) changeName(newName string) { //We can use pointers to modify the values of the struct.
	user.name = newName
}

func main() {
	user1 := User{"Igor", "Silva", 24}
	println(user1.whois())
	println(user1.showAge())

	user1.changeName("Lucas")
	println(user1.whois())
}
