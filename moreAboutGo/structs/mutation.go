package main

import (
	"fmt"
)

// use of pointers for mutation

type person struct {
	name string
}

type employee struct {
	name string
}

// two methods of initializing variable ;
// we can make use of pointers in both the scenerio;
// three ways to mutate the struct or any other type
var x person = person{name: "vikash"}
var emp *employee = &employee{name: "vikash"}

func main() {
	// using functions with pointer as argument
	fmt.Println("x before ", x)
	fmt.Println("x before ", &x)
	change(&x)
	fmt.Println("x after ", x)
	fmt.Println("x after ", &x)

	//  using functions with : pointer as argument
	fmt.Println("EMPLOYEE")
	fmt.Println(*emp)
	changeEmp(emp)
	fmt.Println(*emp)

	// using methods
	// the receviver should be a pointer :
	// but the struct may or maynot be the pointer.
	x.changePerson()
	fmt.Println(x)

	// create a person with constructor
	// using constructor to create struct is beneficial
	// as verification for nil can be verified.
	var y *person = newPerson("vikash")
	fmt.Println("person created ", y)
}

func change(p *person) {
	fmt.Println(&p)
	*p = person{name: "silky"}
}

func (p *person) changePerson() {
	fmt.Println(&p)
	p.name = "priti"
	fmt.Println(*p)

}

func changeEmp(e *employee) {
	e.name = "priti"
}

// using consttructor to build a struct
// example for person struct
func newPerson(name string) *person {
	return &person{name}
}

