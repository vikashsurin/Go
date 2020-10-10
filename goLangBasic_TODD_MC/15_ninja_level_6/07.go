package main

import "fmt"

func main() {

	f := func() {
		for i := 0; i <= 100; i++ {
			fmt.Println(i, "i am anonymous func")
		}
	}
	f()
	x()
	y()
}

var x func() = func() {
	fmt.Println("i am anonymous func")
}
var y = func() {
	fmt.Println("hello")
}
