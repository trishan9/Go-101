package main

import "fmt"

func main() {
	// For Loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i := range 10 {
		fmt.Println(i)
	}

	for {
		fmt.Println("Infinite Loop")
		break
	}

	// While Loop Implementation
	var i int = 0
	for i < 10 {
		fmt.Println(i)
		i += 1
	}
}
