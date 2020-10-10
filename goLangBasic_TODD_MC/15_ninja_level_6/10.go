package main

import "fmt"

func main() {
	ff := kaaa(ff())

	fmt.Println()
}

func kaaa() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
