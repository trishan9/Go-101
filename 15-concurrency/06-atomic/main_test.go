package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkSumWithMutex(b *testing.B) {
	var wg sync.WaitGroup
	var sum int64 = 0
	var mu sync.Mutex

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			sum++
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkSumWithAtomic(b *testing.B) {
	var wg sync.WaitGroup
	var sum int64 = 0

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			atomic.AddInt64(&sum, 1)
			wg.Done()
		}()
	}

	wg.Wait()
}
