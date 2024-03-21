package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	mux := &sync.Mutex{}
	// ch := make(chan int) // Unbuffered channels -> will give deadlock error if there're no any consumer for the channel
	ch := make(chan int, 3) // Buffered channels

	wg.Add(2)

	go func(ch chan int, wg *sync.WaitGroup, mux *sync.Mutex) {
		defer wg.Done()
		// fmt.Println(<-ch) // waits for the message
		// fmt.Println(<-ch)

		for i := 0; i < 3; i++ {
			fmt.Println(<-ch)
		}

		for message := range ch {
			fmt.Println(message) // gives deadlock if channel is not closed
		}
	}(ch, wg, mux)

	go func(ch chan int, wg *sync.WaitGroup, mux *sync.Mutex) {
		defer wg.Done()
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}(ch, wg, mux)

	wg.Wait()
}
