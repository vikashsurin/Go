package main

import "fmt"

func main() {
	x := make([]string, 50, 50)
	x = []string{"albama", "alaska"}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}
