package main

import "fmt"

type person struct {
	firstname string
	lastname  string
}

type agent struct {
	firstname string
	lastname  string
}

type speak interface {
	greeting()
}

func main() {
	p1 := person{
		firstname: "alex",
		lastname:  "Robber",
	}

	a1 := agent{
		firstname: "Rob",
		lastname:  "Pike",
	}

	fmt.Println(p1, a1)
	// p1.greeting()
	// a1.greeting()
	sayHello(p1)
	sayHello(a1)

}
func sayHello(s speak) {
	// fmt.Println(s.greeting())
	s.greeting()
}

func (p person) greeting() {
	fmt.Println("hello moto :", p.firstname)
	// return "no"
}

func (a agent) greeting() {
	fmt.Println("Hello mojo :", a.firstname)
	// return "ok"
}
