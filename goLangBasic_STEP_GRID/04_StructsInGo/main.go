package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo // == contactInfo contactInfo
}

func main() {
	// var alex person
	// alex := person{
	// 	firstName: "Alex",
	// 	lastName:  "Sannator",
	// }

	// alex.firstName = "alex"
	// alex.lastName = "morra"

	// fmt.Println(alex)
	// fmt.Printf("%+v\n", alex)
	// fmt.Printf("%T\n", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 109086,
		},
	}

	fmt.Println(jim)
	fmt.Printf("%T\n%+v \n", jim, jim)

	// jimPointer := &jim
	// jim.updateName("jimmy")
	jim.updateName("Jimmy")
	// jim.updateName("Jimmy")
	jim.print()
	fmt.Println(jim)

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

//receiver for a struct
func (p person) print() {
	fmt.Printf("%+v\n", p)
}
