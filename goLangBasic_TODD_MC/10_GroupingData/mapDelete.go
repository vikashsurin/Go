package main

import "fmt"

func main() {
	m := map[string]int{
		"james":       32,
		"Money Heist": 33,
	}

	fmt.Println(m)
	delete(m, "james")
	fmt.Println(m)
}
