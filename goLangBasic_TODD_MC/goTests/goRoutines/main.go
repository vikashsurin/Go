package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go fmt.Println("hello world")
	defer wg.Done()
	wg.Wait()
}
