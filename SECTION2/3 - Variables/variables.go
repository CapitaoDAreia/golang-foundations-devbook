package main

import "fmt"

func main() {
	var variable1 string = "variable1" //explicit declaration and attribution
	variable2 := "variable2"           //Inference declaration and attribution
	fmt.Println(variable1, variable2)

	var ( //declaring more than one
		variable3  string = "variable3"
		variable4  string = "variable4"
		variable10 int    = 10
	)
	fmt.Println(variable3, variable4, variable10)

	variable3, variable4 = variable4, variable3 //exchanging values
	fmt.Println(variable3, variable4)

	const constant1 string = "Constant 1"
	fmt.Println(constant1)

}
