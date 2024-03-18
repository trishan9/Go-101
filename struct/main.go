package main

import "fmt"

// Write a method for a struct representing a stack that pops the top element.
type MyStruct struct {
	stack []int
}

func (s MyStruct) Pop() MyStruct {
	s.stack = s.stack[:len(s.stack)-1]
	return s
}

func main() {
	var integerStruct MyStruct = MyStruct{[]int{1, 2, 3, 4, 5}}
	fmt.Println(integerStruct)
	integerStruct = integerStruct.Pop()
	fmt.Println(integerStruct)
}
