package main

import "fmt"

/*
	Using an empty interface, we can make a function that receives every type.
*/
func generic(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generic("String")
	generic(178)
	generic(true)
}
