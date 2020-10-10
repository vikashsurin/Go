package main

import "fmt"

func main() {
	c := make(chan int)

	//send
	go foo(c)

	//receive
	bar(c)

	fmt.Println()
}

//send
func foo(c chan<- int) {
	c <- 43

}

//receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}
