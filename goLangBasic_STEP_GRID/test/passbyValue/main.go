package main

import "fmt"

type person struct {
	firstname string
}
type num int
type str string

func main() {

	p1 := person{
		firstname: "Adam",
	}

	//update struct
	p1.firstname = "kola"
	fmt.Println(p1)
	p1.up("alex")
	fmt.Println(p1)

	//update int
	n := num(1)
	fmt.Println(n)
	n.upg(3)
	fmt.Println(n)

	//update string
	string1 := str("stringAlex")
	fmt.Println(string1)
	fmt.Println(&string1)
	// string1 = "dfdf"
	string1 = "thisNewString"
	fmt.Println(string1)
	fmt.Println(&string1)

	string1.upStr("AnotherString")
	fmt.Println(string1)
	fmt.Println(&string1)

}

func (p *person) up(newname string) {
	(*p).firstname = newname
}

func (n *num) upg(newInt int) {
	(*n) = num(newInt)
}

func (s *str) upStr(newString string) {
	*s = str(newString)
}
