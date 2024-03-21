package main

import "fmt"

func main() {
	fmt.Println(sumInt(1, 2, 3))
	fmt.Println(sumFloat64(1.0, 2.3, 3.4))

	fmt.Println(sum[int](1, 2, 3))
	fmt.Println(sum(1.0, 2.3, 3.4)) // Type is inferred automatically

	fmt.Println(min(2.5, 2, 1.1, 3))

	fmt.Println(max(3, 2.9))

	type floatType float64
	var b floatType = 2.9
	fmt.Println(max(3, b))

	user := UserData[string]{
		name: "Trishan",
		id:   "1",
		data: "I am Trishan",
	}
	fmt.Println(user)
}

func sumInt(numbers ...int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func sumFloat64(numbers ...float64) float64 {
	sum := 0.0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func sum[T int | float64](numbers ...T) T {
	var sum T
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func min[T int | float64](numbers ...T) T {
	var min T = numbers[0]
	for _, v := range numbers {
		if v < min {
			min = v
		}
	}
	return min
}

// Type Set
type MaxParams interface {
	int | ~float32 | ~float64
}

func max[T MaxParams](a T, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}

// Using Generics in Maps
type CustomMap[T comparable, V int | string] map[T]V

// Using Generics in Structs
type CustomData interface {
	int | float32 | float64 | string | []rune | []byte
}

type UserData[T CustomData] struct {
	name string
	id   string
	data T
}
