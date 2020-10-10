package main

import "fmt"

func main() {
	type person struct {
		first       string
		last        string
		favFlavours []string
	}

	p1 := person{
		first: "James",
		last:  "Bond",
		favFlavours: []string{
			"chocolate",
			"martini",
			"rum and coke",
		},
	}

	p2 := person{
		first: "Nikola",
		last:  "Tesla",
		favFlavours: []string{

			"strawberry",

			"lavander",
			"cola",
		},
	}
	fmt.Println(p1)

	for i, v := range p1.favFlavours {
		fmt.Println("\t", i, v)
	}
	fmt.Println(p2)

	for i, v := range p2.favFlavours {
		fmt.Println("\t", i, v)
	}
}
