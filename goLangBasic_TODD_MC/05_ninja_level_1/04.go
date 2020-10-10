package main

import "fmt"

type vik int // int is the underlying type

var x vik

//05.go
var y int

func main() {

	fmt.Println(x)
	fmt.Printf("%T\t", x)
	x = 42
	fmt.Println(x)
	//05.go
	y = int(x)

	fmt.Println(y)
	fmt.Printf("%T\t", y)

}
