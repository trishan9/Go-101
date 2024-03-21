package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// regularMap := map[int]int{}
	syncMap := sync.Map{} // Go-routines-Safe Map

	wg.Add(50)
	for i := 0; i < 50; i++ {
		go func() {
			// regularMap[i] = i
			syncMap.Store(i, i)
			wg.Done()
		}()
	}

	wg.Wait()
}
