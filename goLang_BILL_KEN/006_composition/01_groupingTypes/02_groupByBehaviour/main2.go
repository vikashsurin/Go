package main

import "fmt"

type animal struct {
	name  string
	sound string
}
type human struct {
	name  string
	sound string
}

type speaker interface {
	speak()
}

func (a animal) speak() {
	fmt.Println(a.name, "say", a.sound)

}
func (h human) speak() {
	fmt.Println(h.name, "say", h.sound)

}
func main2() {
	dog := animal{
		name:  "dog",
		sound: "bark",
	}
	cat := animal{
		name:  "cat",
		sound: "meow",
	}

	man := human{
		name:  "rob pike",
		sound: "talk",
	}

	speakers := []speaker{
		dog,
		cat,
		man,
	}
	for _, speaker := range speakers {
		speaker.speak()
	}
}
