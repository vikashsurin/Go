package main

import "fmt"

func main() {

	fmt.Println("helo world")

	// we assing a function to a var
	f := func(x int) {
		fmt.Println("i am anonymous func", x)
	}
	f(123)
}
