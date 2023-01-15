package main

import (
	"fmt"
	"time"
)

/*
	The panic function servs to handle with errors.
	Could be used when your function falls into an unexpected return.

	Used with the 'defer' clause you can handle and treat the panic.
*/

func recoverExecution() {
	if r := recover(); r != nil {
		time.Sleep(time.Second)
		fmt.Println("The function has dropped in PANIC, recover initialized.")
		fmt.Println("")
		time.Sleep(time.Second)
		fmt.Println("Calling the function again with other grades...")
		fmt.Println("")
		time.Sleep(time.Second)
		fmt.Println(studentAproved(10, 10, 10))
		time.Sleep(time.Second)
		fmt.Println("Success! Ending the program...")
		fmt.Println("")
	}
}

func studentAproved(grades ...int) (returnedValue bool) {
	defer recoverExecution()

	sum := 0
	average := 0

	for _, value := range grades {
		sum += value
	}
	average = sum / len(grades)
	fmt.Println(average, grades, sum)
	if average > 6 {
		returnedValue = true
		return returnedValue
	} else if average < 6 {
		returnedValue = false
		return returnedValue
	}

	panic("AVERAGE EXACTLY EQUALS 6")
}

func main() {
	fmt.Println(studentAproved(6, 6, 6))
}
