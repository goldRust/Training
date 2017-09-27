package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3, 4, 5, 2, 30, 4, 50, 2, 3, 4, 5)

	//FAN OUT
	//distribute the sq work across two goroutines that both read from in.
	c1 := sq(in)
	c2 := sq(in)
	c3 := sq(in)
	c4 := sq(in)

	//FAN IN
	//consume the merged output from c1 and c2.

	for n := range merge(c1, c2, c4, c3) {
		fmt.Println(n) //Prints 4 then 9 or 9 then 4...
	}
}

func gen(nums ...int) chan int {

	fmt.Printf("TYPE OF NUMS: %T \n", nums)
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out

}

func merge(cs ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	fmt.Printf("TYPE of CS: %T \n", cs) // just FYI
	wg.Add(len(cs))

	for _, c := range cs {
		go func(ch chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
