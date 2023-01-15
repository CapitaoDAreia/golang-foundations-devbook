package main

import (
	"fmt"
)

func modValue(number *int) { //This function receives a pointer to an existent function, that is, receives the reference to an variable.
	*number = *number + 1 //Her this value is modified.
}

func main() {
	numbermain := 10
	fmt.Println(numbermain) //The value.
	modValue(&numbermain)
	fmt.Println(numbermain) //The value modified by reference.
}
