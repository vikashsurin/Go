package main

import "fmt"

func main() {

	friends := []string{"rom", "ram", "kola", "ok"}

	for i, v := range friends {
		friends = friends[:2]

		fmt.Printf("%v [%v]\n", i, &v)
	}

	for i, v := range friends {
		// friends = friends[:2]
		// fmt.Println(len(friends))
		// fmt.Println(friends[1])
		fmt.Printf("%v %v \n", i, &v)
	}
}
