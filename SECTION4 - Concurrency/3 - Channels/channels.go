package main

import (
	"fmt"
	"time"
)

/*
	Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
	More refs: https://go.dev/tour/concurrency/2
*/

func writeScreen(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- text //Here we pass text variable for the channel
		time.Sleep(time.Second)
	}
	close(channel)
}

func main() {
	channel := make(chan string) //Channel is the way that the info of goroutines transits.
	go writeScreen("Function1", channel)

	for message := range channel {
		fmt.Println(message)
	}

	fmt.Println("Fim do programa.")
}
