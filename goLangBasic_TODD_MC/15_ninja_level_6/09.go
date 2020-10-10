package main

import "fmt"

// func main() {

// 	loli := func(xi []int) int {
// 		if len(xi) == 0 {
// 			return 0
// 		}
// 		if len(xi) == 1 {
// 			return xi[0]
// 		}
// 		return xi[0] + xi[len(xi)-1]
// 	}
// 	x := jojo(loli, []int{1, 2, 3, 4, 5, 6})
// 	fmt.Println(x)

// }

// func jojo(f func(xi []int) int, ii []int) int {
// 	n := f(ii)
// 	n++
// 	return n
// }
func main() {

	callbckFun := func(x int) int {
		return x
	}

	z := jojo(callbckFun, 2)
	fmt.Println(z)

}

func jojo(f func(x int) int, num int) int {
	n := f(num)
	n++
	return n
}
