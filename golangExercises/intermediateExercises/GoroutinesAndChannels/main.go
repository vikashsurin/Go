package main

import "fmt"

func main() {

	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i)
			for j := 0; j < 20; j++ {
				fmt.Println("\t\t", j)
			}
		}
		c <- 1 // block code
	}()

	fmt.Println(<-c)
	fmt.Println("hllo")
}
