package main

import (
	"fmt"
	"time"
)

/*
	Concurrency != Parallelism

	Parallelism - Two or more taskes being executed in the same time.

	Concurrency - Tasks being executed sharing resources, but not exactly in the same time.

*/

func writeScreen(text string) {
	for {
		fmt.Println(text)
		time.Sleep((time.Second))
	}
}

func main() {
	/*
		The 'go' keyword  make the rest of the code run independently of the function that has go keyword.
	*/
	go writeScreen("Ihaa!")
	writeScreen("Ihuuuu!")
}
