package main

import "fmt"

func main() {
	/*
		In a nutshell, an anonymous function is a function without a name
	*/
	const parameter string = "One string"
	anonymousFunc := func(parameter string) (parameterReturned string) {
		parameterReturned = parameter + " modified in the anonymous function."
		return
	}(parameter)

	fmt.Println(anonymousFunc)
}
