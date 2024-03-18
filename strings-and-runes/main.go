package main

import (
	"fmt"
	"strings"
)

func main() {
	// Strings
	var myString = "résumé"
	for i, v := range myString {
		fmt.Println(i, v)
	}

	// Runes
	var myRune = []rune("résumé")
	for i, v := range myRune {
		fmt.Printf("The value in index %d is %c\n", i, v)
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
