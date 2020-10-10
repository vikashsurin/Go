package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}
type circle struct {
	radius float64
}

// func (s square) area() float64 {
// 	side := s.side
// 	a := side * side
// 	fmt.Println("The area of a square is : ", a)
// }

func (s square) area() float64 {
	return s.side * s.side
}

// func (c circle) area() float64 {
// 	r := c.radius
// 	a := (22 / 7) * (r * r)
// 	fmt.Println("The area of a circle is : ", a)
// }

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	x := s.area()
	fmt.Println(x)
}

func main() {
	sq := square{
		side: 3,
	}

	ci := circle{
		radius: 3,
	}

	info(sq)
	info(ci)

}
