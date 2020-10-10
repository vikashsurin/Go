package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	kal()

	wg.Add(2)
	go fo()
	go g()

	wg.Wait()
}

func fo() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i, "i am fo")
	}
	wg.Done()
}

func g() {
	for i := 10; i <= 10; i++ {
		fmt.Println(i, "i am go")
	}
	wg.Done()
}

func kal() {
	wg.Add(2)
	func() {
		for i := 0; i <= 10; i++ {
			fmt.Println(i, "f")
		}
		wg.Done()
	}()

	func() {
		for i := 10; i <= 10; i++ {
			fmt.Println(i, "g")
		}
		wg.Done()
	}()
	wg.Wait()
}
