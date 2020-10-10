package main

import "fmt"

func main() {

	a, b := vals()
	fmt.Println(a, b)

	c, d := vals2()
	fmt.Println(c, d)
}

func vals() (int, int) {
	return 2, 4
}

func vals2() (string, string) {
	return "this", "okThere"
}
