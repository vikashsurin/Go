package main

import "fmt"

//goroutine
// func main() {
// 	c := make(chan int)

// 	go func() {
// 		c <- 42
// 	}()
// 	fmt.Println(<-c)
// }

//buffer
// func main() {
// 	c := make(chan int, 1)

// 	c <- 42

// 	fmt.Println(<-c)
// }

//unsuccessful buffer
func main() {

	c := make(chan int, 2)
	c <- 42
	c <- 43
	fmt.Println(<-c)
	fmt.Println(<-c)
}
