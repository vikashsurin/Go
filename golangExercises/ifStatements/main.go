package main

import "fmt"

func main() {

	x := 10
	if x > 0 {
		if x == 10 {
			fmt.Println("stop")
		}
		x = x / 2
	}
	fmt.Println(x)

}
