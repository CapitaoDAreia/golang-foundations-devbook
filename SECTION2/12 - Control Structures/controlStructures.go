package main

import "fmt"

func main() {

	//simple way
	number := 15

	if number > 15 {
		fmt.Println("Bigger")
	} else {
		fmt.Println("Smaller")
	}

	//match and exit
	if number == 15 {
		fmt.Println("Equal")
		// return
	}

	//create a variable anda validate in the same time
	//this variable is not acessible out of this scope
	//we call that 'if init'
	if anotherNumber := number; anotherNumber > 10 {
		fmt.Println("Another number > 10")
	}

	fmt.Println("End!")

}
