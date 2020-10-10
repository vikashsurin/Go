package main

import "fmt"

func main() {

	fmt.Println("helo world")

	//anonymus func
	func(x int) {
		fmt.Println("i am anonymous func", x)
	}(44)
}
