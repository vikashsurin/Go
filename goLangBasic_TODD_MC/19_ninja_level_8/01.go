package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Age   int
}

func main() {

	p1 := person{
		First: "James Bond",
		Age:   32,
	}
	p2 := person{
		First: "James Bond",
		Age:   32,
	}

	people := []person{p1, p2}
	// marshal
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
	fmt.Printf("%T", bs)
}
