package main

import "fmt"

func main() {
	//A slice is a datatype that points to an array
	//But, how slice really works? What happens under the clothes?

	//In this example, we use make func to alocate a variable in a space of the memory.
	//The things with slice dinamic(under the clothes) happens like this.
	slice := make([]float32, 5, 5)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	//If i decide to increment the count of elements in a slice,
	//the compiler automatically grows the same slice to support the new amout of elements
	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}
