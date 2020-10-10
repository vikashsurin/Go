package main

import "fmt"

func main() {
	// x := type{values} // composite literal
	x := []int{3, 45, 33, 6, 9, 8}
	fmt.Println(len(x))

	//access from index position
	fmt.Println(x[3])

	// printing range (index and value)
	for i, v := range x {
		fmt.Println(i, v)
	}
}
