package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
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

	// For Loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// While Loop Implementation
	var i int = 0
	for i < 10 {
		fmt.Println(i)
		i += 1
	}

	// Maps
	var myMap2 map[string]string = make(map[string]string)
	fmt.Println(myMap2)

	var myMap map[string]string = map[string]string{"name": "Trishan", "address": "Nepal"}
	var myName, nameExists = myMap["name"]
	myMap["city"] = "Kathmandu"
	if nameExists {
		fmt.Println(myName)
	} else {
		fmt.Println("Value with this key doesn't exist!")
	}
	fmt.Println(myMap["city"])
	delete(myMap, "city")
	fmt.Println(myMap["city"])

	// Loop with Maps
	for key, value := range myMap {
		fmt.Println(key, value)
	}

	// Array
	var myArr [3]string = [3]string{"name", "address"}
	var myArr2 [2]string = [...]string{"name", "address"}
	var myArr3 = [...]any{"apple", "banana", "orange", 1, 2}
	fmt.Printf("%v %v", myArr, myArr[0])
	fmt.Printf("%v", myArr2)
	fmt.Printf("%v", myArr3)
	fmt.Println(myArr3[3:5])

	// Loop in Array
	for _, v := range myArr3 {
		fmt.Println(v)
	}

	// Slice
	var mySlice2 []string = make([]string, 2, 8) // Preallocations are good for performance!
	fmt.Printf("%v", mySlice2)

	var mySlice []string = []string{"name", "address"}
	mySlice = append(mySlice, "trishan")
	mySlice = append(mySlice, []string{"name", "address"}...)
	fmt.Println(mySlice, len(mySlice), cap(mySlice))

	// Loop in Slice
	for _, v := range mySlice {
		fmt.Println(v)
	}

	// Strings
	var myString = "résumé"
	for i, v := range myString {
		fmt.Println(i, v)
	}

	// Runes
	var myRune = []rune("résumé")
	for i, v := range myRune {
		fmt.Println(i, v)
	}

	var strSlice = []string{"t", "r", "i", "s", "h", "a", "n"}

	// Inefficient Way
	var name1 = ""
	for _, v := range strSlice {
		name1 += v
	}
	fmt.Println(name1)

	// Efficient Way
	var strBuilder strings.Builder
	for _, v := range strSlice {
		strBuilder.WriteString(v)
	}
	var newName = strBuilder.String()
	fmt.Println(newName)
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
