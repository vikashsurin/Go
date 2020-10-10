package main

import "fmt"

func main() {

	c := make(chan int)

	cr := make(<-chan int)
	cs := make(chan<- int)

	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)
}
