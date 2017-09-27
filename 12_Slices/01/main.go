package main

import "fmt"

func main() {
	mySlice := make([]int, 0)
	fmt.Println("_______________________________")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println((cap(mySlice)))
	fmt.Println("_______________________________")
	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Length: ", len(mySlice), " -- Cap: ", cap(mySlice), " -- Value: ", mySlice[i])
	}
}
