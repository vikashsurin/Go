package main

import "fmt"

/*
 constants in go are flexible .
 they are stored in 256 bits.
 they do not have a TYPE,
 rather they have a KIND which can be promoted into TYPES.

*/

func main() {
	//CASE:1
	/*
	   const doesnot have a fixed type , so x and y doesnot yet have a type.
	   so it was possible to add x and y of different kind{or say types}
	   z := x+y

	*/
	const x = 43.2 //KIND FLOAT64
	const y = 2    //KIND INT
	z := x + y     //type is FLOAT64

	//CASE:2
	/*
	   n, k and l are variables not constants.
	   the variable have defined  TYPE  at the time of assignment.
	   so l := n+k doesnot work as n and k are of different types
	*/

	// n := 43.2 //TYPE FLOAT64
	// k := 2    //TYPE
	// l := n + k // addition of float and int is not valid with variables
	// fmt.Println(l)

	//CASE:3
	/*
	 In this case, nn , kk and ll are of the same TYPE.
	 so it was possible to add them. there is not mutation/conversion of types.
	*/
	var nn float64 = 43.2 //TYPE FLOAT64
	var kk float64 = 2    //TYPE INT
	// l := n + k // addition of float and int is not valid with variables
	var ll float64 = nn + kk
	fmt.Println(ll)

	var xx int32 = 3
	var zz int32 = xx

	fmt.Printf("%T\n", z)
	fmt.Println(z)
	fmt.Printf("%T\n", zz)

	// var a int32 = 1
	// var b int32 = 1
	// // const b int32 = 1
	// var c int32 = 1
	// fmt.Printf("%T %T %T ", a, b, c)
	// fmt.Println(&a, &b, &c)
}
