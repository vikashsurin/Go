package main

import "fmt"

func main() {
	weights := []int{70, 80, 77, 40, 55}
	fmt.Println("Average wight of 5 people :", calAvg(weights))
}

func calAvg(w []int) int {
	var tw int
	for _, w := range w {
		tw = tw + w

	}
	return tw / 5
}
