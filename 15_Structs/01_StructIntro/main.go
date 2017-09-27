package main

import "fmt"

type foo int

func main() {
	var myAge foo
	myAge = 33
	fmt.Printf("%T %v \n", myAge, myAge)

	var yourAge int
	yourAge = 29
	fmt.Printf("%T %v \n", yourAge, yourAge)

	// Nope - can't mix tipes ourAge:= myAge+yourAge

	//Convert first
	ourAge := yourAge + int(myAge)

	fmt.Printf("%T %v \n", ourAge, ourAge)
	//how about this convert
	thisAge := myAge + foo(yourAge)
	fmt.Printf("%T %v \n", thisAge, thisAge)
}
