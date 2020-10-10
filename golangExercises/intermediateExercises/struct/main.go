package main

import "fmt"

type house struct {
	noRooms bool
	price   int
	city    string
}

func main() {
	house1 := house{
		noRooms: true,
		price:   20,
		city:    "Jamshedpur",
	}
	fmt.Println(house1)

}
