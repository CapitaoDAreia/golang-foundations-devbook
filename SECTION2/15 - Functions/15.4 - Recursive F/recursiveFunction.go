package main

import "fmt"

//returns the value of one position in fibonacci sequence
func fibonacci(position uint) uint {
	if position <= 1 {
		return position
	}
	return fibonacci(position-2) + fibonacci(position-1)
}

//factorial
func fact(n int) int {
	if n == 0 {
		return 1
	}
	fmt.Println(n * fact(n-1))
	return n * fact(n-1)
}

func main() {
	//Recursive function is a function that, supported by a pile data structure, calls herself.

	fmt.Println(fibonacci(10))
	fmt.Println(fact(10))
}
