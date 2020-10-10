package main

import "fmt"

func main() {
	f := xx()
	fmt.Println(f())
}

func xx() func() int {

	return func() int {
		return 42
	}
}
