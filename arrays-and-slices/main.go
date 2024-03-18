package main

import "fmt"

func main() {
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

	var sliceOfIntegers []int = []int{1, 2, 3, 4, 5}
	fmt.Printf("Sum of all integers is %v\n", calculateSum((sliceOfIntegers)))

	var arrOfIntegers [5]int = [5]int{100, 20, 399, 4, 5}
	fmt.Printf("Greatest of all integers is %v\n", findMax((arrOfIntegers)))
}

// Write a function that takes a slice of integers and returns the sum of all elements.
func calculateSum(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

// Create a program that finds the maximum element in an array of integers.
func findMax(arr [5]int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}
