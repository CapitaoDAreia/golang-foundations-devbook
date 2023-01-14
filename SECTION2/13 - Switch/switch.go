package main

import "fmt"

func dayOfTheWeek(num uint8) string {

	//choose a variable before
	switch num {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return "Inválido"
	}
}

func dayOfTheWeek2(num uint8) string {

	//choose a variable in time of validation
	switch {
	case num == 1:
		return "Domingo"
	case num == 2:
		return "Segunda-Feira"
	case num == 3:
		return "Terça-Feira"
	case num == 4:
		return "Quarta-Feira"
	case num == 5:
		return "Quinta-Feira"
	case num == 6:
		return "Sexta-Feira"
	case num == 7:
		return "Sábado"
	case num >= 8:
		return fmt.Sprintf("Number %v is invalid.", num)
	default:
		return "Invalid!"
	}
}

func main() {

	day := dayOfTheWeek2(19)
	fmt.Println(day)
}
