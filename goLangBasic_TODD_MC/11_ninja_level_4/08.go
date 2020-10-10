package main

import "fmt"

func main() {
	m := map[string][]string{
		"james bond": []string{`shaken not stirred`, `martinis`, `women`},
		"chocolate":  []string{"this is a chocolate", "this is a sweet", "this is a new book"},
		"bar":        []string{"this is a  bar", "this is a restaurant", " this is a milk"},
	}

	for k, val := range m {
		fmt.Println("this is a record for", k)
		for i, v2 := range val {

			fmt.Printf("\t  %v  %v \n", i, v2)
		}
	}
	fmt.Println("\n", m)
}
