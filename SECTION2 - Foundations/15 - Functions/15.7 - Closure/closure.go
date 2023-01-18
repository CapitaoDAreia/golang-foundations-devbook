package main

import "fmt"

func closure() func() { //This function returns a function that prints a string declared and assigned inside her.
	text := "Closure text"

	function := func() {
		fmt.Println(text)
	}

	return function
}

func main() {
	text := "Main text"
	fmt.Println(text)

	newFunc := closure() //Here, the return is stored inside a variable that will be a function.
	newFunc()            //Here, the function received will be called, and will return the variable received inside her origin function.
}
