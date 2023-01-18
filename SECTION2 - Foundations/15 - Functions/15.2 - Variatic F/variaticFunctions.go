package main

import "fmt"

func sum(numbers ...int) (sum int) {
	/*
		In this case you can provide a slice of values that matches 'int' type
		OBS: Remember, in the second () we have the named retur of this function.
	*/
	for _, number := range numbers {
		sum += number
	}
	return
}

func concatenate(fword string, sword string) (result string) {
	/*
		In this case the same happens, but with other types of variable
	*/
	result = fword + " " + sword
	return
}

func main() {
	sumResult := sum(10, 55, 67, 89, 23, 14, 2, 56)
	fmt.Println(sumResult)
	fmt.Println(concatenate("Igor", "Silva"))
}
