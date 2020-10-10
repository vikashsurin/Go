package main

import "fmt"

// print even numbers
func main() {
	x := 1
	for {
		x++
		if x > 100 {
			break
		}
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)

	}

}
