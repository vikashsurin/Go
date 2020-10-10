package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 22, 11, 990, 465}
	xs := []string{"james", "kola", "aplha", "mojo mojo"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}
