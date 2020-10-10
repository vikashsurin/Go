package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(x)

	fmt.Println(x[2:5])
	fmt.Println(x[5:7])
	fmt.Println(x[7:])
	fmt.Println(x[:4])

}
