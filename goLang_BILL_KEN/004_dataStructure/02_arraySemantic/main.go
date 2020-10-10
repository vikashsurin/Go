package main

import "fmt"

type person struct {
	name string
}

func main() {

	u := []string{"rom", "ram"}
	fmt.Println(&u[0], u)
	u = append(u, "kal")
	fmt.Println(&u[0], u)
	foo(u)
}
func foo(u []string) {
	fmt.Println(&u[0], u)
}
