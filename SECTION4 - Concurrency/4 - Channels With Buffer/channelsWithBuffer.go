package main

import "fmt"

//Sends to a buffered channel block only when the buffer is full.
//https://go.dev/tour/concurrency/3

func main() {
	channel := make(chan string, 2)

	channel <- "Function 1"
	channel <- "Function 2"

	message := <-channel
	message2 := <-channel

	fmt.Println(message, " - ", message2)
}
