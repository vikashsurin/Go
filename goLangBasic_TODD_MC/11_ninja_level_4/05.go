package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(x)

	x = append(x[:3], x[5:]...)
	fmt.Println("this is deleting a slice", x)

	x = append(x, 23)
	fmt.Println(x)

	y := []int{23, 323, 45, 66, 43434, 555}
	fmt.Println(y)

	x = append(x, y...)
	fmt.Println(x)
}
