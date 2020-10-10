package main

import "fmt"

func main() {
	// slice
	x := []int{3, 45, 33, 6, 9, 8}
	fmt.Println(x[:])
	fmt.Println(x[1:])
	fmt.Println(x[1:3])

	for i := 0; i <= 5; i++ {
		fmt.Println(i, x[i])
	}
}
