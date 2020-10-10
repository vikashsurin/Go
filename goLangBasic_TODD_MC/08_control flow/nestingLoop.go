package main

import "fmt"

func main() {
	// for init; condition; post {}
	for i := 0; i <= 100; i++ {
		for j := 0; j <= 4; j++ {

			fmt.Println("The outer loop :%d\t The inner loop: %d\n", i, j)
		}

	}

}
