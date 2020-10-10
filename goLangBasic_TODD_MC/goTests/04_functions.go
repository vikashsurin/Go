package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	p1 := person{
		name: "james",
		age:  32,
	}
	foss()
	s, i, b, f := ret()
	f()
	fmt.Println(s, i, b, f)
	fmt.Printf("%T\t", s, i, b, f)

	rf := retFun()
	fmt.Printf("\n%T\t", rf)
	fmt.Println("returned function", rf)
	rf()
	rff := retFunStr()
	rff()
	fmt.Println("retFuncStr :", rff)

	i, b, s, sl, m, st := arg(21, true, "james", []string{"kathrina", "julia", "alphert"}, map[string]int{"one": 1, "two": 2, "three": 3}, p1)
	fmt.Println("arguments :", i, b, s, sl, m, st)

	//callback
	fk := funArg(ff)
	fmt.Println(fk)

}
func ff() int {
	return 43
}

func foss() {
	fmt.Println("hello I am foss")
}

// function with return value
func ret() (string, int, bool, func()) {
	return "human", 43, true, func() { fmt.Println("i am returned") }
}

// return a function from a function
func retFun() func() {
	return func() {
		fmt.Println("i am a returned function")
	}
}
func retFunStr() (kal func()) {
	kal = func() {
		fmt.Println("hello go")
	}
	return
}

//functions taking arguments
//i.e, int, bool, string, slice, map, struct.
func arg(n int, b bool, s string, sl []string, m map[string]int, p person) (int, bool, string, []string, map[string]int, person) {
	return n, b, s, sl, m, p
}

// function taking function as argument or callback
func funArg(f func() int) int {

	return 32
}
