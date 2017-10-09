package main

import (
	"fmt"

	"sync"
)

func main() {
	in := gen()
	//FAN OUT
	//Multiple functions reading from the same channel until that channel is closed.
	//Distribute work across multiple functions (ten goroutines) that all read from in.

	c0 := factorial(in)
	c1 := factorial(in)
	c2 := factorial(in)
	c3 := factorial(in)
	c4 := factorial(in)
	c5 := factorial(in)
	c6 := factorial(in)
	c7 := factorial(in)
	c8 := factorial(in)
	c9 := factorial(in)

	num := 0
	//FAN IN
	for n := range merge(c0, c1, c2, c3, c4, c5, c6, c7, c8, c9) {
		num++
		fmt.Println(num, ": ", n)
	}
}

func gen() <-chan int64 {
	out := make(chan int64)
	go func() {
		for i := 0; i < 100; i++ {
			for j := 3; j < 13; j++ {
				out <- int64(j)
			}
		}
		close(out)
	}()
	return out
}
func factorial(in <-chan int64) <-chan int64 {
	out := make(chan int64)
	go func() {
		for n := range in {
			out <- fact(n)

		}
		close(out)
	}()
	return out
}

func fact(n int64) int64 {
	total := int64(1)
	for i := n; i > 0; i-- {
		total *= int64(i)
	}
	return total
}

func merge(cs ...<-chan int64) chan int64 {
	var wg sync.WaitGroup

	out := make(chan int64)

	output := func(c <-chan int64) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	//Start a goroutine to close out once all the output goroutines are done.
	//This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()

	return out

}
