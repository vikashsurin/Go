package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Sayings: []string{
			"Shaken", "not shaken", "lovely", "awesome", "never", "like",
		},
	}
	bs, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		panic(err)
	}
	fmt.Println(string(bs))
}
