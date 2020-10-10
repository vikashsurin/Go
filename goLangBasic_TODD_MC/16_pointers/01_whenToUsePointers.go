package main

import "fmt"

func main() {
	a := 2
	fmt.Println(a)
	fmt.Printf("%T\t\n", a)
	fmt.Println(&a) // gives us the address where the value is stored in a memory.

	b := &a
	fmt.Println(b)
	fmt.Printf("%T\t\n", b)
	fmt.Println(*b) // *gives the values stored in an adress.
	fmt.Println(&b)
}
