package main

import "fmt"

type Person struct {
	Name string
	age  int
}
type DoubleZero struct {
	Person
	LicenseToKill bool
}

func (p Person) Greeting() {
	fmt.Println("I'm just a regular person.")
}
func (dz DoubleZero) Greeting() {
	fmt.Println("Miss Moneypenny, so good to see you.")
}
func main() {
	p1 := Person{
		"Ian Flemming",
		44,
	}
	p2 := DoubleZero{
		Person{
			"James Bond",
			23,
		},
		true,
	}
	p1.Greeting()
	p2.Greeting()
	p2.Person.Greeting()
}
