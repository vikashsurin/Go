package main

import "fmt"

func main() {

	sum := add(2, 3)
	fmt.Println(sum)
}

func add(n1 int, n2 int) int {
	mull := mul(n1, n2)
	return n1 + n2 + mull
}

func mul(n1, n2 int) int {
	return n1 * n2
}
