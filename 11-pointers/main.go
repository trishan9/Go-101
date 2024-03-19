package main

import "fmt"

func updateName(x string) {
	x = "Trishan Dai" // This won't update the original variable -> Pass by value -> Go makes a copy of the value when passes to functions
}

// Used behind the hood for passing by reference in Slices, Maps, and Functions
func updateNameWithPointer(x *string) {
	*x = "Trishan Dai" // This will update the original value as we are changing the value directly in pointer
}

func main() {
	name := "Trishan"
	fmt.Println("My name is ", name)

	updateName(name) // This will not change the value of name
	fmt.Println("My name is ", name)

	addressOfName := &name
	fmt.Println("My name is ", name)
	updateNameWithPointer(addressOfName)
	fmt.Println("My name is ", name) // This will change the value of name
}
