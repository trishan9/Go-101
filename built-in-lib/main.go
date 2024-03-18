package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	yo := "Hello World!!"

	//* Strings are immutable
	fmt.Println(strings.Contains(yo, "Hello"))
	fmt.Println(strings.ReplaceAll(yo, "ll", "mm"))
	fmt.Println(strings.ToUpper(yo))
	fmt.Println(strings.Index(yo, "ld"))
	fmt.Println(strings.Split(yo, ""))
	fmt.Println(yo)

	ages := []int{45, 20, 35, 75, 44, 76, 88}
	fmt.Println(ages)
	sort.Ints(ages)
	fmt.Println(ages) // prints sorted slice

	index := sort.SearchInts(ages, 75)
	fmt.Println(index)

	names := []string{"trishan", "nischay", "abiral", "wagle", "shrestha"}
	sort.Strings(names)
	fmt.Println(names) // prints sorted slice

	fmt.Println(sort.SearchStrings(names, "trishan"))
}
