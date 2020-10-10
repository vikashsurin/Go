package main

import "fmt"

func main() {
	xs1 := []string{"james", "bond", "hero", "kill bill", "marque"}

	xs2 := []string{"mona", "kona", "rona", "lona", "jona"}
	fmt.Println(xs1)
	fmt.Println(xs2)

	xxs := [][]string{xs1, xs2}
	fmt.Println(xxs)

	for i, v := range xxs {
		fmt.Println(i, v)
	}

	// same but not complete

	for i, xs := range xxs {
		fmt.Println("record", i)
		for j, val := range xs {
			fmt.Printf("\t index position: %v \t value: %v \n", j, val)
		}
	}
}
