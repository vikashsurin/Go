package main

import "fmt"

func main() {
	defer goo()
	gar()
}
func goo() {
	fmt.Println("goo")
}
func gar() {
	fmt.Println("gar")
}
