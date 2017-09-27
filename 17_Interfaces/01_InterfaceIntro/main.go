package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}
type Circle struct {
	radius float64
}
type Rectangle struct {
	height float64
	width  float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (z Square) area() float64 {
	return z.side * z.side
}
func (r Rectangle) area() float64 {
	return r.height * r.width
}

type Shape interface {
	area() float64
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := Square{10}
	c := Circle{10}
	r := Rectangle{10, 12}
	info(s)
	info(c)
	info(r)
}
