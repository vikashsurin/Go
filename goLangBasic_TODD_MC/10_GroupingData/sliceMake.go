package main

import "fmt"

func main() {
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, 2334, 2323, 444, 33, 2222, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 223)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
