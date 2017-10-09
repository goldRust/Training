package main

import "fmt"

type person struct {
	First string
	last  string
	age   int
}

type doubleZero struct {
	person
	First         string
	LicenseToKill bool
}

func main() {
	p1 := doubleZero{
		person{"James", "Bond", 20},
		"Double Zero Seven",
		true,
	}
	p2 := doubleZero{
		person{"Miss", "MoneyPenny", 19},
		"If looks could kill...",
		false,
	}
	fmt.Println(p1.First, p1.person.First)
	fmt.Println(p2.First, p2.person.First)
}
