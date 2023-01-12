package main

import (
	"fmt"
	"github.com/badoux/checkmail"
	"module/auxiliar"
)

// Write function writes a message from their origin
func main() {
	fmt.Println("Main archive...")
	auxiliar.Write("Write from auxiliar package...")

	error := checkmail.ValidateFormat("danieligord7@hotmail")
	fmt.Println(error)
}
