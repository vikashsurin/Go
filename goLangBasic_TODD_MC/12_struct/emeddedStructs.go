package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk   bool
	first string
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk:   true,
		first: "iamFirst",
	}
	p2 := person{
		first: "Hardy",
		last:  "Show",
		age:   23,
	}
	fmt.Println(p2)
	fmt.Println(p2.first)
	fmt.Println(sa1)
	fmt.Println(sa1.person.first)
	fmt.Println(sa1.first)
	fmt.Println(sa1.ltk)
}
