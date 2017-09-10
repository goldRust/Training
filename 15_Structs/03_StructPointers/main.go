package main

import "fmt"

type person struct{
	first string
	last string
	age int
}

type DoubleZero struct {
	person
	LicenseToKill bool
}


func (p person) fullName() string{
	return p.last + p.first + p.last
}

func main(){
	p1:= DoubleZero{
		person{"James","Bond",20},
		true,
	}
p2:= DoubleZero{person{"Miss", "MoneyPenny",19},
false,}

fmt.Println(p1,p2)
}
