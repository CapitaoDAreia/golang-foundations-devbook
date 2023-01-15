package main

import "fmt"

func mathCalcs(x int, y int) (sum int, sub int) {
	/*
		In this case, the firs () refers to parameters
		and the second () refers to the name of the return.

		The variables that the function return are declared in the second ()
		and just assigned in the function.

		The returnm in this case, is implicit.
	*/
	sum = x + y
	if x > y {
		sub = x - y
	} else {
		sub = y - x
	}
	return
}

func main() {
	fmt.Println(mathCalcs(15, 10))
}
