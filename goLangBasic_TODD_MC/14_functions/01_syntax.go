package main

import "fmt"

func main() {

	foo()
	bar("James")
	s1 := woo("Kola Bear")
	fmt.Println(s1)
	x, y := mouse("clever", "John")
	fmt.Println(x)
	fmt.Println(y)
}

// func (r receiver) identifier(parameters) (return(s)) {...}
// side note : parameters -- when we pass
// side note : arguments -- when we call

func foo() {
	fmt.Println("hello i am foo")
}

//everything in go is pass by value
func bar(s string) { // s is an idetifier
	fmt.Println("hello", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from wooo ", "", s)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint("hey ", fn, ln)
	b := true
	return a, b
}
