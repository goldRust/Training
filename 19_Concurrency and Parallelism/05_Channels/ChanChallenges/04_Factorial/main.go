package main

import "fmt"

//func main() {
//	f:= factorial(4)
//	fmt.Println("Total:", f)
//}
//
//
//func factorial(n int)int{
//	total:=1
//	for i:=n; i>0;i--{
//		total *=i
//	}
//	return total
//
//}

//Use goroutines and channels to calc factorials

func main() {
	channel := make(chan int)
	n := 10
	go func() {
		for i := n; i > 0; i-- {
			channel <- i
		}
		close(channel)
	}()

	total := 1
	for i := range channel {
		total *= i
	}
	fmt.Println(total)

}
