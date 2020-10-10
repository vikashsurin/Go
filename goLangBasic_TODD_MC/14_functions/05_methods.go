package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("i am ", s.first, s.last)
}

func main() {

	sa1 := secretAgent{
		person: person{
			first: "james",
			last:  "bond",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	sa1.speak()
}

/*

A method is a function that can be attached to a instance of a struct.
When a struct{} is passed to a method , the struct is called a (receiver).

then ,every instance of that struct can acess the method attached to the struct.

*/
