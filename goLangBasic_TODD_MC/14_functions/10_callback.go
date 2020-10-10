package main

import "fmt"

//passing the function as a argument is known as : callback

func main() {

	ii := []int{23, 4, 4, 4, 3, 23, 45, 4, 600}
	s := summ(ii...)
	fmt.Println(" sum of all numbers", s)

	s2 := even(summ, ii...)
	fmt.Println(" sum of even numbers", s2)

	s3 := odd(summ, ii...)
	fmt.Println(" sum of odd numberes ", s3)
}
func summ(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

/////////here is a callback//////////////////////
func even(f func(xi ...int) int, vi ...int) int {
	var yi []int ////slice
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)

}

func odd(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
