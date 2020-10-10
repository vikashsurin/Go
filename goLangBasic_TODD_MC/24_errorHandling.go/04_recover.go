package main

import "fmt"

func main() {
	n, err := fmt.Println("hello")
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}
