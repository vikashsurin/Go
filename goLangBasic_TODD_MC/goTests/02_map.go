package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	m := map[string]int{
		"john": 22,
		"ron":  21,
	}
	//range over map
	for k, val := range m {
		fmt.Println("KEY :", k, "::", "VALUE :", val)
	}

	m2 := map[string]string{
		"p1": "bond",
		"p2": "potter",
	}
	for k, val := range m2 {
		fmt.Println(k, val)
	}

	m3 := map[string]person{
		"p1": person{
			name: "james",
			age:  32,
		},
		"p2": person{
			name: "kal",
			age:  33,
		},
	}
	for k, val := range m3 {
		fmt.Println(k, val)
	}

	m4 := map[string]person{
		"p1": {
			name: "james",
			age:  32,
		},
		"p2": {
			name: "kal",
			age:  33,
		},
	}
	for k, val := range m3 {
		fmt.Println(k, val)
	}

	//MUTATION
	m5 := map[string]person{
		"p1": {
			name: "Sophia",
			age:  32,
		},
		"p2": {
			name: "Lily",
			age:  33,
		},
	}
	// add to map
	m5["p3"] = person{name: "Scarlet", age: 21}
	m5["p4"] = person{name: "Senorita", age: 24}
	fmt.Println(m5)

	//delete from a map
	delete(m5, "p1")

	fmt.Println(m)
	fmt.Println(m2)
	fmt.Println(m3)
	fmt.Println(m4)
	fmt.Println(m5)
}
