package main

import "fmt"

type person struct {
	Name string
	age  int
}
type doubleZero struct {
	person
	LicenseToKill bool
}

func (p person) greeting() {
	fmt.Println("I'm just a regular person.")
}
func (dz doubleZero) greeting() {
	fmt.Println("Miss Moneypenny, so good to see you.")
}
func main() {
	p1 := person{
		"Ian Flemming",
		44,
	}
	p2 := doubleZero{
		person{
			"James Bond",
			23,
		},
		true,
	}
	p1.greeting()
	p2.greeting()
	p2.person.greeting()
}
