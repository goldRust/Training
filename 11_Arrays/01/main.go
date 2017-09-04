package main

import "fmt"
func main(){
	var x [90]string
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
	for i:=33;i<=122;i++{
		x[i-33]=string(i)
	}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
}
