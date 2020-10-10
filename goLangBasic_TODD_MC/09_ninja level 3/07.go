package main

import "fmt"

func main() {
	x := "James Bond"
	if x == "money penny" {
		fmt.Println("no")
	} else if x == "James Bond" {
		fmt.Println("James Bond")
	} else {
		fmt.Println("lets go to bed")
	}
}
