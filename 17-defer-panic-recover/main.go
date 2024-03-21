package main

import "fmt"

func main() {
	iterate()
	fmt.Println(division(15, 3))
	fmt.Println(division(15, 0))
	fmt.Println(division(18, 6))
}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from: ", r)
	}
}

func iterate() {
	defer handlePanic()
	for i := 0; i < 10; i++ {
		if i == 4 {
			panic("I hate 4!")
		}
		fmt.Println(i)
	}
}

func division[T int | float64](a, b T) T {
	defer handlePanic()
	if b == 0 {
		panic("Division by zero")
	} else {
		return a / b
	}
}
