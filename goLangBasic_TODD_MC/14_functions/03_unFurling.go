package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 6, 7, 8, 9}
	fmt.Println(xi)
	x := sum(xi...)
	fmt.Println("The total is ", x)

}

func sum(x ...int) int {

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position, ", i, "we are adding ,", v, "to the total which is now, ", sum)
	}
	return sum
}
