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
	fmt.Println("i am ", s.first, s.last, " - the secretAgent speak")
}
func (p person) speak() {
	fmt.Println("I am ", p.first, p.last, " - the person speak")
}

type human interface {
	speak()
}

func bal(h human) {
	fmt.Println("I was passed into bal ", h)
}

type hotdog int

func main() {

	sa1 := secretAgent{
		person: person{
			first: "james",
			last:  "bond",
		},
		ltk: true,
	}

	p1 := person{

		first: "kola",
		last:  "Bear",
	}

	// fmt.Println(sa1)
	sa1.speak()
	p1.speak()
	// fmt.Println(p1)
	bal(sa1)
	bal(p1)
	fmt.Printf("%T %T \n", sa1, p1)

	// conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}

/*

An interface is a Type.
Interface is useful when there is a common method , being used by many diffrent struct.

so interface can also be concluded as a place for common methods. and i can have as many methods ....


*/
