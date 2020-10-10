package main

import "fmt"

func main() {
	x := []int{1, 2, 34, 55, 6}
	fmt.Println(x)
	x = append(x, 23, 445, 22, 44, 2323)
	fmt.Println(x)

	y := []int{23, 33, 454, 222}
	fmt.Println(y)
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
