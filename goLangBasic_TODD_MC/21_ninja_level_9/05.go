package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var incrementor int64

	gs := 100
	wg.Add(gs)

	for i := 0; i <= gs; i++ {
		go func() {
			fmt.Println("Goroutines ", runtime.NumGoroutine())
			atomic.AddInt64(&incrementor, 1)
			fmt.Println(atomic.LoadInt64(&incrementor))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Goroutines ", runtime.NumGoroutine())
	fmt.Println("end value", incrementor)
}
