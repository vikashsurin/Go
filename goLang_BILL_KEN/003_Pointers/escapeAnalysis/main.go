package main

import "fmt"

type user struct {
	name string
}

func main() {
	u1 := creatUserV1()
	u2 := creatUserV2()

	fmt.Println("res", u1, u2)

}

/*
---.There is no constructor , there is factory functions.---
*uses value semantic
*returns a value of type user
*/
//FACTORY FUNCTION
func creatUserV1() user {
	u := user{
		name: "vik",
	}
	fmt.Println("v1 :", &u)
	return u
}
func creatUserV2() *user {
	u := &user{
		name: "vik",
	}
	fmt.Println("v2 :", &u)
	return u
}
