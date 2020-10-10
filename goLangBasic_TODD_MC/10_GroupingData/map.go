package main

import "fmt"

func main() {
	m := map[string]int{
		"james":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["james"])
	v, ok := m["barnabas"]
	fmt.Println(v)
	fmt.Println(ok)
}
