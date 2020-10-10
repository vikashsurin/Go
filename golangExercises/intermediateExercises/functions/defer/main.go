package main

import "fmt"

func main() {

	defer fmt.Println("Hello")
	defer fmt.Println("!")
	fmt.Println("world")
}
