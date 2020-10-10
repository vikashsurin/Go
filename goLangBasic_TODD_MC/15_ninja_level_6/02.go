package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := variadic(ii...)
	fmt.Println(n)

	ii2 := []int{1, 2, 3, 4, 5, 3, 6, 7, 32, 12121, 2}
	n2 := varr(ii2)
	fmt.Println(n2)
}

func variadic(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func varr(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
