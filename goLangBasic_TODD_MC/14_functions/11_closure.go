package main

import "fmt"

// The idea of being aware of the scope of the code.

func main() {
	a := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println("hello world")
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
