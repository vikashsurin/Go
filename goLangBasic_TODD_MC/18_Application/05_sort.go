package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{2, 34, 5, 666, 78, 2, 2}

	xs := []string{"James", "Q", "M", "Kola", "Bear"}

	fmt.Println("unsorted", xi)
	sort.Ints(xi) // no return statement
	fmt.Println("sorted", xi)

	fmt.Println("unsorted", xs)
	sort.Strings(xs)
	fmt.Println("sorted", xs)
}
