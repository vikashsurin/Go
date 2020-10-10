package main

import "fmt"

func main() {

	a := struct {
		name      string
		age       int
		sex       string
		friends   map[string]int //map
		favDrinks []string       //slice
	}{
		name: "Lee Cooper",
		age:  21,
		sex:  "M",
		friends: map[string]int{
			"vegeta": 23,
			"goku":   234,
			"gotens": 45,
			"gohan":  44,
			"chichi": 34443,
			"bulma":  4343,
		},
		favDrinks: []string{
			"cola",
			"lemon juice",
			"ice tea",
		},
	}
	fmt.Println(a.name)
	fmt.Println(a.age)
	fmt.Println(a.sex)
	fmt.Println(a.friends)
	fmt.Println(a.favDrinks)

	for k, v := range a.friends { //  for map
		fmt.Println("\t", k, v)
	}

	for i, val := range a.favDrinks { // for slice
		fmt.Println("\t", i, val)
	}
}
