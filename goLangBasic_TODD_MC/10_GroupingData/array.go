package main

import "fmt"

// composite data type or aggregate data type
func main() {
	var x [5]int
	fmt.Println(x)
	x[3] = 43
	fmt.Println(x)
	fmt.Println(len(x))
}

// array types
