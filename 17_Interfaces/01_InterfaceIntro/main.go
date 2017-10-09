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
type rectangle struct {
	height float64
	width  float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (z square) area() float64 {
	return z.side * z.side
}
func (r rectangle) area() float64 {
	return r.height * r.width
}

type shape interface {
	area() float64
}

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := square{10}
	c := circle{10}
	r := rectangle{10, 12}
	info(s)
	info(c)
	info(r)
}
