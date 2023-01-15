package main

import "fmt"

func alunoAprovado(n1 int, n2 int) bool {
	fmt.Println("----- FIRST EXAMPLE -----")
	defer fmt.Println("Média calculada. Retornando resultado...") //this snippet will be executed exactly before any return

	fmt.Println("Calculando média...")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true //before this
	}
	return false //and before that
}

func anotherExample() {
	fmt.Println("----- SECOND EXAMPLE -----")
	fmt.Println("Function1")
	defer fmt.Println("Function2") //this snippet will be the last executed
	fmt.Println("Function3")
	fmt.Println("Function4")
}

func main() {
	fmt.Println(alunoAprovado(1, 6))
	anotherExample()
}
