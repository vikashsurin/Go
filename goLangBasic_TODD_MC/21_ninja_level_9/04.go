package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	incrementor := 0
	gs := 100
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i <= gs; i++ {
		go func() {
			fmt.Println("Goroutines ", runtime.NumGoroutine())
			mu.Lock()
			v := incrementor
			fmt.Println("before", incrementor)
			runtime.Gosched()
			v++
			incrementor = v
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Goroutines ", runtime.NumGoroutine())
	fmt.Println("after", incrementor)
}
