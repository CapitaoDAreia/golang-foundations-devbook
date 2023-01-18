package main

import "fmt"

func main() {
	var variable1 int8 = 10
	var variable2 int8 = 10
	fmt.Println(variable1, variable2)

	pointerVariable1 := &variable1 //this is the way to capture the address of a variable in the memory
	fmt.Println(variable1, pointerVariable1)

	//Note, if i increment the value, the address remains the same
	variable1++
	fmt.Println(variable1, pointerVariable1)

	//Here, i'm takig the value by reference and sharing the real value by pointer
	//the variable3 takes the value related with the memory address saved in *pointerVariable1
	variable3 := *pointerVariable1
	variable1++
	fmt.Println(variable3)

}
