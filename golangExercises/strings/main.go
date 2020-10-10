package main

import "fmt"

var string2 string

type string3 string

func main() {
	string1 := "this is a string 1"
	string2 = "this is a string 2"
	string3 := string3("this is a string 3")

	fmt.Println(string1, string2, string3)
	fmt.Printf("%T\t%T\t%T", string1, string2, string3)
	fmt.Println(len(string1), len(string2), len(string3))
}
