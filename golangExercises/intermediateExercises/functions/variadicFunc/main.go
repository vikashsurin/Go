package main

import "fmt"

func main() {

	students := []string{"ROb", "Jason", "Molly", "jasmine", "Lily"}
	show(students...)
}

func show(ss ...string) {
	for _, v := range ss {
		fmt.Println(v)
	}
}
