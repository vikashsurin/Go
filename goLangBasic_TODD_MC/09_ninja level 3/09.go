package main

import "fmt"

func main() {
	favSport := "surfing"
	switch favSport {
	case "swimming":
		fmt.Println("this will not  print")
	case "dancing":
		fmt.Println("this will print")
	case "surfing":
		fmt.Println("this is surfing")
	default:
		fmt.Println("this is default")
	}
}
