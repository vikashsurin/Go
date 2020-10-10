package main

import "fmt"

func main() {

	s := []string{"sdf", "e", "df"}
	fmt.Println(&s[0])
	for i := range s {
		fmt.Println(i, &s[0])
		s[0] = "null"
	}
	fmt.Println(s)
}
