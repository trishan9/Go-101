package main

import "fmt"

func main() {
	// Stacking Defers
	defer fmt.Println("This will be printed at the end of the function 1")
	defer fmt.Println("This will be printed at the end of the function 2")
	defer fmt.Println("This will be printed at the end of the function 3")

	// Constants, Variables, Data types
	var name string = "Trishan"
	var age uint16 = 18
	var isAdult bool
	a1, b1 := 5, 7
	const PI float32 = 3.14

	if age >= 18 {
		isAdult = true
	}

	greeting := `
	Best Regards

	- Trishan Wagle`

	fmt.Print(a1 + b1)
	fmt.Println("\nValue of PI is", PI)
	fmt.Println("Hello World!", name, isAdult, age, greeting, len(name))
}
