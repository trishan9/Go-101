package main

import (
	"encoding/json"
	"fmt"
)

// Fields in struct must be exported in order to be marshaled.
type Todos struct {
	Id          int      `json:"todoId"`
	Task        string   `json:"task"`
	HiddenField string   `json:"-"`
	Description string   `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}

func main() {
	// Marshal
	todos := []Todos{
		{1, "Go to market", "I am Hidden", "", []string{"home", "daily"}},
		{2, "Learn Java", "I am Hidden", "Learn DSA with Java and Design Patterns also", []string{"cs", "programming", "daily"}},
	}
	todosJson, err := json.MarshalIndent(todos, "", "\t")
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%s\n", todosJson)
	}

	// Unmarshal
	var loadedJson []Todos
	isValidJson := json.Valid(todosJson)
	if isValidJson {
		json.Unmarshal(todosJson, &loadedJson)
		fmt.Printf("%#v", loadedJson)
	}
}
