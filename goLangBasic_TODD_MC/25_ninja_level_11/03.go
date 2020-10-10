package main

import "fmt"

type customError struct {
	info string
}
type hotdog int

func (ce customError) Error() string {
	return fmt.Sprintf("Here is the error  :%v ", ce.info)
}
func main() {
	c1 := customError{
		info: "need more coffee",
	}
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	foo(c1)
}

func foo(e error) {
	fmt.Println("foo ran - ", e, "\n", e.(customError).info)
}
