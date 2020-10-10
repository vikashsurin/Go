package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) speak() {
	fmt.Println("hello i am ", p.name)
}
func (p person) walk() {
	fmt.Println(p.name, " is walking\n")
}

type human interface {
	speak()
	walk()
}

func doSomething(h human) {
	h.speak()
	h.walk()
}
func main() {
	p1 := person{
		name: "James",
		age:  21,
	}
	p2 := person{
		name: "kal",
		age:  22,
	}
	p3 := person{
		name: "laura",
		age:  22,
	}
	p4 := person{
		name: "chici",
		age:  22,
	}
	fmt.Println(p1)
	p1.speak()
	p2.speak()
	p3.speak()
	p4.speak()
	p1.walk()
	p2.walk()
	p3.walk()
	p4.walk()
	doSomething(p1)
	doSomething(p2)
	doSomething(p3)
	doSomething(p4)

	fmt.Printf("%T\t", p1)
}
