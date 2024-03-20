package main

import (
	"fmt"
	"sync"
	"time"
)

// func main() {
// 	t0 := time.Now()
// 	for i := range 5 {
// 		mockDbCall(fmt.Sprintf("Hello %v", i))
// 	}
// 	fmt.Printf("Time taken: %v\n", time.Since(t0)) // Took 10s to complete - without using Goroutines
// }

// func mockDbCall(str string) {
// 	time.Sleep(2 * time.Second)
// 	fmt.Println(str)
// }

// var wg = &sync.WaitGroup{}
var wg sync.WaitGroup
var results []string
var mux sync.RWMutex

func main() {
	t0 := time.Now()
	for i := range 5 {
		go mockDbCall(fmt.Sprintf("Hello %v", i))
		wg.Add(1)
	}
	wg.Wait()
	fmt.Printf("Time taken: %v\n", time.Since(t0)) // Took 2s to complete - using Goroutines
	fmt.Printf("Results %v\n", results)
}

func mockDbCall(str string) {
	defer wg.Done()
	time.Sleep(2 * time.Second) // Supposing that this this DB call takes about 2 seconds to complete
	mux.Lock()
	results = append(results, str)
	mux.Unlock()
	fmt.Println(str)
}
