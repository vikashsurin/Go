package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("this will not  print")
	case true:
		fmt.Println("this will print")
	default:
		fmt.Println("this is default")
	}
}
