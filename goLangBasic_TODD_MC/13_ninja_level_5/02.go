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
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k)
		fmt.Println("\t", v.first)
		fmt.Println("\t", v.last)
		for _, val := range v.favFlavours {
			fmt.Println("\t\t", val)
		}
		fmt.Println("--------------")
	}
}
