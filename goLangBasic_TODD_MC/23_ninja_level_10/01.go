package main

import "fmt"

func main() {
	c := make(chan int, 1)

	c <- 32

	fmt.Println(<-c)
	////////------///////

	cc := make(chan int)

	go func() {
		cc <- 43
	}()
	fmt.Println(<-cc)
}
