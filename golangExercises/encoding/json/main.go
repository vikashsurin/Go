package main

import (
	"encoding/json"
	"fmt"
)


// Person ....
type Person struct {
	Name string
	Age  int32
	Sex  string
}

func main() {
	m := Person{
		Name: "Vikash",
		Age:  23,
		Sex:  "M",
	}
	//Marshal
	// func Marshal(V interface{}) ([]byte, error)
	bs, _ := json.Marshal(m)
	fmt.Println(bs)

	// Unmarshal
	// func Unmarshal(data []byte, V interface{}) error
	var mm Person
	_ = json.Unmarshal(bs, &mm)
	fmt.Println("Unmarshalled :: ", mm)

	var x interface{}
	_ = json.Unmarshal(bs, &x)
	fmt.Println("{}interface :: ", x)
}
