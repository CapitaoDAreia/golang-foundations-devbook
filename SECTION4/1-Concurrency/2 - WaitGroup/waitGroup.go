package main

import (
	"fmt"
	"sync"
	"time"
)

//WaitGroup is used to wait for all the goroutines.

func writeScreen(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}

func main() {
	var waitGroup sync.WaitGroup //Creates a waitgroup

	waitGroup.Add(2)

	go func() {
		writeScreen("First Function")
		waitGroup.Done() //decrement waitgroup.add
	}()

	go func() {
		writeScreen("Second Function")
		waitGroup.Done() //decrement waitgroup.add
	}()

	waitGroup.Wait() //This statement force execution to wait for the goroutines.

}
