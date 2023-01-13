package main

import "fmt"

type person struct {
	name string
	age  uint8
}

type student struct {
	person
	college string
}

func main() {
	student1 := student{person{"Igor", 24}, "Estácio"}
	fmt.Println(student1)
}
