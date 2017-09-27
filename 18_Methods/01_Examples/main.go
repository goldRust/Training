package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
	circumference() float64
}

//Method set of circle includes circumference and area.
func (c circle) circumference() float64 {
	return 2 * math.Pi * c.radius
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func info(s shape) {
	fmt.Println("Circumference:", s.circumference())
	fmt.Println("Area:", s.area())
}
func main() {

	c := circle{5}
	info(c)

}
