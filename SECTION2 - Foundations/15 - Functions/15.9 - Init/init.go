package main

import "fmt"

func init() {
	fmt.Println("Função init being executed...")
	//The 'init' function is executed before 'main' function. One 'init' function is allowed per archive. Yes, archive, not package.
}

func main() {
	fmt.Println("Função main being executed...")
}
