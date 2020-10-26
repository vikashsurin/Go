package main

import "fmt"

func main(){
	const x = 2
	fmt.Printf("%T\n",x)


	const (
		a = iota
		b
		c
		d
	)
	fmt.Println(a,b,c,d)
}