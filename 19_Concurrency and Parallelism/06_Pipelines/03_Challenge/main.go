package main

import "fmt"

func main() {

	for n := range factorial(oneHundred()) {
		fmt.Println(n)
	}
}

func factorial(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			total := 1
			for i := n; i > 0; i-- {
				total *= i
			}
			out <- total
		}
		close(out)
	}()
	return out
}

func oneHundred() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i <= 10; i++ {
			for j := 2; j < 12; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}
