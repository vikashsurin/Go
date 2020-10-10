package main

import "fmt"

func main() {
	switch {
	case
		false:
		fmt.Println("this doesnot print")
	case (2 == 3):
		fmt.Println("this will not print")
	case true:
		fmt.Println("this will print true")
	default:
		fmt.Println("this is default")
	}
	val()
}
func val() {
	switch "Bond" {
	case "kala":
		fmt.Println("this is kala")
	case "Bond":
		fmt.Println("this is bond")
	default:
		fmt.Println("thisis default")

	}
}
