package main

import "fmt"

type example struct {
	//  |flag| coun| ter |  | p | i |
	///  ----1 ------2-------4------
	flag    bool
	counter int16 // int16 takes more space than int32 and int64
	pi      float32
}

//struct fields should be larget to smallest, or there will more padding in between the adress
//group based on correctness
func main() {

	ex1 := example{} // zero value struct
	var ex2 example  //zero value struct

	// anonymous struct or unnamed struct
	// used when a struct is used only once in a code block
	// to reduce type pollution
	anon := struct {
		name string
		age  int
	}{
		name: "Jason",
		age:  23,
	}

	//type interfere or type conversion interference
	// type cannot be converted implicitly vice-versa
	type bill int
	type chill int
	var a bill
	var b chill
	//a = b  this is wrong
	a = bill(b) // this is right
	fmt.Println(a)

}
