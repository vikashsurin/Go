package main

import "fmt"

/*
type person struct {
	first string
	last  string
	age   int
}
*/

func main() {

	p1 := struct {
		first string
		last  string
	}{
		first: "James",
		last:  "Bond",
	}
	fmt.Println(p1)
}
