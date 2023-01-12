package main

import (
	"fmt"
	"module/auxiliar"
)

// Write function writes a message from their origin
func main() {
	fmt.Println("Main archive...")
	auxiliar.Write("Write from auxiliar package...")
}
