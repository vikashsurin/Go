package main

import "fmt"

// print ascii numbers
func main() {

	for i := 33; i <= 122; i++ {

		fmt.Printf("%v\t%#U\n", i, i)

	}

}
