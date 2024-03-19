package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Printf("The sum of 100, 200 and 550 is %v\n", sumOfNumbers(100, 200, 550))

	for i := range 10 {
		fmt.Println(fibo(i))
	}

	fmt.Printf("The factorial of 4 is %v\n", fact(4))

	var result, remainder, errorMessage = integerDivision(15, 3)
	fmt.Println(result, remainder, errorMessage)
}

// Sum of numbers using Variadic Function and Naked Return
func sumOfNumbers(numbers ...int) (sum int) {
	for _, v := range numbers {
		sum += v
	}
	return
}

// Fibonacci using Recursion
func fibo(n int) int {
	if n <= 1 {
		return n
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}

// Factorial using Recursion
func fact(n int) int {
	if n <= 1 {
		return n
	} else {
		return n * fact(n-1)
	}
}

// Function with multiple returns
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
