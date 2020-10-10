package main

import "fmt"

func main() {
	//array
	n := [10]int{1, 2, 3, 4, 5, 6, 7}

	//slice
	var s []int = n[1:10]

	fmt.Println(n)
	fmt.Printf("%T\t", s)
	fmt.Println(s)

	names := [4]string{
		"John",
		"Molly",
		"katrina",
		"Sophia",
	}
	fmt.Printf("%T\t", names)
	fmt.Println(names)

	a := names[0:2]
	fmt.Println(a)
	b := names[1:3]
	fmt.Println(b)

	// a slice doesnot stores any value.
	// it is a reference of an array , and can affect the array when slice is modified
	b[0] = "xxx"
	a[0] = "xxxx"
	fmt.Println(a, b)
	fmt.Println(names)

	//slice literals
	i := []int{1, 2, 3, 4, 5, 6}

	bo := []bool{true, false, true, false, false}

	str := []string{"book", "shelf", "car", "ice", "bike"}
	fmt.Println(i, bo, str)

	// slice of struct
	ss := []struct {
		i int
		s string
		b bool
	}{{1, "joker", true}, {2, "venom", false}}

	fmt.Printf("%T\t", ss)
	fmt.Println(ss)

	for i, v := range ss {
		fmt.Println(i, v.i, v.s, v.b)
	}
	//slice of people of type person.
	type person struct {
		name string
		age  int
	}
	people := []person{
		{"james", 32},
		{"nathan", 33},
	}
	fmt.Printf("%T\t", people)
	fmt.Println(people)

	//range over slice
	for i, val := range people {
		fmt.Println(i, val.name, val.age)
	}

	//slice of slices
	sS := [][]string{
		[]string{"james", "bond", "bear"},
		[]string{"james", "bond", "bear"},
	}
	fmt.Println(sS)

	//append to slices
	//array
	as := []string{"senorita", "julia", "lily", "margrat", "sophia"}
	fmt.Println("before append", as)
	as = append(as, "naethan", "robert", "james", "joker")
	fmt.Println("after append", as)
}
