package main

import "fmt"

var x int = 42
var y string = "james Bond"
var z bool = true

func main() {

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	// 03
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)

	fmt.Println(s)

}
