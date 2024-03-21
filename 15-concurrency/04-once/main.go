package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	sum := 0
	// var mu sync.Mutex

	// wg.Add(1000)
	// for i := 0; i < 1000; i++ {
	// 	go func() {
	// 		mu.Lock()
	// 		sum++ // Sum is 1000
	// 		mu.Unlock()
	// 		wg.Done()
	// 	}()
	// }

	var once sync.Once
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			once.Do(func() {
				sum++ // Sum is 1
			})
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(sum)
}
