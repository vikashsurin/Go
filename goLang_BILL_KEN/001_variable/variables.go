package main

import "fmt"

var a int    //int can change size with respect to architecture.
var b string // a string has  2 values : pointer and length
// |he|l|l|o  <-  pointer
//              length(5)
var c float64
var d bool

func main() {

	var aa int32 = 23
	var bb int32 = 32
	strr := "string"
	fmt.Printf("%T\n%T\n", aa, bb)
	fmt.Printf("%#v\n%#v\n", &aa, &bb)
	fmt.Println(len(strr))
	ss := make([]string, 23)

	fmt.Println(len(ss))
}

// go is all about type
//
