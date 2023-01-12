/*
	There are four types of int/float type in Go.
	Those types are classified by the number of bits supported.

	float/int8, float/int16, float/int32, float/int64
*/

package main

import (
	"errors"
	"fmt"
)

func main() {
	var number int8 = 100 // --  This is OK
	// var number int8 = 100000  --  This is not OK
	fmt.Println(number)

	//unsigned integers
	var number2 uint32 = 1000
	fmt.Println(number2)

	//float numbers
	var realNumber float32 = 123.00
	fmt.Println(realNumber)

	//strings
	var str1 string = "Texto"
	fmt.Println(str1)

	//single quatos returns to the var the number of the char in the ascii table
	char := 'A'
	fmt.Println(char)

	//in a unassigned string or a unassigned number variable, the default value is a empty string and zero, consecutively
	//we call that zeroed values
	var str2 string
	var emptyNumber int8
	fmt.Println(str2, emptyNumber)
	//in a boolean the zeroed value is false
	var boolean bool
	fmt.Println(boolean)

	//error type
	//in this case, the zeroed value is <nil>
	//To attribute a value to an error type variable we must use errors package
	var erro error = errors.New("This is a error!")
	fmt.Println(erro)
}
