package main

import (
	"errors"
	"fmt"
)

func main() {

	a := 15
	b := 6

	var result, remainder, errorMessage = integerDivision(a, b)

	// Conditional Statements
	if errorMessage != nil {
		fmt.Println(errorMessage.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of integer division is %v", result)
	} else {
		fmt.Printf("The result of integer division is %v with remainder %v", result, remainder)
	}

	fmt.Println()

	// Switch Case
	switch {
	case errorMessage != nil:
		fmt.Println(errorMessage.Error())
	case remainder == 0:
		fmt.Printf("The result of integer division is %v", result)
	default:
		fmt.Printf("The result of integer division is %v with remainder %v", result, remainder)
	}

	fmt.Println()

	// Conditional Switch Case
	switch remainder {
	case 0:
		fmt.Println("The division was exact!")
	case 1, 2:
		fmt.Println("The division was close!")
	default:
		fmt.Println("The division was not close!")
	}
}

func integerDivision(a int, b int) (int, int, error) {
	var errorMessage error
	if b == 0 {
		errorMessage = errors.New("value of b can't be zero or negative")
		return 0, 0, errorMessage
	}

	var result int = a / b
	var remainder int = a % b
	return result, remainder, errorMessage
}
