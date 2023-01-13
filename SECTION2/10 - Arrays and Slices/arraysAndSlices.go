package main

import "fmt"

func main() {
	/*
		ARRAYS
	*/

	//declaration
	var array1 [5]int
	array2 := [5]string{"Um", "Dois", "Três", "Quatro", "Cinco"}

	//attribution
	array1[0] = 1

	//"dinamic" attribution
	//in this way, the count of values that you provide into the '{}' determines the length of the array
	array3 := [...]int{1, 2, 3, 4, 5}

	//PRINT ARRAYS
	fmt.Println("----- PRINT ARRAYS -----")
	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)

	/*
		SLICES - different of the array, a slice has a dynamically-sized structure.
		A slice points to an array, a slice is a piece on an array.
	*/
	//attribution and declaration
	slice1 := []int{10, 11, 12}

	//incrementing values in slice
	slice1 = append(slice1, 13, 14)

	//storing a piece of an array in a slice
	slice2 := array3[1:3] //takes index 1 and 2 of array3

	//PRINT SLICES
	fmt.Println("----- PRINT SLICES -----")
	fmt.Println(slice1)
	fmt.Println(slice2)

}
