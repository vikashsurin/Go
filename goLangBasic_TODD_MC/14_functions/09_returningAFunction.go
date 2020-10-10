package main

import "fmt"

func main() {
	s1 := fo() // store returned value
	fmt.Println(s1)

	x := baz()
	fmt.Println(x)

	i := x()
	fmt.Println(i)

}
func fo() string {
	s := "hello world"
	return s
}

func baz() func() int {
	// anonymous function returned
	return func() int {
		return 451
	}
}
