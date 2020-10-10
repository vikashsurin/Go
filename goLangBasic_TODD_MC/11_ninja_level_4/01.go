package main

import "fmt"

func main() {
	x := [5]int{12, 23, 34, 45, 56}

	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
