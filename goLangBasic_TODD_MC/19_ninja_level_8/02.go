package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"First"`
	Age  int    `json:"Age"`
}

func main() {
	j := `[{"First":"James Bond","Age":32},{"First":"James Bond","Age":32}]`
	bs := []byte(j)

	var people []person
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("json : ", j, "\n")
	fmt.Println("bytes : ", bs, "\n")
	fmt.Println(people)

	for i, person := range people {
		fmt.Println("person #", i)
		fmt.Println(person.Name, person.Age)
	}
}
