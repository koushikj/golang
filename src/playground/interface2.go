package main

import (
	"fmt"
	"math"
)

type square struct {
	length float32
}

type circle struct {
	radius float32
}

func (s square) area() float32 {
	return s.length * s.length
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float32
}

func findArea(s shape) float32 {
	return s.area()
}

func main() {
	fmt.Printf("%T\n", 1201.2112) //to find system default float type - whether 32 or 64 bit

	s := square{10}
	c := circle{10}

	//using normal strut method call
	fmt.Println(s.area())
	fmt.Println(c.area())

	//using interface
	fmt.Println(findArea(s))
	fmt.Println(findArea(c))
}
