package main

import "fmt"

func main() {
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
}
