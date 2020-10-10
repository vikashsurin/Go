package main

import "fmt"

func main() {

	fmt.Println("2 + 3 = ", sum(2, 3))
	fmt.Println("2 + 4 = ", sum(2, 4))
	fmt.Println("2 + 5 = ", sum(2, 5))
}

func sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
