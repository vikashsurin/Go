package main

import "fmt"

func main() {

	defer one()
	two()

}

func one() {
	fmt.Println("first")

}
func two() {
	fmt.Println("second")
}
