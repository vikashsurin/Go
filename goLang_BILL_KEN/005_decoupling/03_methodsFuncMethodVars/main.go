package main

import "fmt"

type user struct {
	name string
	age  int
}

/*
 value semantic.
 makes the copy of the data and modify its data.
 so the underlying data remains unmodified.
*/
func (u user) changeName() user {
	u.name = "cena"
	return u
}

/*
 Pointer semantic.
points to the original data  and modify its data.
so the underlying data experiences a change.
*/
func (u *user) namo() *user {
	u.name = "halkat"
	return u
}
func main() {

	u1 := user{
		name: "John",
		age:  23,
	}
	fmt.Println(u1) //before modification

	kk := u1.changeName()
	fmt.Println(kk) // value semantic mod*
	fmt.Println(u1) // original data remains same

	h := u1.namo()
	fmt.Printf("%T\n", *h) //pointer semantic mod*
	fmt.Println(u1)        //original data changes.
}
