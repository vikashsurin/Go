package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

//attaching a speak() method to person
func (p person) speak() {
	fmt.Println("I am ", p.first, p.last, "and I am ", p.age, " years old.")
}
func main() {

	p1 := person{
		first: "james",
		last:  "Bond",
		age:   32,
	}
	p1.speak()
}
