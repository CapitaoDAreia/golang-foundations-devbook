package main

import "fmt"

func sum(n1 int8, n2 int8) int8 { //specifies parameters, return and their types
	return n1 + n2
}

func otherFuncExample(n1, n2 int16) int16 { //specifies parameters and their types. In this way, only one info to parameters is necessary
	return n1 + n2
}

func mathCalc(n1, n2 int8) (int8, int8) { //that function returns two values
	sum := n1 + n2
	sub := n1 - n2

	return sum, sub
}

func main() {
	//calling outside function
	result := sum(10, 10)
	fmt.Println(result)

	//function stored in a variable
	var function = func(text string) {
		fmt.Println(text)
	}
	function("Function Executed...")

	//storing two values
	result1MathCalc, result2MathCalc := mathCalc(20, 15)
	fmt.Println(result1MathCalc, result2MathCalc)

	//storing one value and discarding one of them
	result3MathCalc, _ := mathCalc(20, 10) //the _ is used for discard a value
	fmt.Println(result3MathCalc)
}
