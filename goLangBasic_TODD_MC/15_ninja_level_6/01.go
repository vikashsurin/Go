package main

import "fmt"

func main() {
	a := fooo()
	fmt.Println(a)

	b, c := f()
	fmt.Println(b, c)

	d := fo1(33)
	fmt.Println(d)

	e, f := f2(545, "Johnny")
	fmt.Println(e, f)

}

func fooo() int {
	return 42
}

func f() (int, string) {
	return 133, "george"
}

func fo1(n int) int {
	return n
}

func f2(n int, s string) (int, string) {
	return n, s

}
