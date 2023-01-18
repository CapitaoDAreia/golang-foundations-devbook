package main

import (
	"fmt"
	"time"
)

func example1() {
	//Example1 - usual for structure
	i := 0

	fmt.Println("Example 1")
	for i <= 10 {
		i++
		fmt.Printf("Incrementing I... %v\n", i) //control time, the code "sleeps" for one second
		time.Sleep(time.Second)
	}
}

func example2() {
	//Example2 - classic for structure
	fmt.Println("Example 1")
	for j := 0; j <= 10; j++ {
		fmt.Printf("Incrementing J... %v\n", j)
		time.Sleep(time.Second)
	}
}

func example3() {
	//Example3 - more sofisticated for structure
	fmt.Println("Example 3")
	names := [3]string{"Primeiro", "Segundo", "Terceiro"}

	for index, value := range names {
		fmt.Printf("Index: %v - ", index)
		fmt.Printf("Value: %v\n", value)
	}
}

func example4() {
	//Example4 - iterating over strings
	fmt.Println("Example 4 - Iterating over strings")
	for index, letter := range "WORD" {
		fmt.Printf("Index: %v | ", index)
		fmt.Printf("Letter: %v\n", string(letter)) //the function 'string()' is necessary because without her the return is the position in the ascii table
	}
}

func example5() {
	//Example5 - iterating over map
	fmt.Println("Example 5 - Iterating over map structure")
	user := map[string]string{
		"name":    "Igor",
		"surname": "Silva",
	}
	for key, value := range user {
		fmt.Printf("Key: %v | ", key)
		fmt.Printf("Value: %v | \n", value)
	}
}

func main() {
	example5()
}
