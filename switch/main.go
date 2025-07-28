package main

import "fmt"

func main() {
	day := 3

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("i am just a default")
	}

	month := "january"
	switch month {
	case "january", "February", "March":
		fmt.Println("winter")
	case "april", "may", "june":
		fmt.Println("spring")
	default:
		fmt.Println("i am just a default")
	}

	temperature := 10
	switch {
	case temperature < 0:
		fmt.Println("Freezinggg")
	case temperature >= 5:
		fmt.Println("perfect")
	}


}