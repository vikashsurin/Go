package main

import "fmt"

func main() {

	fmt.Println("2 + 3 = ", mySum(2, 3))
	fmt.Println("2 + 4 = ", mySum(2, 4))
	fmt.Println("2 + 5 = ", mySum(2, 5))
}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
