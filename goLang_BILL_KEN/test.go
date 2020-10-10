package main

import "fmt"

type animal struct {
	name  string
	sound string
}
type human struct {
	name  string
	sound string
}

func main() {
	x := 5 - 8*(2+1)%6
	fmt.Println(x)
}
