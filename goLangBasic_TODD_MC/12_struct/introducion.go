package main

// struct = a data structure which allows us to compose values of diffrent types
//composite or aggregate data type

import "fmt"

type person struct {
	first string
	last  string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
	}
	p2 := person{
		first: "Hardy",
		last:  "Show",
	}
	fmt.Println(p1, p2)
	fmt.Println(p1.first, p2.first)
}
